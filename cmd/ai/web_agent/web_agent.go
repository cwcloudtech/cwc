package web_agent

import (
	"bytes"
	"cwc/cmd/ai/agent"
	"cwc/config"
	"cwc/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

var (
	listenAddr string
	listenPort int
	serverURL  string
	modelName  string
	provider   string
)

type webAgentSettings struct {
	MaxTokens int `json:"max_tokens"`
}

type webAgentMessage struct {
	Role    string `json:"role"`
	Message string `json:"message"`
}

type webAgentRequest struct {
	Settings webAgentSettings  `json:"settings"`
	Message  string            `json:"message"`
	Messages []webAgentMessage `json:"messages"`
}

type webAgentResponse struct {
	Status  string                 `json:"status"`
	Message string                 `json:"message"`
	Usage   map[string]interface{} `json:"usage"`
}

type gitLabProject struct {
	ID     int    `json:"id"`
	WebURL string `json:"web_url"`
}

type gitLabIssue struct {
	IID   int    `json:"iid"`
	Title string `json:"title"`
}

type gitLabUser struct {
	Name     string `json:"name"`
	Username string `json:"username"`
}

type gitLabObjectAttributes struct {
	Note         string `json:"note"`
	NoteableType string `json:"noteable_type"`
}

type gitLabNoteEvent struct {
	ObjectKind       string                 `json:"object_kind"`
	EventType        string                 `json:"event_type"`
	Project          gitLabProject          `json:"project"`
	Issue            *gitLabIssue           `json:"issue"`
	User             gitLabUser             `json:"user"`
	ObjectAttributes gitLabObjectAttributes `json:"object_attributes"`
}

type gitLabIssueCommentRequest struct {
	Body string `json:"body"`
}

var WebAgentCmd = &cobra.Command{
	Use:   "web-agent",
	Short: "Start AI agent HTTP API",
	Long:  "Start an HTTP API that accepts AI agent requests and returns agent responses",
	RunE: func(cmd *cobra.Command, args []string) error {
		agt := agent.NewLLMAgent(serverURL, modelName, provider)
		fmt.Printf("Using model: %s (provider: %s)\n", modelName, provider)

		mux := http.NewServeMux()
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			handleWebAgentRequest(w, r, agt)
		})

		mux.HandleFunc("/gitlab", func(w http.ResponseWriter, r *http.Request) {
			handleGitLabWebhook(w, r, agt)
		})

		mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
		})

		addr := fmt.Sprintf("%s:%d", listenAddr, listenPort)
		fmt.Printf("Starting AI web-agent API on http://%s\n", addr)
		return http.ListenAndServe(addr, mux)
	},
}

func init() {
	WebAgentCmd.DisableFlagsInUseLine = true
	WebAgentCmd.Flags().IntVarP(&listenPort, "port", "p", 8081, "Web-agent listen port")
	WebAgentCmd.Flags().StringVarP(&listenAddr, "address", "a", "127.0.0.1", "Web-agent listen address")
	WebAgentCmd.Flags().StringVarP(&serverURL, "server", "s", "http://127.0.0.1:8080/mcp", "The MCP server URL")
	WebAgentCmd.Flags().StringVarP(&modelName, "model", "m", config.GetDefaultAiModel(), "The LLM model to use")
	WebAgentCmd.Flags().StringVar(&provider, "provider", config.GetDefaultAiProvider(), "LLM provider: openrouter, openai, google, deepseek or anthropic")
}

func handleWebAgentRequest(w http.ResponseWriter, r *http.Request, llmAgent *agent.LLMAgent) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_ = json.NewEncoder(w).Encode(webAgentResponse{
			Status:  "error",
			Message: "method not allowed",
			Usage:   map[string]interface{}{},
		})
		return
	}

	var req webAgentRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(webAgentResponse{
			Status:  "error",
			Message: "invalid JSON payload",
			Usage:   map[string]interface{}{},
		})
		return
	}

	messages := make([]agent.AgentConversationMessage, 0, len(req.Messages)+1)
	for _, message := range req.Messages {
		role := strings.ToLower(strings.TrimSpace(message.Role))
		if utils.IsBlank(role) {
			role = "user"
		}

		messages = append(messages, agent.AgentConversationMessage{
			Role:    role,
			Message: message.Message,
		})
	}

	if utils.IsNotBlank(req.Message) {
		messages = append(messages, agent.AgentConversationMessage{Role: "user", Message: req.Message})
	}

	if len(messages) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(webAgentResponse{
			Status:  "error",
			Message: "either message or messages is required",
			Usage:   map[string]interface{}{},
		})
		return
	}

	result, err := llmAgent.ProcessConversationWithUsage(messages, agent.AgentSettings{MaxTokens: req.Settings.MaxTokens})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(webAgentResponse{
			Status:  "error",
			Message: err.Error(),
			Usage:   map[string]interface{}{},
		})
		return
	}

	usage := map[string]interface{}{}
	if result.Usage != nil {
		usage = map[string]interface{}{
			"total":      result.Usage.Total,
			"prompt":     result.Usage.Prompt,
			"completion": result.Usage.Completion,
		}
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(webAgentResponse{
		Status:  "ok",
		Message: utils.FormatWebAgentMessage(result.Response),
		Usage:   usage,
	})
}

