/*
Copyright © 2022 comwork.io contact.comwork.io

*/
package project

import (
	"cwc/cmd/project/create"
	"cwc/cmd/project/delete"
	"cwc/cmd/project/ls"

	"github.com/spf13/cobra"
)

// projectCmd represents the project command
var ProjectCmd = &cobra.Command{
	Use:   "project",
	Short: "Manage your projects in the cloud",
	Long: `This command lets you manage your projects in the cloud.
Several actions are associated with this command such as creating a project, deleting a project
and listing your available project`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	ProjectCmd.AddCommand(ls.LsCmd)
	ProjectCmd.AddCommand(create.CreateCmd)
	ProjectCmd.AddCommand(delete.DeleteCmd)
}
