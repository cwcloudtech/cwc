package transfer

import (
	"cwc/handlers/admin"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	registryId string
	email      string
)

var TransferCmd = &cobra.Command{
	Use:   "transfer",
	Short: "Transfer a particular registry ownership",
	Long:  `This command lets you transfer a particular registry ownership to another user by his email.`,
	Run: func(cmd *cobra.Command, args []string) {
		admin.HandleTransferRegistryOwnership(&registryId, &email)
	},
}

func init() {
	TransferCmd.Flags().StringVarP(&registryId, "registry", "r", "", "The registry id")
	TransferCmd.Flags().StringVarP(&email, "email", "e", "", "The email of the new owner")

	err := TransferCmd.MarkFlagRequired("registry")
	if nil != err {
		fmt.Println(err)
	}

	err = TransferCmd.MarkFlagRequired("email")
	if nil != err {
		fmt.Println(err)
	}
}
