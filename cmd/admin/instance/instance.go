/*
Copyright © 2022 comwork.io contact@comwork.io
*/
package instance

import (
	"cwc/cmd/admin/instance/create"
	"cwc/cmd/admin/instance/delete"
	"cwc/cmd/admin/instance/ls"
	"cwc/cmd/admin/instance/update"
	"cwc/cmd/admin/instance/refresh"

	"github.com/spf13/cobra"
)

// instanceCmd represents the instance command
var InstanceCmd = &cobra.Command{
	Use:   "instance",
	Short: "Manage your virtual machines in the cloud",
	Long: `This command lets you manage your virtual machines in the cloud.
Several actions are associated with this command such as creating an instance, updating an instance, deleting an instance
and listing your available instance`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	InstanceCmd.DisableFlagsInUseLine = true
	InstanceCmd.AddCommand(ls.LsCmd)
	InstanceCmd.AddCommand(update.UpdateCmd)
	InstanceCmd.AddCommand(create.CreateCmd)
	InstanceCmd.AddCommand(delete.DeleteCmd)
	InstanceCmd.AddCommand(refresh.RefreshCmd)

}
