package renew

import (
	"cwc/handlers/admin"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	registryId string
)

var RenewCmd = &cobra.Command{
	Use:   "renew",
	Short: "Renew a particular registry credentials",
	Long:  `This command lets you renew a particular registry credentials.`,
	Run: func(cmd *cobra.Command, args []string) {
		admin.HandleRenewRegistryCredentials(&registryId)
	},
}

func init() {
	RenewCmd.Flags().StringVarP(&registryId, "registry", "r", "", "The registry id")

	err := RenewCmd.MarkFlagRequired("registry")
	if nil != err {
		fmt.Println(err)
	}
}
