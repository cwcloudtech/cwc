/*
Copyright © 2022 comwork.io contact.comwork.io

*/
package environment

import (
	"cwc/cmd/environment/ls"

	"github.com/spf13/cobra"
)

// providerCmd represents the provider command
var EnvironmentCmd = &cobra.Command{
	Use:   "environment",
	Short: "Get informations about available environments that you can associate to a virtual machine",
	Long:  `Get informations about available environments that you can associate to a virtual machine`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	EnvironmentCmd.DisableFlagsInUseLine = true
	EnvironmentCmd.AddCommand(ls.LsCmd)
}
