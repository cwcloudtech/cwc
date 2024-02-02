package ls

import (
	"cwc/client"
	"cwc/handlers/user"
	"cwc/utils"

	"github.com/spf13/cobra"
)

var (
	bucketId string
	pretty   bool
)

// lsCmd represents the ls command
var LsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List available buckets",
	Long: `This command lets you list your available buckets in the cloud
This command takes no arguments`,
	Run: func(cmd *cobra.Command, args []string) {
		c, err := client.NewClient()
		utils.ExitIfError(err)
		if utils.IsBlank(bucketId) {
			buckets, err := c.GetAllBuckets()
			utils.ExitIfError(err)
			user.HandleGetBuckets(buckets, &pretty)
		} else {
			bucket, err := c.GetBucket(*&bucketId)
			utils.ExitIfError(err)
			user.HandleGetBucket(bucket, &pretty)
		}
	},
}

func init() {
	LsCmd.Flags().StringVarP(&bucketId, "bucket", "b", "", "The bucket id")
	LsCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "Pretty print the output (optional)")
}
