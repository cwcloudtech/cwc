package monitor

import (
	"cwc/cmd/admin/monitor/create"
	"cwc/cmd/admin/monitor/delete"
	"cwc/cmd/admin/monitor/ls"
	"cwc/cmd/admin/monitor/update"

	"github.com/spf13/cobra"
)

var MonitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "Manage your monitors in the cloud",
	Long: `This command lets you Manage your monitors in the cloud.
Several actions are associated with this command such listing your available monitors`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	MonitorCmd.DisableFlagsInUseLine = true
	MonitorCmd.AddCommand(ls.LsCmd)
	MonitorCmd.AddCommand(create.CreateCmd)
	MonitorCmd.AddCommand(update.UpdateCmd)
	MonitorCmd.AddCommand(delete.DeleteCmd)
}
