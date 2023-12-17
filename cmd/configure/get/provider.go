package get

import (
	"cwc/handlers/user"

	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var GetProviderCmd = &cobra.Command{
	Use:   "provider",
	Short: "Get the default provider",
	Long:  `This command lets you retrieve the default provider`,
	Run: func(cmd *cobra.Command, args []string) {
		user.HandlerGetDefaultProvider()
	},
}

func init() {
}
