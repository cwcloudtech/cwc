/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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
	Short: "Delete a particular virtual machine",
	Long: `This command lets you delete a particular instance.
To use this command you have to provide the instance ID that you want to delete`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	TypeCmd.DisableFlagsInUseLine = true
	TypeCmd.AddCommand(ls.LsCmd)
}
