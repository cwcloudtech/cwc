package delete

import (
	"cwc/handlers/admin"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	deviceId string
)

var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a particular device",
	Long: `This command lets you delete a particular device. To use this command you have to provide the device ID that you want to delete.`,
	Run: func(cmd *cobra.Command, args []string) {
		admin.HandleDeleteDevice(&deviceId)
	},
}

func init() {
	DeleteCmd.Flags().StringVarP(&deviceId, "deviceId", "d", "", "The device id")

	err := DeleteCmd.MarkFlagRequired("deviceId")
	if nil != err {
		fmt.Println(err)
	}
}
