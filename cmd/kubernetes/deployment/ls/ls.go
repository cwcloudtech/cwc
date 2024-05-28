package ls

import (
	"cwc/client"
	"cwc/handlers/user"
	"cwc/utils"
	"github.com/spf13/cobra"
)

var deploymentId string
var pretty bool

var LsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List available deployments",
	Long:  `This command lets you list your available deployments in the cloud`,
	Run: func(cmd *cobra.Command, args []string) {
		c, err := client.NewClient()
		utils.ExitIfError(err)
		if utils.IsBlank(deploymentId) {
			deployments, err := c.GetAllDeployments()
			utils.ExitIfError(err)
			user.HandleGetDeployments(deployments, &pretty)
		} else {
			deployment, err := c.GetDeploymentById(*&deploymentId)
			utils.ExitIfError(err)
			user.HandleGetDeployment(deployment, &pretty)
		}
	},
}

func init() {
	LsCmd.Flags().StringVarP(&deploymentId, "id", "d", "", "The deployment id")
	LsCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "Pretty print the output (optional)")
}
