/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package get

import (
	"cwc/handlers"

	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var GetEndpointCmd = &cobra.Command{
	Use:   "endpoint",
	Short: "Get the default endpoint",
	Long: `This command lets you retrieve the default endpoint`,
	Run: func(cmd *cobra.Command, args []string) {
			handlers.HandlerGetDefaultEndpoint()

	},
}

func init() {

}
