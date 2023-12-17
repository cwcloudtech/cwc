package get

import (
	"cwc/handlers/user"

	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var GetRegionCmd = &cobra.Command{
	Use:   "region",
	Short: "Get the default region",
	Long:  `This command lets you retrieve the default region`,
	Run: func(cmd *cobra.Command, args []string) {
		user.HandlerGetDefaultRegion()
	},
}

func init() {
}
