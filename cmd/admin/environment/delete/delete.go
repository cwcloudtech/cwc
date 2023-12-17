package delete

import (
	"cwc/handlers/admin"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	envId string
)

// deleteCmd represents the delete command
var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a particular environment",
	Long: `This command lets you delete a particular environment.
To use this command you have to provide the environment ID that you want to delete`,
	Run: func(cmd *cobra.Command, args []string) {
		admin.HandleDeleteEnvironment(&envId)
	},
}

func init() {
	DeleteCmd.Flags().StringVarP(&envId, "environment", "e", "", "The environment id")

	err := DeleteCmd.MarkFlagRequired("environment")
	if nil != err {
		fmt.Println(err)
	}
}
