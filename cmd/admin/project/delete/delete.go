package delete

import (
	"cwc/handlers/admin"

	"github.com/spf13/cobra"
)

var (
	projectId   string
	projectName string
	projectUrl  string
)

// deleteCmd represents the delete command
var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a particular project",
	Long: `This command lets you delete a particular project.
To use this command you have to provide the project ID that you want to delete
NOTE: The project needs to be empty and doesnt hold any instances`,
	Run: func(cmd *cobra.Command, args []string) {
		admin.HandleDeleteProject(&projectId, &projectName, &projectUrl)
	},
}

func init() {
	DeleteCmd.Flags().StringVarP(&projectId, "id", "p", "", "The project id")
	DeleteCmd.Flags().StringVarP(&projectName, "name", "n", "", "The project name")
	DeleteCmd.Flags().StringVarP(&projectUrl, "url", "u", "", "The project url")
}
