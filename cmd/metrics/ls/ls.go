package ls

import (
	"cwc/client"
	"cwc/handlers/user"
	"cwc/utils"

	"github.com/spf13/cobra"
)

var (
	pretty bool
)

var LsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all metrics",
	Long: `This command lets you list all available metrics with their labels and values.
This command takes no arguments`,
	Run: func(cmd *cobra.Command, args []string) {
		samples, err := client.GetMetrics()
		utils.ExitIfError(err)
		user.HandleListMetrics(samples, &pretty)
	},
}

func init() {
	LsCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "Pretty print the output as a table (optional)")
}
