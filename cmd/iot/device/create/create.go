package create

import (
	"cwc/client"
	"cwc/handlers/user"
	"cwc/utils"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	device      client.Device
	pretty 	bool = false
)

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a device in the cloud",
	Long:  "This command lets you create a device in the cloud.",
	Run: func(cmd *cobra.Command, args []string) {
		created_device, err := user.PrepareAddDevice(&device)
		utils.ExitIfError(err)
		user.HandleAddDevice(created_device, &pretty)
	},
}

func init() {
	CreateCmd.Flags().StringVarP(&device.Username, "username", "u", "", "Username of the device (email)")
	CreateCmd.Flags().StringVarP(&device.Typeobject_id, "object_type_id", "o", "", "Object type id of the device")

	err := CreateCmd.MarkFlagRequired("username")
	if nil != err {
		fmt.Println(err)
	}
	err = CreateCmd.MarkFlagRequired("object_type_id")
	if nil != err {
		fmt.Println(err)
	}
}