package ls

import (
	"cwc/client"
	"cwc/handlers/user"
	"cwc/utils"

	"github.com/spf13/cobra"
)

var (
	instanceId string
	pretty     bool
)

// lsCmd represents the ls command
var LsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List availble instances in the cloud",
	Long: `This command lets you list your available instances in the cloud
This command takes no arguments`,
	Run: func(cmd *cobra.Command, args []string) {
		c, err := client.NewClient()
		utils.ExitIfError(err)
		if utils.IsBlank(instanceId) {
			instances, err := c.GetAllInstances()
			utils.ExitIfError(err)
			user.HandleGetInstances(instances, &pretty)
		} else {
			instance, err := c.GetInstance(*&instanceId)
			utils.ExitIfError(err)
			user.HandleGetInstance(instance, &pretty)
		}
	},
}

func init() {
	LsCmd.Flags().StringVarP(&instanceId, "instance", "i", "", "The instance id")
	LsCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "Pretty print the output (optional)")
}
