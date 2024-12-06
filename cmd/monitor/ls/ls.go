package ls

import (
	"cwc/client"
	"cwc/handlers/user"
	"cwc/utils"

	"github.com/spf13/cobra"
)

var (
	monitorId string
	pretty    bool = false
)

// lsCmd represents the ls command
var LsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List available monitors",
	Long: `This command lets you list the available monitors in the cloud that can be associeted to an instance
This command takes no arguments`,
	Run: func(cmd *cobra.Command, args []string) {
		c, err := client.NewClient()
		utils.ExitIfError(err)
		if utils.IsBlank(monitorId) {
			monitors, err := c.GetAllMonitors()
			utils.ExitIfError(err)
			user.HandleGetMonitors(monitors, &pretty)
		} else {
			monitor, err := c.GetMonitorById(*&monitorId)
			utils.ExitIfError(err)
			user.HandleGetMonitor(monitor, &pretty)
		}
	},
}

func init() {
	LsCmd.Flags().StringVarP(&monitorId, "monitor", "m", "", "The monitor id")
	LsCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "Pretty print the output (optional)")
}
