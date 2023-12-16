package create

import (
	"cwc/client"
	"cwc/handlers/user"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	function    client.Function
	interactive bool = false
)

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a function in the cloud",
	Long:  `This command lets you create a function in the cloud.`,
	Run: func(cmd *cobra.Command, args []string) {
		user.HandleAddFunction(&function, &interactive)
	},
}

func init() {
	CreateCmd.Flags().BoolVarP(&function.Is_public, "is_public", "p", false, "Is the function public? (optional)")
	CreateCmd.Flags().StringVarP(&function.Content.Name, "name", "n", "", "Name of the function")
	CreateCmd.Flags().StringVarP(&function.Content.Language, "language", "l", "", "Language of the function")
	CreateCmd.Flags().BoolVarP(&interactive, "interactive", "i", false, "Interactive mode")
	CreateCmd.Flags().StringVarP(&function.Content.Regexp, "regexp", "r", "", "Arguments matching regexp (optional)")
	CreateCmd.Flags().StringVarP(&function.Content.Callback_url, "callback-url", "u", "", "Callback URL of the function (optional)")
	CreateCmd.Flags().StringVarP(&function.Content.Callback_authorization_header, "callback-authorization-header", "a", "", "Callback Authorization Header of the function (optional)")
	CreateCmd.Flags().StringSliceVarP(&function.Content.Args, "args", "g", []string{}, "Arguments of the function")
	CreateCmd.Flags().StringVarP(&function.Content.Code, "code", "c", "", "Code of the function (optional)")

	err := CreateCmd.MarkFlagRequired("name")
	if nil != err {
		fmt.Println(err)
	}

	err = CreateCmd.MarkFlagRequired("language")
	if nil != err {
		fmt.Println(err)
	}
}
