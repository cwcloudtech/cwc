package ls

import (
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
		if utils.IsBlank(bucketId) {
			user.HandleGetBuckets(&pretty)
		} else {
			user.HandleGetBucket(&bucketId, &pretty)
		}
	},
}

func init() {
	LsCmd.Flags().StringVarP(&bucketId, "bucket", "b", "", "The bucket id")
	LsCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "Pretty print the output (optional)")
}
