package delete

import (
	"cwc/handlers/user"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	formId string
)

var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a particular contact form",
	Long: `This command lets you delete a particular contact form.
To use this command you have to provide the form ID that you want to delete.`,
	Run: func(cmd *cobra.Command, args []string) {
		user.HandleDeleteForm(&formId)
	},
}

func init() {
	DeleteCmd.Flags().StringVarP(&formId, "form", "f", "", "The contact form id")

	err := DeleteCmd.MarkFlagRequired("form")
	if nil != err {
		fmt.Println(err)
	}
}
