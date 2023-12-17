package ls

import (
	"cwc/handlers/admin"
	"cwc/utils"

	"github.com/spf13/cobra"
)

var (
	registryId string
)

// lsCmd represents the ls command
var LsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List available registries",
	Long: `This command lets you list your available registries in the cloud
This command takes no arguments`,
	Run: func(cmd *cobra.Command, args []string) {
		if utils.IsBlank(registryId) {
			admin.HandleGetRegistries()
		} else {
			admin.HandleGetRegistry(&registryId)
		}
	},
}

func init() {
	LsCmd.Flags().StringVarP(&registryId, "registry", "r", "", "The registry id")
}
