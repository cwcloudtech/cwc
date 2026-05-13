package web_agent

import (
	"cwc/cmd/ai/agent"
	"cwc/config"
	"encoding/json"
	"fmt"
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
	Status   string                 `json:"status"`
	Response string                 `json:"response"`
	Usage    map[string]interface{} `json:"usage"`
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
			Status:   "error",
			Response: "method not allowed",
			Usage:    map[string]interface{}{},
		})
		return
	}

	var req webAgentRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(webAgentResponse{
			Status:   "error",
			Response: "invalid JSON payload",
			Usage:    map[string]interface{}{},
		})
		return
	}

	messages := make([]agent.AgentConversationMessage, 0, len(req.Messages)+1)
	for _, message := range req.Messages {
		role := strings.ToLower(strings.TrimSpace(message.Role))
		if role == "" {
			role = "user"
		}

		messages = append(messages, agent.AgentConversationMessage{
			Role:    role,
			Message: message.Message,
		})
	}

	if strings.TrimSpace(req.Message) != "" {
		messages = append(messages, agent.AgentConversationMessage{Role: "user", Message: req.Message})
	}

	if len(messages) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(webAgentResponse{
			Status:   "error",
			Response: "either message or messages is required",
			Usage:    map[string]interface{}{},
		})
		return
	}

	result, err := llmAgent.ProcessConversationWithUsage(messages, agent.AgentSettings{MaxTokens: req.Settings.MaxTokens})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(webAgentResponse{
			Status:   "error",
			Response: err.Error(),
			Usage:    map[string]interface{}{},
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
		Status:   "ok",
		Response: result.Response,
		Usage:    usage,
	})
}
