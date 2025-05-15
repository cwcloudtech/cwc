package delete

import (
	"cwc/handlers/admin"
	"cwc/utils"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	key    string
	userId string
	pretty bool = false
)

var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a user's key-value entry from storage",
	Long:  "This command deletes a user's key-value entry from storage by its key as admin.",
	Run: func(cmd *cobra.Command, args []string) {
		err := admin.HandleDeleteUserStorageKV(userId, key, &pretty)
		utils.ExitIfError(err)
	},
}

func init() {
	DeleteCmd.Flags().StringVarP(&key, "key", "k", "", "Key of the storage entry to delete")
	DeleteCmd.Flags().StringVarP(&userId, "user-id", "u", "", "ID of the user who owns the key")
	DeleteCmd.Flags().BoolVar(&pretty, "pretty", false, "Pretty print the output (optional)")

	err := DeleteCmd.MarkFlagRequired("key")
	if nil != err {
		fmt.Println(err)
	}

	err = DeleteCmd.MarkFlagRequired("user-id")
	if nil != err {
		fmt.Println(err)
	}
}
