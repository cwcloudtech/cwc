/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package delete

import (
	"cwc/handlers"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	projectId string
)

// deleteCmd represents the delete command
var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a particular project",
	Long: `This command lets you delete a particular project.
To use this command you have to provide the project ID that you want to delete
NOTE: The project needs to be empty and doesnt hold any instances`,
	Run: func(cmd *cobra.Command, args []string) {
		handlers.HandleDeleteProject(&projectId)
	},
}

func init() {

	DeleteCmd.Flags().StringVarP(&projectId, "project_id", "p", "", "The project id")

	if err := DeleteCmd.MarkFlagRequired("project_id"); err != nil {
		fmt.Println(err)
	}
}
