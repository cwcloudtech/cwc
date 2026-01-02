package monitor

import (
	"cwc/cmd/monitor/update"
	"cwc/cmd/monitor/create"
	"cwc/cmd/monitor/ls"
	"cwc/cmd/monitor/delete"

	"github.com/spf13/cobra"
)

// providerCmd represents the provider command
var MonitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "Manage your monitors in the cloud",
	Long:  `This command lets you Manage your monitors in the cloud.
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
