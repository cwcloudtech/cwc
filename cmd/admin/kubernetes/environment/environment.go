package environment

import (
	"cwc/cmd/admin/kubernetes/environment/ls"

	"github.com/spf13/cobra"
)

var EnvironmentCmd = &cobra.Command{
	Use:   "environment",
	Short: "Manage your kubernetes environments in the cloud",
	Long: `This command lets you Manage your kubernetes environments in the cloud.
Several actions are associated with this command such listing your available environments`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	EnvironmentCmd.DisableFlagsInUseLine = true
	EnvironmentCmd.AddCommand(ls.LsCmd)
}
