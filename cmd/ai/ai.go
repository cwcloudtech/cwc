package ai

import (
	"cwc/cmd/ai/adapters"
	"cwc/cmd/ai/prompt"

	"github.com/spf13/cobra"
)

// bucketCmd represents the bucket command
var AiCmd = &cobra.Command{
	Use:   "ai",
	Short: "Cwai APIs",
	Long:  `This command lets you call the CWAI endpoints`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	AiCmd.DisableFlagsInUseLine = true
	AiCmd.AddCommand(adapters.AiAdaptersCmd)
	AiCmd.AddCommand(prompt.PromptCmd)
}
