package device

import (
	"cwc/cmd/admin/iot/device/delete"
	"cwc/cmd/admin/iot/device/ls"

	"github.com/spf13/cobra"
)

var DeviceCmd = &cobra.Command{
	Use:   "device",
	Short: "Manage your devices in the cloud",
	Long: `This command lets you Manage your devices in the cloud.
Several actions are associated with this command such as creating a device`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	DeviceCmd.DisableFlagsInUseLine = true
	DeviceCmd.AddCommand(ls.LsCmd)
	DeviceCmd.AddCommand(delete.DeleteCmd)
}
