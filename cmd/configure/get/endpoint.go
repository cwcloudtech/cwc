/*
Copyright © 2022 comwork.io contact@comwork.io
*/
package get

import (
	"cwc/handlers/user"

	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var GetEndpointCmd = &cobra.Command{
	Use:   "endpoint",
	Short: "Get the default endpoint",
	Long:  `This command lets you retrieve the default endpoint`,
	Run: func(cmd *cobra.Command, args []string) {
		user.HandlerGetDefaultEndpoint()
	},
}

func init() {

}
