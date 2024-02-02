package user

import (
	"cwc/cmd/admin/user/delete"
	"cwc/cmd/admin/user/ls"

	"github.com/spf13/cobra"
)

// providerCmd represents the provider command
var UserCmd = &cobra.Command{
	Use:   "user",
	Short: "Manage users",
	Long:  `Manage all users available on the platform`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	UserCmd.DisableFlagsInUseLine = true
	UserCmd.AddCommand(ls.LsCmd)
	UserCmd.AddCommand(delete.DeleteCmd)
}
