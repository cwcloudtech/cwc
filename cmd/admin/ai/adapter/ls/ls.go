package ls

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

var LsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all available external AI adapters",
	Long:  `This command lists all available external AI adapters in the system. Use --adapter-id to get details of a specific adapter.`,
	Run: func(cmd *cobra.Command, args []string) {
		c, err := adminClient.NewClient()
		utils.ExitIfError(err)

		if adapterId != "" {
			adapter, err := c.GetAIAdapterById(adapterId)
			utils.ExitIfError(err)
			admin.HandleGetAIAdapter(adapter, &pretty)
			return
		}

		adapters, err := c.GetAllAIAdapters()
		utils.ExitIfError(err)
		admin.HandleGetAIAdapters(adapters, &pretty)
	},
}

func init() {
	LsCmd.Flags().StringVarP(&adapterId, "adapter-id", "a", "", "ID of a specific AI adapter to get details (optional)")
	LsCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "Pretty print the output (optional)")
}
