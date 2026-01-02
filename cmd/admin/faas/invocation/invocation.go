package invocation

import (
	"cwc/cmd/admin/faas/invocation/ls"

	"github.com/spf13/cobra"
)

var InvocationCmd = &cobra.Command{
	Use:   "invocation",
	Short: "Manage your invocations with cwcloud",
	Long: `This command lets you manage your invocations with cwcloud.
Several actions are associated with this command such as
listing your available invocations`,

	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	InvocationCmd.DisableFlagsInUseLine = true
	InvocationCmd.AddCommand(ls.LsCmd)
}
