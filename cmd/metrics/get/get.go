package get

import (
	"cwc/client"
	"cwc/handlers/user"
	"cwc/utils"

	"github.com/spf13/cobra"
)

var (
	pretty    bool
	valueOnly bool
	filter    string
)

var GetCmd = &cobra.Command{
	Use:   "get <metric_name>",
	Short: "Get all samples for a specific metric",
	Long: `This command lets you retrieve all samples (with labels and values) for a given metric name.
You can filter samples by label with --filter using either "label:value" or "label=value".
Example: cwc metrics get cpu_all --filter instance:node-1`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		samples, err := client.GetMetricByName(name, filter)
		utils.ExitIfError(err)
		user.HandleGetMetric(name, samples, &pretty, valueOnly)
	},
}

func init() {
	GetCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "Pretty print the output as a table (optional)")
	GetCmd.Flags().BoolVar(&valueOnly, "value", false, "Display only raw sample value(s), one per line")
	GetCmd.Flags().StringVar(&filter, "filter", "", "Filter samples by label, format: label:value or label=value")
}
