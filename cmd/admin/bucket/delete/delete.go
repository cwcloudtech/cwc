/*
Copyright © 2022 comwork.io contact@comwork.io
*/
package delete

import (
	"cwc/handlers/admin"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	bucketId string
)

// deleteCmd represents the delete command
var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a particular S3 bucket",
	Long: `This command lets you delete a particular S3 bucket.
To use this command you have to provide the bucket ID that you want to delete`,
	Run: func(cmd *cobra.Command, args []string) {
		admin.HandleDeleteBucket(&bucketId)
	},
}

func init() {

	DeleteCmd.Flags().StringVarP(&bucketId, "bucket", "b", "", "The bucket id")

	if err := DeleteCmd.MarkFlagRequired("bucket"); err != nil {
		fmt.Println(err)
	}
}
