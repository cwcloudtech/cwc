package delete

import (
	"cwc/handlers/user"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	functionId string
)

var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a particular function",
	Long: `This command lets you delete a particular function.
To use this command you have to provide the function ID that you want to delete.`,

	Run: func(cmd *cobra.Command, args []string) {
		user.HandleDeleteFunction(&functionId)
	},
}

func init() {
	DeleteCmd.Flags().StringVarP(&functionId, "function", "f", "", "The function id")

	err := DeleteCmd.MarkFlagRequired("function")
	if nil != err {
		fmt.Println(err)
	}
}
