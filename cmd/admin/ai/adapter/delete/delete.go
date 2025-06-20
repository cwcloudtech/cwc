package delete

import (
	adminClient "cwc/admin"
	"cwc/handlers/admin"
	"cwc/utils"

	"github.com/spf13/cobra"
)

var (
	adapterId string
	pretty    bool
)

var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an AI adapter (admin)",
	Long:  `This command deletes an existing AI adapter.`,
	Run: func(cmd *cobra.Command, args []string) {
		c, err := adminClient.NewClient()
		utils.ExitIfError(err)

		response, err := c.DeleteAIAdapter(adapterId)
		utils.ExitIfError(err)

		admin.HandleDeleteAIAdapter(response, &pretty)
	},
}

func init() {
	DeleteCmd.Flags().StringVarP(&adapterId, "adapter-id", "a", "", "ID of the AI adapter to delete (required)")
	DeleteCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "Pretty print the output (optional)")

	DeleteCmd.MarkFlagRequired("adapter-id")
}
