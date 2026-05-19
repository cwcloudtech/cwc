package metrics

import (
	"cwc/cmd/metrics/get"
	"cwc/cmd/metrics/ls"

	"github.com/spf13/cobra"
)

var MetricsCmd = &cobra.Command{
	Use:   "metrics",
	Short: "Get metrics from the CWCloud API",
	Long:  `Get metrics from the CWCloud API in Prometheus format`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	MetricsCmd.DisableFlagsInUseLine = true
	MetricsCmd.AddCommand(ls.LsCmd)
	MetricsCmd.AddCommand(get.GetCmd)
}
