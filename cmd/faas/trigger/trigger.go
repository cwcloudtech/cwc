package trigger

import (
	"cwc/cmd/faas/trigger/ls"
	"cwc/cmd/faas/trigger/triggerKinds"
	"cwc/cmd/faas/trigger/create"
	"cwc/cmd/faas/trigger/delete"
	"cwc/cmd/faas/trigger/truncate"

	"github.com/spf13/cobra"
)

var TriggerCmd = &cobra.Command{
	Use:   "trigger",
	Short: "Manage your triggers in the cloud",
	Long: `This command lets you manage your triggers in the cloud.
Several actions are associated with this command such as update a trigger, deleting a trigger
and listing your available triggers`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	TriggerCmd.DisableFlagsInUseLine = true
	TriggerCmd.AddCommand(ls.LsCmd)
	TriggerCmd.AddCommand(triggerKinds.TriggerKindsCmd)
	TriggerCmd.AddCommand(delete.DeleteCmd)
	TriggerCmd.AddCommand(create.CreateCmd)
	TriggerCmd.AddCommand(truncate.TruncateCmd)
}