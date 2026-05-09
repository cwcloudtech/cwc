package set

import (
	"cwc/config"
	"cwc/utils"
	"fmt"

	"github.com/spf13/cobra"
)

var SetKeyCmd = &cobra.Command{
	Use:   "key <keyname> <value>",
	Short: "Set an arbitrary configuration key",
	Long:  `This command lets you update an arbitrary key in the default configuration file`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.ExitIfNeeded("You have to provide a key and a value", len(args) < 2)

		keyName := args[0]
		value := args[1]
		config.UpdateFileKeyValue("config", keyName, value)
		fmt.Printf("%s = %v\n", keyName, value)
	},
}

func init() {
}
