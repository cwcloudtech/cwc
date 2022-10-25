/*
Copyright © 2022 comwork.io contact.comwork.io

*/
package ls

import (
	"cwc/handlers/user"

	"github.com/spf13/cobra"
)

var (
	bucketId string
)

// lsCmd represents the ls command
var LsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List available buckets",
	Long: `This command lets you list your available buckets in the cloud
This command takes no arguments`,
	Run: func(cmd *cobra.Command, args []string) {
		if *&bucketId == "" {

			user.HandleGetBuckets()
		} else {
			user.HandleGetBucket(&bucketId)
		}
	},
}

func init() {

	LsCmd.Flags().StringVarP(&bucketId, "bucket", "b", "", "The bucket id")

}
