package agent

import (
	"bytes"
	"context"
	"cwc/config"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strings"

	mcp_golang "github.com/metoro-io/mcp-golang"
	mcp_http_transport "github.com/metoro-io/mcp-golang/transport/http"
)

// LLMAgent represents an LLM that can call MCP server tools
type LLMAgent struct {
	serverURL       string
	modelName       string
	provider        string
	client          *mcp_golang.Client
	httpClient      *http.Client
	anthropicAPIKey string
	openAIAPIKey    string
	openRouterAPIKey string
	providerBaseUrl  string
}

// NewLLMAgent creates a new LLM agent connected to an MCP server
func NewLLMAgent(serverURL string, modelName string, provider string) *LLMAgent {
	providerName := strings.ToLower(strings.TrimSpace(provider))
	
	// Set default model based on provider if not specified
	if strings.TrimSpace(modelName) == "" {
		switch providerName {
		case "openrouter":
			modelName = "meta-llama/llama-3.3-8b-instruct"
		case "anthropic":
			modelName = "claude-haiku-4-5"
		default:
			modelName = "gpt-4o-mini"
		}
	}
	
	baseURL := ""
	tokenKey := ""
	
	switch providerName {
	case "openrouter":
		baseURL = strings.TrimSpace(config.GetOpenRouterBaseURL())
		tokenKey = strings.TrimSpace(config.GetOpenRouterAPIKey())
	case "anthropic":
		baseURL = strings.TrimSpace(config.GetAnthropicBaseURL())
		tokenKey = strings.TrimSpace(config.GetAnthropicAPIKey())
	default:
		baseURL = strings.TrimSpace(config.GetOpenAIBaseURL())
		tokenKey = strings.TrimSpace(config.GetOpenAIAPIKey())
	}

	return &LLMAgent{
		serverURL:        serverURL,
		modelName:        modelName,
		provider:         providerName,
		httpClient:       &http.Client{},
		anthropicAPIKey:  tokenKey,
		openAIAPIKey:     tokenKey,
		openRouterAPIKey: tokenKey,
		providerBaseUrl:  strings.TrimRight(baseURL, "/"),
	}
}

// ClaudeMessage represents a message in Claude API format
type ClaudeMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ClaudeRequest represents a request to Claude API
type ClaudeRequest struct {
	Model     string          `json:"model"`
	MaxTokens int             `json:"max_tokens"`
	Messages  []ClaudeMessage `json:"messages"`
	Tools     []ClaudeTool    `json:"tools,omitempty"`
}

// ClaudeTool represents a tool definition for Claude
type ClaudeTool struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	InputSchema interface{} `json:"input_schema"`
}

// ClaudeResponse represents Claude's response
type ClaudeResponse struct {
	Content []ContentBlock `json:"content"`
	Error   *ErrorInfo     `json:"error,omitempty"`
}

// ContentBlock represents a block in Claude's response
type ContentBlock struct {
	Type  string      `json:"type"`
	Text  string      `json:"text,omitempty"`
	ID    string      `json:"id,omitempty"`
	Name  string      `json:"name,omitempty"`
	Input interface{} `json:"input,omitempty"`
}

// ErrorInfo represents error information
type ErrorInfo struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

type OpenAIMessage struct {
	Role       string           `json:"role"`
	Content    string           `json:"content,omitempty"`
	ToolCalls  []OpenAIToolCall `json:"tool_calls,omitempty"`
	ToolCallID string           `json:"tool_call_id,omitempty"`
}

type OpenAITool struct {
	Type     string             `json:"type"`
	Function OpenAIFunctionTool `json:"function"`
}

type OpenAIFunctionTool struct {
	Name        string      `json:"name"`
	Description string      `json:"description,omitempty"`
	Parameters  interface{} `json:"parameters"`
}

type OpenAIToolCall struct {
	ID       string             `json:"id"`
	Type     string             `json:"type"`
	Function OpenAIFunctionCall `json:"function"`
}

type OpenAIFunctionCall struct {
	Name      string `json:"name"`
	Arguments string `json:"arguments"`
}

type OpenAIChatCompletionRequest struct {
	Model      string          `json:"model"`
	Messages   []OpenAIMessage `json:"messages"`
	Tools      []OpenAITool    `json:"tools,omitempty"`
	ToolChoice string          `json:"tool_choice,omitempty"`
}

type OpenAIChatCompletionResponse struct {
	Choices []struct {
		Message OpenAIMessage `json:"message"`
	} `json:"choices"`
	Error *ErrorInfo `json:"error,omitempty"`
}

