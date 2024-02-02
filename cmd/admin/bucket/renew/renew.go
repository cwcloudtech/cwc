package renew

import (
	"cwc/handlers/admin"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	bucketId string
)

var RenewCmd = &cobra.Command{
	Use:   "renew",
	Short: "Renew a particular s3 bucket credentials",
	Long:  `This command lets you renew a particular s3 bucket credentials.`,
	Run: func(cmd *cobra.Command, args []string) {
		admin.HandleRenewBucketCredentials(&bucketId)
	},
}

func init() {
	RenewCmd.Flags().StringVarP(&bucketId, "bucket", "b", "", "The bucket id")

	err := RenewCmd.MarkFlagRequired("bucket")
	if nil != err {
		fmt.Println(err)
	}

}
