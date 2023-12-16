package delete

import (
	"cwc/handlers/user"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	invocationId string
)

var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a particular invocation",
	Long: `This command lets you delete a particular invocation.
To use this command you have to provide the invocation ID that you want to delete.`,
	Run: func(cmd *cobra.Command, args []string) {
		user.HandleDeleteInvocation(&invocationId)
	},
}

func init() {
	DeleteCmd.Flags().StringVarP(&invocationId, "invocation", "i", "", "The invocation id")

	err := DeleteCmd.MarkFlagRequired("invocation")
	if nil != err {
		fmt.Println(err)
	}
}
