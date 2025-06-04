package kv

import (
	"cwc/cmd/storage/kv/create"
	"cwc/cmd/storage/kv/delete"
	"cwc/cmd/storage/kv/ls"
	"cwc/cmd/storage/kv/update"

	"github.com/spf13/cobra"
)

var KVCmd = &cobra.Command{
	Use:   "kv",
	Short: "Manage your key-value storage",
	Long: `This command lets you manage your key-value storage.
Several actions are associated with this command such as creating, updating,
retrieving and deleting key-value pairs in your storage.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	KVCmd.DisableFlagsInUseLine = true
	KVCmd.AddCommand(create.CreateCmd)
	KVCmd.AddCommand(ls.LsCmd)
	KVCmd.AddCommand(update.UpdateCmd)
	KVCmd.AddCommand(delete.DeleteCmd)
}
