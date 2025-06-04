package delete

import (
	"cwc/handlers/user"
	"cwc/utils"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	key    string
	pretty bool = false
)

var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a key-value entry from your storage",
	Long:  "This command deletes a key-value entry from your cloud storage by its key.",
	Run: func(cmd *cobra.Command, args []string) {
		err := user.HandleDeleteStorageKV(key, &pretty)
		utils.ExitIfError(err)
	},
}

func init() {
	DeleteCmd.Flags().StringVarP(&key, "key", "k", "", "Key of the storage entry to delete")
	DeleteCmd.Flags().BoolVar(&pretty, "pretty", false, "Pretty print the output (optional)")

	err := DeleteCmd.MarkFlagRequired("key")
	if nil != err {
		fmt.Println(err)
	}
}
