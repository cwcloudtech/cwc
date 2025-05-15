package ls

import (
	"cwc/handlers/admin"
	"cwc/utils"

	"github.com/spf13/cobra"
)

var (
	search     string
	userId     string
	startIndex int  = 0
	maxResults int  = 20
	pretty     bool = false
)

var LsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List key-value entries in storage as admin",
	Long:  "This command lists key-value entries in storage as admin. If --user-id is provided, it will list only entries for that specific user, otherwise it lists entries for all users.",
	Run: func(cmd *cobra.Command, args []string) {
		err := admin.HandleListStorageKVs(search, userId, startIndex, maxResults, &pretty)
		utils.ExitIfError(err)
	},
}

func init() {
	LsCmd.Flags().StringVarP(&search, "search", "s", "", "Search term to find in storage keys (optional)")
	LsCmd.Flags().StringVarP(&userId, "user-id", "u", "", "Filter by user ID (optional)")
	LsCmd.Flags().IntVarP(&startIndex, "start-index", "i", 0, "Start index for pagination (optional)")
	LsCmd.Flags().IntVarP(&maxResults, "max-results", "m", 20, "Maximum number of results to return (optional)")
	LsCmd.Flags().BoolVar(&pretty, "pretty", false, "Pretty print the output (optional)")
}
