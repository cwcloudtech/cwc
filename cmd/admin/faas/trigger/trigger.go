package trigger

import (
	"cwc/cmd/admin/faas/trigger/ls"

	"github.com/spf13/cobra"
)

var TriggerCmd = &cobra.Command{
	Use:   "trigger",
	Short: "Manage your triggers in the cloud",
	Long: `This command lets you manage your triggers in the cloud.
Several actions are associated with this command such as listing your available triggers`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	TriggerCmd.DisableFlagsInUseLine = true
	TriggerCmd.AddCommand(ls.LsCmd)
}