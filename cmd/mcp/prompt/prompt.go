package prompt

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

// PromptCmd sends a prompt to a tiny LLM which calls the MCP server.
var PromptCmd = &cobra.Command{
	Use:   "prompt",
	Short: "Send a prompt to a tiny LLM backed by MCP server",
	Long:  "Send a prompt to a tiny LLM that uses MCP server tools to execute the request",
	RunE: func(cmd *cobra.Command, args []string) error {
		if promptText == "" {
			return fmt.Errorf("prompt is required (use -p or --prompt)")
		}

		// Create the LLM agent with access to the MCP server
		agent := NewLLMAgent(serverURL, modelName, provider)

		// Process the prompt using the LLM
		result, err := agent.ProcessPrompt(promptText)
		if err != nil {
			return fmt.Errorf("failed to process prompt: %w", err)
		}

		// Print the result
		fmt.Println(result)
		return nil
	},
}

func init() {
	PromptCmd.Flags().StringVarP(&promptText, "prompt", "p", "", "The prompt to send to the LLM")
	PromptCmd.Flags().StringVarP(&serverURL, "server", "s", "http://127.0.0.1:8080/mcp", "The MCP server URL")
	PromptCmd.Flags().StringVarP(&modelName, "model", "m", "meta-llama/llama-3.3-8b-instruct:free", "The LLM model to use")
	PromptCmd.Flags().StringVar(&provider, "provider", "openai", "LLM provider: openai or anthropic")
	PromptCmd.MarkFlagRequired("prompt")
}
