package keys

import (
	"cwc/handlers/user"

	"github.com/spf13/cobra"
)

var (
	pretty bool
)

var KeysCmd = &cobra.Command{
	Use:   "keys",
	Short: "List available configuration keys",
	Long:  `This command lets you list all available configuration keys`,
	Run: func(cmd *cobra.Command, args []string) {
		user.HandleGetConfigKeys(&pretty)
	},
}

func init() {
	KeysCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "Pretty print the output (optional)")
}
