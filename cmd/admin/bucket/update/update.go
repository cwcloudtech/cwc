/*
Copyright © 2022 comwork.io contact@comwork.io
*/
package update

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
var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a particular s3 bucket credentials",
	Long: `This command lets you update a particular bucket credentials (access_key, secret_key of the bucket)
To use this command you have to provide the bucket ID`,
	Run: func(cmd *cobra.Command, args []string) {
		admin.HandleUpdateBucket(&bucketId, &email)
	},
}

func init() {
	UpdateCmd.Flags().StringVarP(&bucketId, "bucket", "b", "", "The bucket id")
	UpdateCmd.Flags().StringVarP(&email, "transfer bucket", "t", "", "Transfer the bucket to another user by his email")

	err := UpdateCmd.MarkFlagRequired("bucket")
	if nil != err {
		fmt.Println(err)
	}
}
