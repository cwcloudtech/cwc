/*
Copyright © 2022 comwork.io contact@comwork.io
*/
package environment

import (
	"cwc/cmd/admin/environment/create"
	"cwc/cmd/admin/environment/delete"
	"cwc/cmd/admin/environment/ls"

	"github.com/spf13/cobra"
)

// providerCmd represents the provider command
var EnvironmentCmd = &cobra.Command{
	Use:   "environment",
	Short: "Handle environments",
	Long:  `Handle environments`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	EnvironmentCmd.DisableFlagsInUseLine = true
	EnvironmentCmd.AddCommand(create.CreateCmd)
	EnvironmentCmd.AddCommand(delete.DeleteCmd)
	EnvironmentCmd.AddCommand(ls.LsCmd)

}