// ProcessPrompt processes a prompt using the configured LLM with MCP server tools
func (agent *LLMAgent) ProcessPrompt(prompt string) (string, error) {
	ctx := context.Background()

	if !agent.hasProviderCredentials() {
		return "", fmt.Errorf("LLM credentials are required for provider %s", agent.provider)
	}

	if err := agent.initializeMCPClient(ctx); err != nil {
		return "", fmt.Errorf("failed to initialize MCP client: %w", err)
	}

	toolsResp, err := agent.client.ListTools(ctx, nil)
	if err != nil {
		return "", fmt.Errorf("failed to list MCP tools: %w", err)
	}

	if toolsResp == nil || len(toolsResp.Tools) == 0 {
		return "", fmt.Errorf("no tools available from MCP server")
	}

	dynamicRunCmdSchema := agent.buildDynamicRunCommandSchema(ctx)

	claudeTools := make([]ClaudeTool, 0, len(toolsResp.Tools))
	openAITools := make([]OpenAITool, 0, len(toolsResp.Tools))
	for _, tool := range toolsResp.Tools {
		description := ""
		if tool.Description != nil {
			description = *tool.Description
		}

		inputSchema := map[string]interface{}{
			"type":       "object",
			"properties": map[string]interface{}{},
		}
		if tool.InputSchema != nil {
			if typedSchema, ok := tool.InputSchema.(map[string]interface{}); ok {
				inputSchema = typedSchema
			}
		}
		if tool.Name == "run_cwc_command" && dynamicRunCmdSchema != nil {
			inputSchema = dynamicRunCmdSchema
		}

		claudeTools = append(claudeTools, ClaudeTool{
			Name:        tool.Name,
			Description: description,
			InputSchema: inputSchema,
		})

		openAITools = append(openAITools, OpenAITool{
			Type: "function",
			Function: OpenAIFunctionTool{
				Name:        tool.Name,
				Description: description,
				Parameters:  inputSchema,
			},
		})
	}

	openAITools = selectPreferredToolsForPrompt(
		prompt,
		openAITools,
		128,
		func(tool OpenAITool) string { return tool.Function.Name },
		func(tool OpenAITool) string { return tool.Function.Description },
	)
	claudeTools = selectPreferredToolsForPrompt(
		prompt,
		claudeTools,
		128,
		func(tool ClaudeTool) string { return tool.Name },
		func(tool ClaudeTool) string { return tool.Description },
	)

	modelText := ""
	if agent.provider == "anthropic" {
		modelText, err = agent.runAnthropic(ctx, prompt, claudeTools)
	} else {
		modelText, err = agent.runOpenAI(ctx, prompt, openAITools)
	}
	if err == nil && strings.TrimSpace(modelText) != "" {
		return modelText, nil
	}

	if err != nil {
		return "", err
	}

	return "No response from model", nil
}

func (agent *LLMAgent) buildDynamicRunCommandSchema(ctx context.Context) map[string]interface{} {
	raw, err := agent.callTool(ctx, "list_cwc_commands", map[string]interface{}{})
	if err != nil || strings.TrimSpace(raw) == "" {
		return nil
	}

	output := raw
	if idx := strings.Index(raw, "output:\n"); idx >= 0 {
		output = raw[idx+len("output:\n"):]
	}

	commands := extractTopLevelCWCCommands(output)
	if len(commands) == 0 {
		return nil
	}

	return map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"command": map[string]interface{}{
				"type":        "string",
				"description": "Top-level cwc command name",
				"enum":        commands,
			},
			"args": map[string]interface{}{
				"type":        "array",
				"items":       map[string]interface{}{"type": "string"},
				"description": "Subcommand, resource id/name, flags and values",
			},
		},
		"required": []string{"command"},
	}
}

func extractTopLevelCWCCommands(helpText string) []string {
	lines := strings.Split(helpText, "\n")
	inSection := false
	unique := map[string]bool{}

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			if inSection {
				break
			}
			continue
		}

		if strings.HasPrefix(trimmed, "Available Commands:") {
			inSection = true
			continue
		}
		if !inSection {
			continue
		}

		if strings.HasSuffix(trimmed, ":") {
			break
		}

		parts := strings.Fields(trimmed)
		if len(parts) == 0 {
			continue
		}

		cmd := strings.TrimSpace(parts[0])
		if cmd == "help" {
			continue
		}
		if strings.Contains(cmd, "[") || strings.Contains(cmd, "<") {
			continue
		}
		unique[cmd] = true
	}

	if len(unique) == 0 {
		return nil
	}

	commands := make([]string, 0, len(unique))
	for cmd := range unique {
		commands = append(commands, cmd)
	}
	sort.Strings(commands)
	return commands
}

