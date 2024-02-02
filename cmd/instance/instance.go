package instance

import (
	"cwc/cmd/instance/create"
	"cwc/cmd/instance/delete"
	"cwc/cmd/instance/ls"
	"cwc/cmd/instance/type"

	"cwc/cmd/instance/update"

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
	InstanceCmd.AddCommand(size.TypeCmd)

}
