/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package dnszones

import (
	"cwc/cmd/dnszones/ls"

	"github.com/spf13/cobra"
)

// providerCmd represents the provider command
var DnsZonesCmd = &cobra.Command{
	Use:   "dnszones",
	Short: "Get informations about Dns Zones in which you can create your resources such as virtual machines.",
	Long:  `Get informations about Dns Zones in which you can create your resources such as virtual machines.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	DnsZonesCmd.DisableFlagsInUseLine = true
	DnsZonesCmd.AddCommand(ls.LsCmd)
}
