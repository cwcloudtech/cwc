/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package set

import (
	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var SetCmd = &cobra.Command{
	Use:   "set",
	Short: "Update your default configurations",
	Long: `This command lets you update your default configurations`,
	Run: func(cmd *cobra.Command, args []string) {

		cmd.Help()
	},
}

func init() {
	SetCmd.DisableFlagsInUseLine = true
	SetCmd.AddCommand(SetEndpointCmd)
	SetCmd.AddCommand(SetProviderCmd)
	SetCmd.AddCommand(SetRegionCmd)
}
