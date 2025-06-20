package adapter

import (
	"cwc/cmd/admin/ai/adapter/create"
	"cwc/cmd/admin/ai/adapter/delete"
	"cwc/cmd/admin/ai/adapter/ls"
	"cwc/cmd/admin/ai/adapter/update"

	"github.com/spf13/cobra"
)

var AdapterCmd = &cobra.Command{
	Use:   "adapter",
	Short: "Manage all available external adapters AI adapters",
	Long:  `This command lets you manage all avaialable external AI adapters as an administrator.`,
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
