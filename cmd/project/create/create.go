/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package create

import (
	"cwc/handlers"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	name      string
	host      string
	token     string
	git       string
	namespace string
)

// createCmd represents the create command
var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a project in the cloud",
	Long: `This command lets you create a project in the cloud.
You have to provide the project name.
You can also provide your Gitlab host and access token and git username to save the project in another Gitlab instance`,
	Run: func(cmd *cobra.Command, args []string) {
		handlers.HandleAddProject(&name, &host, &token, &git, &namespace)
	},
}

func init() {
	CreateCmd.Flags().StringVarP(&name, "name", "n", "", "The project name")
	CreateCmd.Flags().StringVarP(&host, "host", "l", "", "Gitlab host")
	CreateCmd.Flags().StringVarP(&token, "token", "t", "", "Gitlab Token")
	CreateCmd.Flags().StringVarP(&git, "git", "g", "", "Git username")
	CreateCmd.Flags().StringVarP(&namespace, "namespace", "s", "", "Gitlab Group ID")

	if err := CreateCmd.MarkFlagRequired("name"); err != nil {
		fmt.Println(err)
	}
}
