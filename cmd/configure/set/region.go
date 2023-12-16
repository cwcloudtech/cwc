/*
Copyright © 2022 comwork.io contact@comwork.io
*/
package set

import (
	"cwc/handlers/user"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var SetRegionCmd = &cobra.Command{
	Use:   "region",
	Short: "Set the default endpoint",
	Long:  `This command lets you update the default endpoint`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			fmt.Println("cwc: you have to provide a value")
			os.Exit(1)
		}
		value := args[0]
		user.HandlerSetDefaultRegion(value)
	},
}

func init() {
}
