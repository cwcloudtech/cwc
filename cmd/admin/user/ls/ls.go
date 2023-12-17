package ls

import (
	"cwc/handlers/admin"

	"github.com/spf13/cobra"
)

var (
	userId string
)

// lsCmd represents the ls command
var LsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List available environments",
	Long: `This command lets you list the available environment in the cloud that can be associeted to an instance
This command takes no arguments`,
	Run: func(cmd *cobra.Command, args []string) {
		if *&userId == "" {

			admin.HandleGetUsers()
		} else {
			admin.HandleGetUser(&userId)
		}
	},
}

func init() {
	LsCmd.Flags().StringVarP(&userId, "user", "u", "", "The user id")
}
