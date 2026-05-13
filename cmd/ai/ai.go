package ai

import (
	"cwc/cmd/ai/adapter"
	"cwc/cmd/ai/agent"
	"cwc/cmd/ai/mcp"
	"cwc/cmd/ai/prompt"
	"cwc/cmd/ai/web_agent"

	"github.com/spf13/cobra"
)

var AiCmd = &cobra.Command{
	Use:   "ai",
	Short: "AI features and tools",
	Long:  `This command lets you call the AI agent, manage prompts, and more.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	AiCmd.DisableFlagsInUseLine = true
	AiCmd.AddCommand(adapter.AdapterCmd)
	AiCmd.AddCommand(agent.AgentCmd)
	AiCmd.AddCommand(mcp.McpCmd)
	AiCmd.AddCommand(prompt.PromptCmd)
	AiCmd.AddCommand(web_agent.WebAgentCmd)
}
