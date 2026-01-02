package environment

import (
	"cwc/cmd/admin/kubernetes/environment/ls"

	"github.com/spf13/cobra"
)

var EnvironmentCmd = &cobra.Command{
	Use:   "environment",
	Short: "Manage your kubernetes environments with cwcloud",
	Long: `This command lets you manage your kubernetes environments with cwcloud.
Several actions are associated with this command such listing your available environments`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	EnvironmentCmd.DisableFlagsInUseLine = true
	EnvironmentCmd.AddCommand(ls.LsCmd)
}
