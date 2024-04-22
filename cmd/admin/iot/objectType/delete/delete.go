package delete

import (
	"cwc/handlers/admin"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	objectTypeId string
)

var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a particular object type",
	Long: `This command lets you delete a particular object type.
To use this command you have to provide the object type ID that you want to delete.`,
	Run: func(cmd *cobra.Command, args []string) {
		admin.HandleDeleteObjectType(&objectTypeId)
	},
}

func init() {
	DeleteCmd.Flags().StringVarP(&objectTypeId, "objectType", "o", "", "The object type id")

	err := DeleteCmd.MarkFlagRequired("objectType")
	if nil != err {
		fmt.Println(err)
	}
}
