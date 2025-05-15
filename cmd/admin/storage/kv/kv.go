package kv

import (
	"cwc/cmd/admin/storage/kv/delete"
	"cwc/cmd/admin/storage/kv/ls"

	"github.com/spf13/cobra"
)

var KVCmd = &cobra.Command{
	Use:   "kv",
	Short: "Manage key-value storage as admin",
	Long: `This command lets you manage key-value storage as an admin.
Several actions are associated with this command such as listing, retrieving and 
deleting key-value pairs for any user.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	KVCmd.DisableFlagsInUseLine = true
	KVCmd.AddCommand(ls.LsCmd)
	KVCmd.AddCommand(delete.DeleteCmd)
}
