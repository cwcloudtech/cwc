package bucket

import (
	"cwc/cmd/bucket/delete"
	"cwc/cmd/bucket/ls"
	"cwc/cmd/bucket/renew"

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
	BucketCmd.AddCommand(delete.DeleteCmd)
	BucketCmd.AddCommand(renew.RenewCmd)
}
