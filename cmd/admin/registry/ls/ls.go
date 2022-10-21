/*
Copyright © 2022 comwork.io contact.comwork.io

*/
package ls

import (
	"cwc/handlers/admin"

	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var LsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List available registries",
	Long: `This command lets you list your available registries in the cloud
This command takes no arguments`,
	Run: func(cmd *cobra.Command, args []string) {
		admin.HandleGetRegistry()
	},
}

func init() {

}