func selectPreferredToolsForPrompt[T any](
	prompt string,
	tools []T,
	maxTools int,
	nameFn func(T) string,
	descriptionFn func(T) string,
) []T {
	if len(tools) <= maxTools || maxTools <= 0 {
		return tools
	}

	normalizedPrompt := strings.ToLower(prompt)
	replacer := strings.NewReplacer("-", " ", "_", " ", "/", " ", ",", " ", ".", " ", ":", " ", ";", " ")
	normalizedPrompt = replacer.Replace(normalizedPrompt)
	promptTokens := map[string]bool{}
	for _, part := range strings.Fields(normalizedPrompt) {
		if len(part) >= 2 {
			promptTokens[part] = true
		}
	}

	essential := map[string]bool{
		"list_cwc_commands":   true,
		"get_cwc_command_help": true,
		"run_cwc_command":     true,
	}

	type scoredTool struct {
		tool  T
		score int
	}
	scored := make([]scoredTool, 0, len(tools))

	for _, tool := range tools {
		toolName := nameFn(tool)
		name := strings.ToLower(toolName)
		desc := strings.ToLower(descriptionFn(tool))
		score := 0

		if essential[strings.ToLower(toolName)] {
			score += 1000
		}

		nameForTokens := replacer.Replace(name)
		for _, token := range strings.Fields(nameForTokens) {
			if promptTokens[token] {
				score += 10
			}
		}

		for token := range promptTokens {
			if strings.Contains(name, token) {
				score += 6
			}
			if token != "" && strings.Contains(desc, token) {
				score += 2
			}
		}

		scored = append(scored, scoredTool{tool: tool, score: score})
	}

	sort.SliceStable(scored, func(i, j int) bool {
		if scored[i].score == scored[j].score {
			return nameFn(scored[i].tool) < nameFn(scored[j].tool)
		}
		return scored[i].score > scored[j].score
	})

	selected := make([]T, 0, maxTools)
	for i := 0; i < len(scored) && len(selected) < maxTools; i++ {
		selected = append(selected, scored[i].tool)
	}

	return selected
}

func (agent *LLMAgent) hasProviderCredentials() bool {
	if agent.provider == "anthropic" {
		return strings.TrimSpace(agent.anthropicAPIKey) != ""
	}
	if agent.provider == "openrouter" {
		return true
	}
	return strings.TrimSpace(agent.openAIAPIKey) != ""
}

func (agent *LLMAgent) runAnthropic(ctx context.Context, prompt string, claudeTools []ClaudeTool) (string, error) {
	claudeReq := ClaudeRequest{
		Model:     agent.modelName,
		MaxTokens: 1024,
		Tools:     claudeTools,
		Messages: []ClaudeMessage{
			{
				Role:    "user",
				Content: prompt,
			},
		},
	}

	response, err := agent.callClaude(ctx, claudeReq)
	if err != nil {
		return "", err
	}

	for _, block := range response.Content {
		if block.Type != "tool_use" {
			continue
		}
		argsMap, ok := block.Input.(map[string]interface{})
		if !ok {
			inputJSON, _ := json.Marshal(block.Input)
			if unmarshalErr := json.Unmarshal(inputJSON, &argsMap); unmarshalErr != nil {
				return "", fmt.Errorf("failed to parse tool input: %w", unmarshalErr)
			}
		}
		return agent.callTool(ctx, block.Name, argsMap)
	}

	for _, block := range response.Content {
		if block.Type == "text" {
			return block.Text, nil
		}
	}

	return "", nil
}

func (agent *LLMAgent) runOpenAI(ctx context.Context, prompt string, openAITools []OpenAITool) (string, error) {
	req := OpenAIChatCompletionRequest{
		Model:    agent.modelName,
		Messages: []OpenAIMessage{{Role: "user", Content: prompt}},
	}
	if len(openAITools) > 0 {
		req.Tools = openAITools
		req.ToolChoice = "auto"
	}

	resp, err := agent.callOpenAI(ctx, req)
	if err != nil {
		return "", err
	}
	if len(resp.Choices) == 0 {
		return "", nil
	}

	msg := resp.Choices[0].Message
	for _, toolCall := range msg.ToolCalls {
		if toolCall.Type != "function" {
			continue
		}
		argsMap := map[string]interface{}{}
		if strings.TrimSpace(toolCall.Function.Arguments) != "" {
			if unmarshalErr := json.Unmarshal([]byte(toolCall.Function.Arguments), &argsMap); unmarshalErr != nil {
				return "", fmt.Errorf("failed to parse OpenAI tool arguments: %w", unmarshalErr)
			}
		}
		return agent.callTool(ctx, toolCall.Function.Name, argsMap)
	}

	return msg.Content, nil
}

