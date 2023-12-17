package ls

import (
	"cwc/handlers/user"
	"cwc/utils"

	"github.com/spf13/cobra"
)

var (
	instanceId string
)

// lsCmd represents the ls command
var LsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List availble instances in the cloud",
	Long: `This command lets you list your available instances in the cloud
This command takes no arguments`,
	Run: func(cmd *cobra.Command, args []string) {
		if utils.IsBlank(instanceId) {
			user.HandleGetInstances()
		} else {
			user.HandleGetInstance(&instanceId)
		}
	},
}

func init() {
	LsCmd.Flags().StringVarP(&instanceId, "instance", "i", "", "The instance id")
}
