package ls

import (
	adminClient "cwc/admin"
	"cwc/handlers/admin"
	"cwc/utils"
	"github.com/spf13/cobra"
)

var (
	monitorId string
	pretty    bool = false
)

var LsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List available monitors",
	Long: `This command lets you list your available monitors in the cloud
This command takes no arguments`,
	Run: func(cmd *cobra.Command, args []string) {
		c, err := adminClient.NewClient()
		utils.ExitIfError(err)
		if utils.IsBlank(monitorId) {
			monitors, err := c.GetAllMonitors()
			utils.ExitIfError(err)
			admin.HandleGetMonitors(monitors, &pretty)
		} else {
			monitor, err := c.GetMonitorById(monitorId)
			utils.ExitIfError(err)
			admin.HandleGetMonitor(monitor, &pretty)
		}
	},
}

func init() {
	LsCmd.Flags().StringVarP(&monitorId, "id", "m", "", "The monitor id")
	LsCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "Pretty print the output (optional)")
}
