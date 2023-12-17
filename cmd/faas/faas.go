package faas

import (
	"cwc/cmd/faas/function"
	"cwc/cmd/faas/invocation"
	"cwc/cmd/faas/languages"
	"cwc/cmd/faas/trigger"

	"github.com/spf13/cobra"
)

var FaasCmd = &cobra.Command{
	Use:   "faas",
	Short: "Manage your serverless functions in the cloud",
	Long: `This command lets you manage your functions as a service in the cloud.
Several actions are associated with this command such as update a function, deleting a function
and listing your available functions`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	FaasCmd.DisableFlagsInUseLine = true
	FaasCmd.AddCommand(languages.LanguagesCmd)
	FaasCmd.AddCommand(function.FunctionCmd)
	FaasCmd.AddCommand(invocation.InvocationCmd)
	FaasCmd.AddCommand(trigger.TriggerCmd)
}
