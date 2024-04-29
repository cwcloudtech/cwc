package ls

import (
	"cwc/client"
	"cwc/handlers/user"
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
		c, err := client.NewClient()
		utils.ExitIfError(err)
		if utils.IsBlank(deviceId) {
			devices, err := c.GetAllDevices()
			utils.ExitIfError(err)
			user.HandleGetDevices(devices, &pretty)
		} else {
			device, err := c.GetDeviceById(deviceId)
			utils.ExitIfError(err)
			user.HandleGetDevice(device, &pretty)
		}
	},
}


func init() {
	LsCmd.Flags().StringVarP(&deviceId, "id", "i", "", "The device id")
	LsCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "Pretty print the output (optional)")
}
