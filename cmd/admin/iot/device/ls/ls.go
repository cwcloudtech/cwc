package ls

import (
	adminClient "cwc/admin"
	"cwc/handlers/admin"
	"cwc/utils"

	"github.com/spf13/cobra"
)

var (
	deviceId string
	pretty   bool = false
)

var LsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List available devices",
	Long: `This command lets you list your available devices in the cloud
This command takes no arguments`,
	Run: func(cmd *cobra.Command, args []string) {
		c, err := adminClient.NewClient()
		utils.ExitIfError(err)
		if utils.IsBlank(deviceId) {
			devices, err := c.GetAllDevices()
			utils.ExitIfError(err)
			admin.HandleGetDevices(devices, &pretty)
		}
	},
}


func init() {
	LsCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "Pretty print the output (optional)")
}
