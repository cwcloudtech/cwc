/*
Copyright © 2022 comwork.io contact@comwork.io
*/
package delete

import (
	"cwc/handlers/admin"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	userId string
)

// deleteCmd represents the delete command
var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a particular project",
	Long: `This command lets you delete a particular project.
To use this command you have to provide the project ID that you want to delete
NOTE: The project needs to be empty and doesnt hold any instances`,
	Run: func(cmd *cobra.Command, args []string) {
		admin.HandleDeleteUser(&userId)
	},
}

func init() {

	DeleteCmd.Flags().StringVarP(&userId, "user", "u", "", "The user id")
	if err := DeleteCmd.MarkFlagRequired("user"); err != nil {
		fmt.Println(err)
	}

}
