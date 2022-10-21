/*
Copyright © 2022 comwork.io contact.comwork.io

*/
package update

import (
	"cwc/handlers/user"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	bucketId string
)

// updateCmd represents the update command
var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a particular s3 bucket credentials",
	Long: `This command lets you update a particular bucket credentials (access_key, secret_key of the bucket)
To use this command you have to provide the bucket ID`,
	Run: func(cmd *cobra.Command, args []string) {
		user.HandleUpdateBucket(&bucketId)
	},
}

func init() {

	UpdateCmd.Flags().StringVarP(&bucketId, "bucket", "b", "", "The bucket id")

	if err := UpdateCmd.MarkFlagRequired("bucket"); err != nil {
		fmt.Println(err)
	}

}
