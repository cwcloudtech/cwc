/*
Copyright © 2022 comwork.io contact.comwork.io

*/
package update

import (
	"cwc/handlers/user"
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
		user.HandleUpdateInstance(&instanceId, &status)
	},
}

func init() {

	UpdateCmd.Flags().StringVarP(&instanceId, "instance", "i", "", "The instance id")
	UpdateCmd.Flags().StringVarP(&status, "status", "s", "", "Instance status (poweroff, poweron, reboot)")

	if err := UpdateCmd.MarkFlagRequired("instance"); err != nil {
		fmt.Println(err)
	}

	if err := UpdateCmd.MarkFlagRequired("status"); err != nil {
		fmt.Println(err)
	}
}
