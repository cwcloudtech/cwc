package create

import (
	"cwc/client"
	"cwc/handlers/user"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	content client.InvocationAddContent
	interactive bool = false
	argumentsValues []string
)

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an invocation in the cloud",
	Long: `This command lets you create an invocation in the cloud.`,
	Run: func(cmd *cobra.Command, args []string) {
		user.HandleAddInvocation(&content, &argumentsValues, &interactive)
	},
}

func init() {
	CreateCmd.Flags().StringVarP(&content.Function_id, "function_id", "f", "", "The function id")
	CreateCmd.Flags().BoolVarP(&interactive, "interactive", "i", false, "Interactive mode")
	CreateCmd.Flags().StringArrayVarP(&argumentsValues, "args", "a", []string{}, "The invocation arguments values")

	if err := CreateCmd.MarkFlagRequired("function_id"); err != nil {
		fmt.Println(err)
	}
}