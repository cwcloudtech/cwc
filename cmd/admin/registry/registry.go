package registry

import (
	"cwc/cmd/admin/registry/create"
	"cwc/cmd/admin/registry/delete"
	"cwc/cmd/admin/registry/ls"
	"cwc/cmd/admin/registry/renew"
	"cwc/cmd/admin/registry/transfer"

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
	RegistryCmd.AddCommand(ls.LsCmd)
	RegistryCmd.AddCommand(create.CreateCmd)
	RegistryCmd.AddCommand(transfer.TransferCmd)
	RegistryCmd.AddCommand(renew.RenewCmd)
	RegistryCmd.AddCommand(delete.DeleteCmd)
}
