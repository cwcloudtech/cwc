/*
Copyright © 2022 comwork.io contact@comwork.io
*/
package ls

import (
	"cwc/handlers/admin"

	"github.com/spf13/cobra"
)

var (
	projectId   string
	projectName string
	projectUrl  string
)

// lsCmd represents the ls command
var LsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List available projects",
	Long: `This command lets you list your available projects in the cloud
This command takes no arguments`,
	Run: func(cmd *cobra.Command, args []string) {

		admin.HandleGetProjects(&projectId, &projectName, &projectUrl)
	},
}

func init() {

	LsCmd.Flags().StringVarP(&projectId, "id", "p", "", "The project id")
	LsCmd.Flags().StringVarP(&projectName, "name", "n", "", "The project name")
	LsCmd.Flags().StringVarP(&projectUrl, "url", "u", "", "The project url")

}
