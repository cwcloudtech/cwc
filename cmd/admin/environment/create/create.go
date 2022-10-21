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
	name        string
	roles       []string
	subdomains  []string
	main_role   string
	path        string
	description string
	privacy     bool
)

// createCmd represents the create command
var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an environment in the cloud",
	Long:  `This command lets you create an environment in the cloud`,
	Run: func(cmd *cobra.Command, args []string) {
		admin.HandleAddEnvironment(&name, &path, &roles, &main_role, &privacy, &description, &subdomains)
	},
}

func init() {
	CreateCmd.Flags().StringVarP(&name, "name", "n", "", "The environment name")
	CreateCmd.Flags().StringSliceVarP(&roles, "roles", "r", []string{}, "The environment roles")
	CreateCmd.Flags().StringSliceVarP(&roles, "subdomains", "s", []string{}, "The environment subdomains")
	CreateCmd.Flags().StringVarP(&main_role, "main-role", "m", "", "The environment main role")
	CreateCmd.Flags().StringVarP(&path, "path", "p", "", "The environment path")
	CreateCmd.Flags().StringVarP(&description, "description", "d", "", "The environment description")
	CreateCmd.Flags().BoolVarP(&privacy, "private", "a", false, "The environment privacy")

	if err := CreateCmd.MarkFlagRequired("name"); err != nil {
		fmt.Println(err)
	}
	if err := CreateCmd.MarkFlagRequired("roles"); err != nil {
		fmt.Println(err)
	}
	if err := CreateCmd.MarkFlagRequired("main-role"); err != nil {
		fmt.Println(err)
	}
	if err := CreateCmd.MarkFlagRequired("description"); err != nil {
		fmt.Println(err)
	}
}
