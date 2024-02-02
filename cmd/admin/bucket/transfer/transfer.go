package transfer

import (
	"cwc/handlers/admin"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	bucketId string
	email    string
)

// updateCmd represents the update command
var TransferCmd = &cobra.Command{
	Use:   "transfer",
	Short: "Transfer a particular s3 bucket ownership",
	Long:  `This command lets you transfer a particular s3 bucket ownership to another user by his email.`,
	Run: func(cmd *cobra.Command, args []string) {
		admin.HandleTransferBucketOwnership(&bucketId, &email)
	},
}

func init() {
	TransferCmd.Flags().StringVarP(&bucketId, "bucket", "b", "", "The bucket id")
	TransferCmd.Flags().StringVarP(&email, "email", "e", "", "The email of the new owner")

	err := TransferCmd.MarkFlagRequired("bucket")
	if nil != err {
		fmt.Println(err)
	}

	err = TransferCmd.MarkFlagRequired("email")
	if nil != err {
		fmt.Println(err)
	}

}
