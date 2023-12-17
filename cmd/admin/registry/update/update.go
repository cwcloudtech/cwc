package update

import (
	"cwc/handlers/admin"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	registryId string
	email      string
)

// updateCmd represents the update command
var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a particular registry credentials",
	Long: `This command lets you update a particular bucket credentials (access_key, secret_key of the registry)
To use this command you have to provide the registry ID`,
	Run: func(cmd *cobra.Command, args []string) {
		admin.HandleUpdateRegistry(&registryId, &email)
	},
}

func init() {
	UpdateCmd.Flags().StringVarP(&registryId, "registry", "r", "", "The registry id")
	UpdateCmd.Flags().StringVarP(&email, "transfer registry", "t", "", "Transfer the registry to another user by his email")

	err := UpdateCmd.MarkFlagRequired("registry")
	if nil != err {
		fmt.Println(err)
	}
}
