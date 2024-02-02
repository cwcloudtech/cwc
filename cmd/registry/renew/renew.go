package renew

import (
	"cwc/handlers/user"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	registryId string
)

// updateCmd represents the update command
var RenewCmd = &cobra.Command{
	Use:   "renew",
	Short: "Renew a particular registry credentials",
	Long: `This command lets you renew a particular bucket credentials (access_key, secret_key of the registry)
To use this command you have to provide the registry ID`,
	Run: func(cmd *cobra.Command, args []string) {
		user.HandleRenewRegistry(&registryId)
	},
}

func init() {
	RenewCmd.Flags().StringVarP(&registryId, "registry", "r", "", "The registry id")

	err := RenewCmd.MarkFlagRequired("registry")
	if nil != err {
		fmt.Println(err)
	}
}
