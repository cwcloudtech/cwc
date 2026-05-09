package agent

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
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
	openAIBaseURL   string
}

// NewLLMAgent creates a new LLM agent connected to an MCP server
func NewLLMAgent(serverURL string, modelName string, provider string) *LLMAgent {
	baseURL := strings.TrimSpace(os.Getenv("OPENAI_BASE_URL"))
	if baseURL == "" {
		if strings.Contains(modelName, ":free") {
			baseURL = "https://openrouter.ai/api/v1"
		} else {
			baseURL = "https://api.openai.com/v1"
		}
	}

	return &LLMAgent{
		serverURL:       serverURL,
		modelName:       modelName,
		provider:        strings.ToLower(strings.TrimSpace(provider)),
		httpClient:      &http.Client{},
		anthropicAPIKey: strings.TrimSpace(os.Getenv("ANTHROPIC_API_KEY")),
		openAIAPIKey:    strings.TrimSpace(firstNonEmpty(os.Getenv("OPENAI_API_KEY")),
		openAIBaseURL:   strings.TrimRight(baseURL, "/"),
	}
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if strings.TrimSpace(value) != "" {
			return value
		}
	}
	return ""
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

	if err := agent.initializeMCPClient(ctx); err != nil {
		return "", fmt.Errorf("failed to initialize MCP client: %w", err)
	}

	if !agent.hasProviderCredentials() {
		command, args, ok := deriveCommandFromPrompt(prompt)
		if !ok {
			return "", fmt.Errorf("could not map prompt to a cwc command without LLM credentials; set OPENAI_API_KEY or ANTHROPIC_API_KEY, or use an explicit command-like prompt such as 'instance ls'")
		}
		return agent.callTool(ctx, "run_cwc_command", map[string]interface{}{
			"command": command,
			"args":    args,
		})
	}

	// Get available tools from the MCP server
	toolsResp, err := agent.client.ListTools(ctx, nil)
	if err != nil {
		return "", fmt.Errorf("failed to list MCP tools: %w", err)
	}

	if toolsResp == nil || len(toolsResp.Tools) == 0 {
		return "", fmt.Errorf("no tools available from MCP server")
	}

	toolSchema := map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"command": map[string]interface{}{
				"type":        "string",
				"description": "The cwc command to run",
			},
			"args": map[string]interface{}{
				"type":        "array",
				"items":       map[string]interface{}{"type": "string"},
				"description": "Command arguments",
			},
		},
		"required": []string{"command"},
	}

	// Build Claude and OpenAI tool definitions from MCP tools
	claudeTools := make([]ClaudeTool, 0, len(toolsResp.Tools))
	openAITools := make([]OpenAITool, 0, len(toolsResp.Tools))
	for _, tool := range toolsResp.Tools {
		description := ""
		if tool.Description != nil {
			description = *tool.Description
		}
		claudeTool := ClaudeTool{
			Name:        tool.Name,
			Description: description,
			InputSchema: toolSchema,
		}
		openAITool := OpenAITool{
			Type: "function",
			Function: OpenAIFunctionTool{
				Name:        tool.Name,
				Description: description,
				Parameters:  toolSchema,
			},
		}
		claudeTools = append(claudeTools, claudeTool)
		openAITools = append(openAITools, openAITool)
	}

	modelText := ""
	if agent.provider == "anthropic" {
		modelText, err = agent.runAnthropic(ctx, prompt, claudeTools)
	} else {
		modelText, err = agent.runOpenAI(ctx, prompt, openAITools)
	}
	if err == nil && strings.TrimSpace(modelText) != "" {
		return modelText, nil
	}

	if command, args, ok := deriveCommandFromPrompt(prompt); ok {
		return agent.callTool(ctx, "run_cwc_command", map[string]interface{}{
			"command": command,
			"args":    args,
		})
	}

	if err != nil {
		return "", err
	}

	return "No response from model", nil
}

func (agent *LLMAgent) hasProviderCredentials() bool {
	if agent.provider == "anthropic" {
		return strings.TrimSpace(agent.anthropicAPIKey) != ""
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
		Model:      agent.modelName,
		Messages:   []OpenAIMessage{{Role: "user", Content: prompt}},
		Tools:      openAITools,
		ToolChoice: "auto",
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

func deriveCommandFromPrompt(prompt string) (string, []string, bool) {
	normalized := strings.TrimSpace(strings.ToLower(prompt))
	if normalized == "" {
		return "", nil, false
	}

	if strings.HasPrefix(normalized, "cwc ") {
		parts := strings.Fields(strings.TrimSpace(prompt))
		if len(parts) >= 2 {
			return parts[1], parts[2:], true
		}
	}

	if (strings.Contains(normalized, "instance") || strings.Contains(normalized, "instances")) &&
		(strings.Contains(normalized, "list") || strings.Contains(normalized, "listen") || strings.Contains(normalized, "show")) {
		return "instance", []string{"ls"}, true
	}

	parts := strings.Fields(strings.TrimSpace(prompt))
	if len(parts) >= 2 {
		return parts[0], parts[1:], true
	}

	return "", nil, false
}

// callClaude calls the Claude API
func (agent *LLMAgent) callClaude(ctx context.Context, req ClaudeRequest) (*ClaudeResponse, error) {
	if agent.anthropicAPIKey == "" {
		return nil, fmt.Errorf("ANTHROPIC_API_KEY environment variable is not set")
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", "https://api.anthropic.com/v1/messages", bytes.NewReader(reqBody))
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
	if agent.openAIAPIKey == "" {
		return nil, fmt.Errorf("OPENAI_API_KEY environment variable is not set")
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal OpenAI request: %w", err)
	}

	endpoint := agent.openAIBaseURL + "/chat/completions"
	httpReq, err := http.NewRequestWithContext(ctx, "POST", endpoint, bytes.NewReader(reqBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create OpenAI request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+agent.openAIAPIKey)
	if strings.Contains(agent.openAIBaseURL, "openrouter.ai") {
		httpReq.Header.Set("HTTP-Referer", "https://github.com/ineumann/cwc")
		httpReq.Header.Set("X-Title", "cwc-mcp-prompt")
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
