package mcp

import (
	"cwc/cmd/mcp/start"

	"github.com/spf13/cobra"
)

// McpCmd represents the MCP command group.
var McpCmd = &cobra.Command{
	Use:   "mcp",
	Short: "Manage the cwc MCP server",
	Long:  "Manage the cwc MCP server",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	McpCmd.DisableFlagsInUseLine = true
	McpCmd.AddCommand(start.StartCmd)
}
