package create

import (
	"cwc/client"
	"cwc/handlers/user"
	"cwc/utils"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	data 	client.Data
)

var CreateCmd = &cobra.Command{
	Use:  "create",
	Short: "Create a data in the cloud",
	Long:  "This command lets you create a data in the cloud.",
	Run: func(cmd *cobra.Command, args []string) {
		created_data, err := user.PrepareAddData(&data)
		utils.ExitIfError(err)
		user.HandleAddData(created_data)
	},
}

func init() {
	CreateCmd.Flags().StringVarP(&data.Device_id, "device_id", "d", "", "Device id of the data")
	CreateCmd.Flags().StringVarP(&data.Content, "content", "c", "", "Content of the data")

	err := CreateCmd.MarkFlagRequired("device_id")
	if nil != err {
		fmt.Println(err)
	}

	err = CreateCmd.MarkFlagRequired("content")
	if nil != err {
		fmt.Println(err)
	}
}
