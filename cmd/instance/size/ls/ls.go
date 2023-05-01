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
	Short: "List availble instances types",
	Long:  `List availble instances types`,
	Run: func(cmd *cobra.Command, args []string) {
		user.HandleListInstancesTypes()
	},
}

func init() {

}
