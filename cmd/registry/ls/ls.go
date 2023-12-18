package ls

import (
	"cwc/handlers/user"
	"cwc/utils"

	"github.com/spf13/cobra"
)

var (
	registryId string
	pretty     bool
)

// lsCmd represents the ls command
var LsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List available registries",
	Long: `This command lets you list your available registries in the cloud
This command takes no arguments`,
	Run: func(cmd *cobra.Command, args []string) {
		if utils.IsBlank(registryId) {
			user.HandleGetRegistries()
		} else {
			user.HandleGetRegistry(&registryId, &pretty)
		}
	},
}

func init() {
	LsCmd.Flags().StringVarP(&registryId, "registry", "r", "", "The registry id")
	LsCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "Pretty print the output (optional)")
}
