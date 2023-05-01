/*
Copyright © 2022 comwork.io contact@comwork.io
*/
package provider

import (
	"cwc/cmd/provider/ls"

	"github.com/spf13/cobra"
)

// providerCmd represents the provider command
var ProviderCmd = &cobra.Command{
	Use:   "provider",
	Short: "Get informations about available providers that you created cloud resources on them",
	Long:  `Get informations about available providers that you created cloud resources on them`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	ProviderCmd.DisableFlagsInUseLine = true
	ProviderCmd.AddCommand(ls.LsCmd)
}
