/*
Copyright © 2022 comwork.io contact@comwork.io
*/
package ls

import (
	"cwc/handlers/user"

	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var LsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List available Dns Zones",
	Long: `This command lets you list the available Dns Zones in the cloud
This command takes no arguments`,
	Run: func(cmd *cobra.Command, args []string) {
		user.HandleListDnsZones()
	},
}

func init() {

}
