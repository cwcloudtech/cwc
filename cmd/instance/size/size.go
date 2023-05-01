/*
Copyright © 2022 comwork.io contact@comwork.io
*/
package size

import (
	"cwc/cmd/instance/size/ls"

	"github.com/spf13/cobra"
)

var (
	instanceId string
)

// deleteCmd represents the delete command
var TypeCmd = &cobra.Command{
	Use:   "type",
	Short: "Get informations about available instances types (size)",
	Long:  "Get informations about available instances types (size)",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	TypeCmd.DisableFlagsInUseLine = true
	TypeCmd.AddCommand(ls.LsCmd)
}
