package create

import (
	"cwc/client"
	"cwc/handlers/user"
	"fmt"
	"cwc/utils"
	"github.com/spf13/cobra"
)

var (
	content         client.InvocationAddContent
	interactive     bool = false
	argumentsValues []string
	synchronous     bool = false
	pretty          bool = false
)

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an invocation in the cloud",
	Long:  "This command lets you create an invocation in the cloud.",
	Run: func(cmd *cobra.Command, args []string) {
		created_invocation, err := user.PrepareAddInvocation(&content, &argumentsValues, &interactive, &synchronous)
		utils.ExitIfError(err)
		user.HandleAddInvocation(created_invocation, &pretty)
	},
}

func init() {
	CreateCmd.Flags().StringVarP(&content.Function_id, "function_id", "f", "", "The function id")
	CreateCmd.Flags().BoolVarP(&interactive, "interactive", "i", false, "Interactive mode (optional)")
	CreateCmd.Flags().StringArrayVarP(&argumentsValues, "args", "a", []string{}, "The invocation arguments values")
	CreateCmd.Flags().BoolVarP(&synchronous, "synchronous", "s", false, "Synchronous invocation (optional)")
	CreateCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "Pretty print the output (optional)")

	err := CreateCmd.MarkFlagRequired("function_id")
	if nil != err {
		fmt.Println(err)
	}
}
