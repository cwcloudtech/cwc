package agent

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	promptText string
	serverURL  string
	modelName  string
	provider   string
)

// AgentCmd sends a prompt to an AI agent that can call MCP server tools.
var AgentCmd = &cobra.Command{
	Use:   "agent",
	Short: "Send a prompt to an AI agent backed by MCP server tools",
	Long:  "Send a prompt to an AI agent that uses MCP server tools to execute the request",
	RunE: func(cmd *cobra.Command, args []string) error {
		if promptText == "" {
			return fmt.Errorf("prompt is required (use -p or --prompt)")
		}

		agent := NewLLMAgent(serverURL, modelName, provider)
		result, err := agent.ProcessPrompt(promptText)
		if err != nil {
			return fmt.Errorf("failed to process prompt: %w", err)
		}

		fmt.Println(result)
		return nil
	},
}

func init() {
	AgentCmd.Flags().StringVarP(&promptText, "prompt", "p", "", "The prompt to send to the LLM")
	AgentCmd.Flags().StringVarP(&serverURL, "server", "s", "http://127.0.0.1:8080/mcp", "The MCP server URL")
	AgentCmd.Flags().StringVarP(&modelName, "model", "m", "meta-llama/llama-3.3-8b-instruct:free", "The LLM model to use")
	AgentCmd.Flags().StringVar(&provider, "provider", "openai", "LLM provider: openai or anthropic")
	AgentCmd.MarkFlagRequired("prompt")
}
