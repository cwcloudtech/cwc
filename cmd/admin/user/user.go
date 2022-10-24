/*
Copyright © 2022 comwork.io contact.comwork.io

*/
package user

import (
	"cwc/cmd/admin/user/ls"

	"github.com/spf13/cobra"
)

// providerCmd represents the provider command
var UserCmd = &cobra.Command{
	Use:   "user",
	Short: "Get informations about available environments that you can associate to a virtual machine",
	Long:  `Get informations about available environments that you can associate to a virtual machine`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	UserCmd.DisableFlagsInUseLine = true
	UserCmd.AddCommand(ls.LsCmd)
}
