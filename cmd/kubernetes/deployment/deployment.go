package deployment

import (
	"cwc/cmd/kubernetes/deployment/create"
	"cwc/cmd/kubernetes/deployment/ls"

	"github.com/spf13/cobra"
)

var DeploymentCmd = &cobra.Command{
	Use:   "deployment",
	Short: "Manage your kubernetes deployments",
	Long: `This command lets you manage your kubernetes deployments.
Several actions are associated with this command such as creating a deployment, listing deployments
and deleting a deployment`,

	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	DeploymentCmd.DisableFlagsInUseLine = true
	DeploymentCmd.AddCommand(ls.LsCmd)
	DeploymentCmd.AddCommand(create.CreateCmd)
}
