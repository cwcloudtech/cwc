ackage registry

import (
	"cwc/cmd/registry/delete"
	"cwc/cmd/registry/ls"
	"cwc/cmd/registry/update"

	"github.com/spf13/cobra"
)

// RegistryCmd represents the bucket command
var RegistryCmd = &cobra.Command{
	Use:   "registry",
	Short: "Manage your registries in the cloud",
	Long: `This command lets you manage your registries in the cloud.
Several actions are associated with this command such as update a registry, deleting a registry
and listing your available registries`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	RegistryCmd.DisableFlagsInUseLine = true
	RegistryCmd.AddCommand(ls.LsCmd)
	RegistryCmd.AddCommand(update.UpdateCmd)
	RegistryCmd.AddCommand(delete.DeleteCmd)
}
