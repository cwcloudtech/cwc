package mcp

import (
	"cwc/cmd/ai/mcp/start"

	"github.com/spf13/cobra"
)

// McpCmd represents the MCP command group under ai.
var McpCmd = &cobra.Command{
	Use:   "mcp",
	Short: "Start the cwc MCP server",
	Long:  "Start the cwc MCP server",
	RunE:  start.StartCmd.RunE,
}

func init() {
	McpCmd.DisableFlagsInUseLine = true
	McpCmd.Flags().AddFlagSet(start.StartCmd.Flags())
}
