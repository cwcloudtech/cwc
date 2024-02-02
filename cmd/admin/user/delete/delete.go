package delete

import (
	"cwc/handlers/admin"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	userId string
)

// deleteCmd represents the delete command
var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a particular user",
	Long:  `Delete a particular user by providing the user id`,
	Run: func(cmd *cobra.Command, args []string) {
		admin.HandleDeleteUser(&userId)
	},
}

func init() {
	DeleteCmd.Flags().StringVarP(&userId, "user", "u", "", "The user id")

	err := DeleteCmd.MarkFlagRequired("user")
	if nil != err {
		fmt.Println(err)
	}
}
