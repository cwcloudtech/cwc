package kubernetes

import (
	"cwc/cmd/kubernetes/configure"
	"cwc/cmd/kubernetes/deployment"

	"github.com/spf13/cobra"
)

var KubernetesCmd = &cobra.Command{
	Use:   "kubernetes",
	Short: "Manage your kubernetes resources",
	Long: `This command lets you manage your kubernetes resources.
Several actions are associated with this command such as creating a deployment, listing deployments
and deleting a deployment`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	KubernetesCmd.DisableFlagsInUseLine = true
	KubernetesCmd.AddCommand(deployment.DeploymentCmd)
	KubernetesCmd.AddCommand(configure.ConfigureCmd)
}
