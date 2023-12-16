/*
Copyright © 2022 comwork.io contact@comwork.io
*/
package update

import (
	"cwc/handlers/admin"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	instanceId string
	status     string
)

// updateCmd represents the update command
var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a particular instance status",
	Long: `This command lets you update a particular instance status such as poweroff, poweron, reboot
To use this command you have to provide the instance ID and the desired status`,
	Run: func(cmd *cobra.Command, args []string) {
		admin.HandleUpdateInstance(&instanceId, &status)
	},
}

func init() {
	UpdateCmd.Flags().StringVarP(&instanceId, "instance", "i", "", "The instance id")
	UpdateCmd.Flags().StringVarP(&status, "status", "s", "", "Instance status (poweroff, poweron, reboot)")

	err := UpdateCmd.MarkFlagRequired("instance")
	if nil != err {
		fmt.Println(err)
	}

	err = UpdateCmd.MarkFlagRequired("status")
	if nil != err {
		fmt.Println(err)
	}
}
