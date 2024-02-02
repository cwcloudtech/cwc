package ls

import (
	adminClient "cwc/admin"
	"cwc/handlers/admin"
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
		c, err := adminClient.NewClient()
		utils.ExitIfError(err)
		if utils.IsBlank(registryId) {
			registries, err := c.GetAllRegistries()
			utils.ExitIfError(err)
			admin.HandleGetRegistries(registries, &pretty)
		} else {
			registry, err := c.GetRegistry(registryId)
			utils.ExitIfError(err)
			admin.HandleGetRegistry(registry, &pretty)
		}
	},
}

func init() {
	LsCmd.Flags().StringVarP(&registryId, "registry", "r", "", "The registry id")
	LsCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "Pretty print the output (optional)")
}
