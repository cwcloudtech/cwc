package ls

import (
	"cwc/client"
	"cwc/handlers/user"
	"cwc/utils"

	"github.com/spf13/cobra"
)

var (
	formId string
	pretty bool = false
)

// lsCmd represents the ls command
var LsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List available contact forms",
	Long: `This command lets you list the available contact forms in the cloud
This command takes no arguments`,
	Run: func(cmd *cobra.Command, args []string) {
		c, err := client.NewClient()
		utils.ExitIfError(err)
		if utils.IsBlank(formId) {
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
