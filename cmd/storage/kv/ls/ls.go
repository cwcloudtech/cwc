package ls

import (
	"cwc/handlers/user"
	"cwc/utils"

	"github.com/spf13/cobra"
)

var (
	search     string
	startIndex int  = 0
	maxResults int  = 20
	pretty     bool = false
)

var LsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all your key-value entries in storage",
	Long:  "This command lists all your key-value entries in your cloud storage.",
	Run: func(cmd *cobra.Command, args []string) {
		err := user.HandleListStorageKVs(search, startIndex, maxResults, &pretty)
		utils.ExitIfError(err)
	},
}

func init() {
	LsCmd.Flags().StringVarP(&search, "search", "s", "", "Search term to find in storage keys (optional)")
	LsCmd.Flags().IntVarP(&startIndex, "start-index", "i", 0, "Start index for pagination (optional)")
	LsCmd.Flags().IntVarP(&maxResults, "max-results", "m", 20, "Maximum number of results to return (optional)")
	LsCmd.Flags().BoolVar(&pretty, "pretty", false, "Pretty print the output (optional)")
}
