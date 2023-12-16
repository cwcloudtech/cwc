package invocation

import (
	"cwc/cmd/faas/invocation/ls"
	"cwc/cmd/faas/invocation/delete"
	"cwc/cmd/faas/invocation/create"
	"cwc/cmd/faas/invocation/truncate"

	"github.com/spf13/cobra"
)

var InvocationCmd = &cobra.Command{
	Use:   "invocation",
	Short: "Manage your invocations in the cloud",
	Long: `This command lets you manage your invocations in the cloud.
Several actions are associated with this command such as update a invocation, deleting a invocation
and listing your available invocations`,

	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	InvocationCmd.DisableFlagsInUseLine = true
	InvocationCmd.AddCommand(ls.LsCmd)
	InvocationCmd.AddCommand(create.CreateCmd)
	InvocationCmd.AddCommand(delete.DeleteCmd)
	InvocationCmd.AddCommand(truncate.TruncateCmd)
}