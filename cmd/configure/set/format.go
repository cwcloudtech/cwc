/*
Copyright © 2022 comwork.io contact@comwork.io
*/
package set

import (
	"cwc/handlers/user"
	"cwc/utils"

	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var SetFormatCmd = &cobra.Command{
	Use:   "format",
	Short: "Set the default format",
	Long:  `This command lets you update the default format`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.ExitIfNeeded("You have to provide a value", len(args) == 0)

		value := args[0]
		user.HandlerSetDefaultFormat(value)
	},
}

func init() {
}
