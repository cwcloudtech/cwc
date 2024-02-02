package renew

import (
	"cwc/handlers/user"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	bucketId string
)

var RenewCmd = &cobra.Command{
	Use:   "renew",
	Short: "Renew a particular s3 bucket credentials",
	Long: `This command lets you renew a particular bucket credentials (access_key, secret_key of the bucket)
To use this command you have to provide the bucket ID`,
	Run: func(cmd *cobra.Command, args []string) {
		user.HandleRenewBucket(&bucketId)
	},
}

func init() {
	RenewCmd.Flags().StringVarP(&bucketId, "bucket", "b", "", "The bucket id")

	err := RenewCmd.MarkFlagRequired("bucket")
	if nil != err {
		fmt.Println(err)
	}
}
