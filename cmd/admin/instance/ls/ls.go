package ls

import (
	adminClient "cwc/admin"
	"cwc/handlers/admin"
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
	Short: "List availble instances with cwcloud",
	Long: `This command lets you list your available instances with cwcloud
This command takes no arguments`,
	Run: func(cmd *cobra.Command, args []string) {
		c, err := adminClient.NewClient()
		utils.ExitIfError(err)
		if utils.IsBlank(instanceId) {
			instances, err := c.AdminGetAllInstances()
			utils.ExitIfError(err)
			admin.HandleGetInstances(instances, &pretty)
		} else {
			instance, err := c.GetInstance(instanceId)
			utils.ExitIfError(err)
			admin.HandleGetInstance(instance, &pretty)
		}
	},
}

func init() {
	LsCmd.Flags().StringVarP(&instanceId, "instance", "i", "", "The instance id")
	LsCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "Pretty print the output (optional)")
}
