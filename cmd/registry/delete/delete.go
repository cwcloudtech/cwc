package delete

import (
	"cwc/handlers/user"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	registryId string
)

// deleteCmd represents the delete command
var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a particular registry",
	Long: `This command lets you delete a particular registry.
To use this command you have to provide the registry ID that you want to delete`,
	Run: func(cmd *cobra.Command, args []string) {
		user.HandleDeleteRegistry(&registryId)
	},
}

func init() {
	DeleteCmd.Flags().StringVarP(&registryId, "registry", "r", "", "The registry id")

	err := DeleteCmd.MarkFlagRequired("registry")
	if nil != err {
		fmt.Println(err)
	}
}
