/*
Copyright © 2022 comwork.io contact.comwork.io

*/
package set

import (
	"cwc/handlers/user"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var SetFormatCmd = &cobra.Command{
	Use:   "format",
	Short: "Set the default format",
	Long:  `This command lets you update the default format`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("cwc: you have to provide a value")
			os.Exit(1)
		}
		value := args[0]
		user.HandlerSetDefaultFormat(value)

	},
}

func init() {

}
