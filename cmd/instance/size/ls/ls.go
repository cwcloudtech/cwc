/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package ls

import (
	"cwc/handlers"

	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var LsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List availble instances types",
	Long:  `List availble instances types`,
	Run: func(cmd *cobra.Command, args []string) {
		handlers.HandleListInstancesTypes()
	},
}

func init() {

}
