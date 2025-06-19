package adapter

import (
	"cwc/cmd/ai/adapter/create"
	"cwc/cmd/ai/adapter/delete"
	"cwc/cmd/ai/adapter/ls"
	"cwc/cmd/ai/adapter/update"

	"github.com/spf13/cobra"
)

var AdapterCmd = &cobra.Command{
	Use:   "adapter",
	Short: "See default and public external AI adapters and manage the ones created by you.",
	Long:  `This command lets you see default and public external AI adapters and manage the ones created by you.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	AdapterCmd.DisableFlagsInUseLine = true
	AdapterCmd.AddCommand(ls.LsCmd)
	AdapterCmd.AddCommand(create.CreateCmd)
	AdapterCmd.AddCommand(update.UpdateCmd)
	AdapterCmd.AddCommand(delete.DeleteCmd)
}
