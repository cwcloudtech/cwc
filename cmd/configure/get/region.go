/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package get

import (
	"cwc/handlers"

	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var GetRegionCmd = &cobra.Command{
	Use:   "region",
	Short: "Get the default region",
	Long: `This command lets you retrieve the default region`,
	Run: func(cmd *cobra.Command, args []string) {
			handlers.HandlerGetDefaultRegion()

	},
}

func init() {

}
