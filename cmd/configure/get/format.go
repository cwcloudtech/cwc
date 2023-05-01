/*
Copyright © 2022 comwork.io contact@comwork.io
*/
package get

import (
	"cwc/handlers/user"

	"github.com/spf13/cobra"
)

var GetFormatCmd = &cobra.Command{
	Use:   "format",
	Short: "Get the default format",
	Long:  `This command lets you retrieve the default format`,
	Run: func(cmd *cobra.Command, args []string) {
		user.HandlerGetDefaultFormat()
	},
}

func init() {

}
