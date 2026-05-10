package get

import (
	"cwc/handlers/user"

	"github.com/spf13/cobra"
)

var GetKeyCmd = &cobra.Command{
	Use:   "key <key>",
	Short: "Get a configuration value by key",
	Long:  `This command lets you retrieve any configuration value by its key`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		user.HandlerGetConfigKey(args[0])
	},
}

func init() {
}
