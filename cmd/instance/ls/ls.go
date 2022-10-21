/*
Copyright © 2022 comwork.io contact.comwork.io

*/
package ls

import (
	"cwc/handlers/user"

	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var LsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List availble instances in the cloud",
	Long: `This command lets you list your available instances in the cloud
This command takes no arguments`,
	Run: func(cmd *cobra.Command, args []string) {
		user.HandleGetInstance()
	},
}

func init() {

}
