package create

import (
	"cwc/handlers/admin"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	name       string
	reg_type   string
	user_email string
)

// createCmd represents the create command
var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a registry with cwcloud",
	Long:  `This command lets you create a registry with cwcloud`,
	Run: func(cmd *cobra.Command, args []string) {
		admin.HandleAddRegistry(&user_email, &name, &reg_type)
	},
}

func init() {
	CreateCmd.Flags().StringVarP(&name, "name", "n", "", "The registry name")
	CreateCmd.Flags().StringVarP(&user_email, "user", "u", "", "user associated with the project")
	CreateCmd.Flags().StringVarP(&reg_type, "type", "t", "", "The registry type (private/public-read)")

	err := CreateCmd.MarkFlagRequired("name")
	if nil != err {
		fmt.Println(err)
	}

	err = CreateCmd.MarkFlagRequired("user")
	if nil != err {
		fmt.Println(err)
	}
}
