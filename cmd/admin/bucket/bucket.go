package bucket

import (
	"cwc/cmd/admin/bucket/create"
	"cwc/cmd/admin/bucket/delete"
	"cwc/cmd/admin/bucket/ls"
	"cwc/cmd/admin/bucket/update"

	"github.com/spf13/cobra"
)

// bucketCmd represents the bucket command
var BucketCmd = &cobra.Command{
	Use:   "bucket",
	Short: "Manage your S3 buckets in the cloud",
	Long: `This command lets you manage your S3 buckets in the cloud.
Several actions are associated with this command such as update a bucket, deleting a bucket
and listing your available buckets`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	BucketCmd.DisableFlagsInUseLine = true
	BucketCmd.AddCommand(ls.LsCmd)
	BucketCmd.AddCommand(create.CreateCmd)

	BucketCmd.AddCommand(update.UpdateCmd)
	BucketCmd.AddCommand(delete.DeleteCmd)
}