func handleGitLabWebhook(w http.ResponseWriter, r *http.Request, llmAgent *agent.LLMAgent) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_ = json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": "method not allowed"})
		return
	}

	webhookSecret := strings.TrimSpace(config.GetGitLabWebhookSecret())
	if utils.IsNotBlank(webhookSecret) && strings.TrimSpace(r.Header.Get("X-Gitlab-Token")) != webhookSecret {
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": "invalid webhook token"})
		return
	}

	var event gitLabNoteEvent
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&event); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": "invalid JSON payload"})
		return
	}

	if !isGitLabIssueCommand(event) {
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]string{"status": "ignored"})
		return
	}

	prompt := strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(event.ObjectAttributes.Note), "/cwc"))
	responseText := ""
	if utils.IsBlank(prompt) {
		responseText = "Usage: /cwc <prompt>"
	} else {
		messages := []agent.AgentConversationMessage{{
			Role:    "user",
			Message: buildGitLabPrompt(event, prompt),
		}}

		result, err := llmAgent.ProcessConversationWithUsage(messages, agent.AgentSettings{})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": err.Error()})
			return
		}

		responseText = utils.FormatWebAgentMessage(result.Response)
	}

	if err := postGitLabIssueComment(r, event, responseText); err != nil {
		w.WriteHeader(http.StatusBadGateway)
		_ = json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func isGitLabIssueCommand(event gitLabNoteEvent) bool {
	if !strings.EqualFold(strings.TrimSpace(event.ObjectKind), "note") {
		return false
	}

	if !strings.EqualFold(strings.TrimSpace(event.ObjectAttributes.NoteableType), "Issue") {
		return false
	}

	if event.Issue == nil || event.Project.ID <= 0 || event.Issue.IID <= 0 {
		return false
	}

	note := strings.TrimSpace(event.ObjectAttributes.Note)
	return strings.HasPrefix(note, "/cwc")
}

func buildGitLabPrompt(event gitLabNoteEvent, command string) string {
	parts := []string{
		"You are responding to a GitLab issue comment.",
		fmt.Sprintf("Project ID: %d", event.Project.ID),
	}

	if event.Issue != nil {
		parts = append(parts, fmt.Sprintf("Issue IID: %d", event.Issue.IID))
		if strings.TrimSpace(event.Issue.Title) != "" {
			parts = append(parts, fmt.Sprintf("Issue title: %s", event.Issue.Title))
		}
	}

	if strings.TrimSpace(event.User.Username) != "" {
		parts = append(parts, fmt.Sprintf("Requested by GitLab user: @%s", event.User.Username))
	} else if strings.TrimSpace(event.User.Name) != "" {
		parts = append(parts, fmt.Sprintf("Requested by GitLab user: %s", event.User.Name))
	}

	parts = append(parts, "Reply with the final answer only, suitable for posting as a GitLab issue comment.")
	parts = append(parts, "User request:")
	parts = append(parts, command)

	return strings.Join(parts, "\n")
}

func postGitLabIssueComment(r *http.Request, event gitLabNoteEvent, body string) error {
	baseUrl := config.GetGitLabBaseURL()
	if utils.IsBlank(baseUrl) {
		return fmt.Errorf("gitlab_base_url configuration is required to post issue comments")
	}

	apiUrl := fmt.Sprintf("%s/api/v4", baseUrl)

	gitlabToken := strings.TrimSpace(config.GetGitLabToken())
	if utils.IsBlank(gitlabToken) {
		return fmt.Errorf("gitlab_token is required to post issue comments")
	}

	commentReq := gitLabIssueCommentRequest{Body: body}
	payload, err := json.Marshal(commentReq)
	if err != nil {
		return err
	}

	commentURL := fmt.Sprintf("%s/projects/%d/issues/%d/notes", apiUrl, event.Project.ID, event.Issue.IID)
	httpReq, err := http.NewRequest(http.MethodPost, commentURL, bytes.NewReader(payload))
	if err != nil {
		return err
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("PRIVATE-TOKEN", gitlabToken)

	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode >= 200 && httpResp.StatusCode < 300 {
		return nil
	}

	responseBody, readErr := io.ReadAll(io.LimitReader(httpResp.Body, 4096))
	if readErr != nil {
		return fmt.Errorf("gitlab comment failed with status %d", httpResp.StatusCode)
	}

	return fmt.Errorf("gitlab comment failed with status %d: %s", httpResp.StatusCode, strings.TrimSpace(string(responseBody)))
}
