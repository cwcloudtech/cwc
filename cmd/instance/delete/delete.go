package delete

import (
	"cwc/handlers/user"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	instanceId string
)

// deleteCmd represents the delete command
var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a particular virtual machine",
	Long: `This command lets you delete a particular instance.
To use this command you have to provide the instance ID that you want to delete`,
	Run: func(cmd *cobra.Command, args []string) {
		user.HandleDeleteInstance(&instanceId)
	},
}

func init() {
	DeleteCmd.Flags().StringVarP(&instanceId, "instance", "i", "", "The instance id")

	err := DeleteCmd.MarkFlagRequired("instance")
	if nil != err {
		fmt.Println(err)
	}
}
