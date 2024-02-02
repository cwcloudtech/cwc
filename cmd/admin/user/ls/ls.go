package ls

import (
	adminClient "cwc/admin"
	"cwc/handlers/admin"
	"cwc/utils"
	"github.com/spf13/cobra"
)

var (
	userId string
	pretty bool = false
)

// lsCmd represents the ls command
var LsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List available users",
	Long:  `List available users and their details`,
	Run: func(cmd *cobra.Command, args []string) {
		c, err := adminClient.NewClient()
		utils.ExitIfError(err)
		if utils.IsBlank(userId) {
			responseUsers, err := c.GetAllUsers()
			utils.ExitIfError(err)
			admin.HandleGetUsers(responseUsers, &pretty)
		} else {
			responseUser, err := c.GetUser(userId)
			utils.ExitIfError(err)
			admin.HandleGetUser(responseUser, &pretty)
		}
	},
}

func init() {
	LsCmd.Flags().StringVarP(&userId, "user", "u", "", "The user id")
	LsCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "Pretty print the output (optional)")
}
