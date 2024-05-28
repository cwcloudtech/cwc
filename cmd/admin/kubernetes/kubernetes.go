package kubernetes

import (
	"cwc/cmd/admin/kubernetes/cluster"
	"cwc/cmd/admin/kubernetes/environment"
	"github.com/spf13/cobra"
)

var KubernetesCmd = &cobra.Command{
	Use:   "kubernetes",
	Short: "Manage your kubernetes environments in the cloud",
	Long: `This command lets you Manage your kubernetes environments in the cloud.
Several actions are associated with this command such listing your available environments`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	KubernetesCmd.DisableFlagsInUseLine = true
	KubernetesCmd.AddCommand(environment.EnvironmentCmd)
	KubernetesCmd.AddCommand(cluster.ClusterCmd)
}
