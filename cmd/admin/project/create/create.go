/*
Copyright © 2022 comwork.io contact@comwork.io
*/
package create

import (
	"cwc/handlers/admin"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	name       string
	host       string
	token      string
	git        string
	namespace  string
	user_email string
)

// createCmd represents the create command
var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a project in the cloud",
	Long: `This command lets you create a project in the cloud.
You have to provide the project name.
You can also provide your Gitlab host and access token and git username to save the project in another Gitlab instance`,
	Run: func(cmd *cobra.Command, args []string) {
		admin.HandleAddProject(&user_email, &name, &host, &token, &git, &namespace)
	},
}

func init() {
	CreateCmd.Flags().StringVarP(&name, "name", "n", "", "The project name")
	CreateCmd.Flags().StringVarP(&host, "host", "l", "", "Gitlab host")
	CreateCmd.Flags().StringVarP(&user_email, "user", "u", "", "user associeted with the project")
	CreateCmd.Flags().StringVarP(&token, "token", "t", "", "Gitlab Token")
	CreateCmd.Flags().StringVarP(&git, "git", "g", "", "Git username")
	CreateCmd.Flags().StringVarP(&namespace, "namespace", "s", "", "Gitlab Group ID")

	if err := CreateCmd.MarkFlagRequired("name"); err != nil {
		fmt.Println(err)
	}
	if err := CreateCmd.MarkFlagRequired("user"); err != nil {
		fmt.Println(err)
	}
}
