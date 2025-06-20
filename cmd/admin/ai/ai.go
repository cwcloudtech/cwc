package ai

import (
	"cwc/cmd/admin/ai/adapter"

	"github.com/spf13/cobra"
)

var AiCmd = &cobra.Command{
	Use:   "ai",
	Short: "Admin AI management commands",
	Long:  `This command lets you manage AI adapters as an administrator`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	AiCmd.DisableFlagsInUseLine = true
	AiCmd.AddCommand(adapter.AdapterCmd)
}
