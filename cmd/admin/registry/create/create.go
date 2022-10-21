/*
Copyright © 2022 comwork.io contact.comwork.io

*/
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
	Short: "Create a registry in the cloud",
	Long:  `This command lets you create a registry in the cloud`,
	Run: func(cmd *cobra.Command, args []string) {
		admin.HandleAddRegistry(&user_email, &name, &reg_type)
	},
}

func init() {
	CreateCmd.Flags().StringVarP(&name, "name", "n", "", "The registry name")
	CreateCmd.Flags().StringVarP(&user_email, "user", "u", "", "user associeted with the project")
	CreateCmd.Flags().StringVarP(&reg_type, "type", "t", "", "The registry type (private/public)")

	if err := CreateCmd.MarkFlagRequired("name"); err != nil {
		fmt.Println(err)
	}
	if err := CreateCmd.MarkFlagRequired("user"); err != nil {
		fmt.Println(err)
	}
}
