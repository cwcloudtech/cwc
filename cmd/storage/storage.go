package storage

import (
	"cwc/cmd/storage/kv"

	"github.com/spf13/cobra"
)

var StorageCmd = &cobra.Command{
	Use:   "storage",
	Short: "Manage your storage resources",
	Long: `This command lets you manage your storage resources with cwcloud.
Several actions are associated with this command such as creating, updating, 
retrieving and deleting key-value pairs in your storage.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	StorageCmd.DisableFlagsInUseLine = true
	StorageCmd.AddCommand(kv.KVCmd)
}
