package prompt

import (
	"cwc/handlers/user"

	"github.com/spf13/cobra"
)

var (
	startIndex    *int
	maxResults    *int
	historyPretty bool = false
)

var HistoryCmd = &cobra.Command{
	Use:   "history",
	Short: "Get prompt history",
	Long:  `This command allows you to retrieve your prompt history in the form of lists`,
	Run: func(cmd *cobra.Command, args []string) {
		user.HandleGetPromptHistory(startIndex, maxResults, &historyPretty)
	},
}

func init() {
	startIdx := 0
	maxRes := 7
	startIndex = &startIdx
	maxResults = &maxRes

	HistoryCmd.Flags().IntVarP(startIndex, "start", "s", 0, "Start index for pagination (default: 0)")
	HistoryCmd.Flags().IntVarP(maxResults, "max", "m", 7, "Maximum results to return (default: 7)")
	HistoryCmd.Flags().BoolVarP(&historyPretty, "pretty", "p", false, "Pretty print the output (optional)")
}