func (agent *LLMAgent) initializeMCPClient(ctx context.Context) error {
	parsedURL, err := url.Parse(agent.serverURL)
	if err != nil {
		return fmt.Errorf("invalid server URL: %w", err)
	}
	if parsedURL.Scheme == "" || parsedURL.Host == "" {
		return fmt.Errorf("server URL must include scheme and host, e.g. http://127.0.0.1:8080/mcp")
	}

	endpoint := parsedURL.Path
	if endpoint == "" {
		endpoint = "/mcp"
	}

	baseURL := fmt.Sprintf("%s://%s", parsedURL.Scheme, parsedURL.Host)
	transport := mcp_http_transport.NewHTTPClientTransport(endpoint).WithBaseURL(baseURL)
	agent.client = mcp_golang.NewClient(transport)

	_, err = agent.client.Initialize(ctx)
	return err
}

// callClaude calls the Claude API
func (agent *LLMAgent) callClaude(ctx context.Context, req ClaudeRequest) (*ClaudeResponse, error) {
	if agent.anthropicAPIKey == "" {
		return nil, fmt.Errorf("anthropic_api_key config is not set")
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", agent.providerBaseUrl + "/messages", bytes.NewReader(reqBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("x-api-key", agent.anthropicAPIKey)
	httpReq.Header.Set("anthropic-version", "2023-06-01")

	httpResp, err := agent.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to call Claude: %w", err)
	}
	defer httpResp.Body.Close()

	body, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if httpResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Claude API error: %s", string(body))
	}

	var response ClaudeResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	if response.Error != nil {
		return nil, fmt.Errorf("Claude error: %s", response.Error.Message)
	}

	return &response, nil
}

func (agent *LLMAgent) callOpenAI(ctx context.Context, req OpenAIChatCompletionRequest) (*OpenAIChatCompletionResponse, error) {
	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal OpenAI request: %w", err)
	}

	endpoint := agent.providerBaseUrl + "/chat/completions"
	httpReq, err := http.NewRequestWithContext(ctx, "POST", endpoint, bytes.NewReader(reqBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create OpenAI request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	if agent.provider == "openrouter" && agent.openRouterAPIKey != "" {
		httpReq.Header.Set("Authorization", "Bearer " + agent.openRouterAPIKey)
	} else if agent.provider == "openai" && agent.openAIAPIKey != "" {
		httpReq.Header.Set("Authorization", "Bearer " + agent.openAIAPIKey)
	} else if agent.provider == "openai" {
		return nil, fmt.Errorf("openai_api_key config is not set")
	}

	httpResp, err := agent.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to call OpenAI API: %w", err)
	}
	defer httpResp.Body.Close()

	body, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read OpenAI response: %w", err)
	}

	if httpResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("OpenAI API error: %s", string(body))
	}

	var response OpenAIChatCompletionResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to parse OpenAI response: %w", err)
	}
	if response.Error != nil {
		return nil, fmt.Errorf("OpenAI error: %s", response.Error.Message)
	}

	return &response, nil
}

// callTool calls a tool on the MCP server
func (agent *LLMAgent) callTool(ctx context.Context, toolName string, args map[string]interface{}) (string, error) {
	result, err := agent.client.CallTool(ctx, toolName, args)
	if err != nil {
		return "", fmt.Errorf("failed to call tool %s: %w", toolName, err)
	}

	if result == nil || len(result.Content) == 0 {
		return "", nil
	}

	output := make([]string, 0, len(result.Content))
	for _, content := range result.Content {
		if content == nil {
			continue
		}
		if content.TextContent != nil {
			output = append(output, content.TextContent.Text)
			continue
		}
		fallback, marshalErr := json.Marshal(content)
		if marshalErr == nil {
			output = append(output, string(fallback))
		}
	}

	if len(output) == 0 {
		return "", nil
	}

	return strings.Join(output, "\n"), nil
}
