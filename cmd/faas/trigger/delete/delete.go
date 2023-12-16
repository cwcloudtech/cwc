package delete

import (
	"cwc/handlers/user"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	triggerId   string
)

var DeleteCmd = &cobra.Command{
	Use:  "delete",
	Short: "Delete a trigger",
	Long: `This command lets you delete a trigger.
To use this command you have to provide the trigger ID that you want to delete.`,

	Run: func(cmd *cobra.Command, args []string) {
		user.HandleDeleteTrigger(&triggerId)
	},

}

func init() {
	DeleteCmd.Flags().StringVarP(&triggerId, "trigger", "t", "", "The trigger id")

	if err := DeleteCmd.MarkFlagRequired("trigger"); err != nil {
		fmt.Println(err)
	}
}