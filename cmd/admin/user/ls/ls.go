package ls

import (
	"cwc/handlers/admin"
	"cwc/utils"

	"github.com/spf13/cobra"
)

var (
	userId string
	pretty bool
)

// lsCmd represents the ls command
var LsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List available environments",
	Long: `This command lets you list the available environment in the cloud that can be associeted to an instance
This command takes no arguments`,
	Run: func(cmd *cobra.Command, args []string) {
		if utils.IsBlank(userId) {
			admin.HandleGetUsers(&pretty)
		} else {
			admin.HandleGetUser(&userId, &pretty)
		}
	},
}

func init() {
	LsCmd.Flags().StringVarP(&userId, "user", "u", "", "The user id")
	LsCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "Pretty print the output (optional)")
}
