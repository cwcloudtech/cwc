package storage

import (
	"cwc/cmd/admin/storage/kv"

	"github.com/spf13/cobra"
)

var StorageCmd = &cobra.Command{
	Use:   "storage",
	Short: "Manage storage resources as an admin",
	Long: `This command lets you manage storage resources as an admin.
Several actions are associated with this command such as listing, retrieving and 
deleting key-value pairs for any user.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	StorageCmd.DisableFlagsInUseLine = true
	StorageCmd.AddCommand(kv.KVCmd)
}
