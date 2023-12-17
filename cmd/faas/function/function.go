package function

import (
	"cwc/cmd/faas/function/create"
	"cwc/cmd/faas/function/delete"
	"cwc/cmd/faas/function/ls"
	"cwc/cmd/faas/function/update"

	"github.com/spf13/cobra"
)

var FunctionCmd = &cobra.Command{
	Use:   "function",
	Short: "Manage your serverless functions in the cloud",
	Long: `This command lets you Manage your serverless functions in the cloud.
Several actions are associated with this command such as update a function, deleting a function
and listing your available functions`,

	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	FunctionCmd.DisableFlagsInUseLine = true
	FunctionCmd.AddCommand(ls.LsCmd)
	FunctionCmd.AddCommand(delete.DeleteCmd)
	FunctionCmd.AddCommand(create.CreateCmd)
	FunctionCmd.AddCommand(update.UpdateCmd)
}
