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
)

var GetCmd = &cobra.Command{
	Use:   "get <metric_name>",
	Short: "Get all samples for a specific metric",
	Long: `This command lets you retrieve all samples (with labels and values) for a given metric name.
Example: cwc metrics get cpu_all`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		samples, err := client.GetMetricByName(name)
		utils.ExitIfError(err)
		user.HandleGetMetric(name, samples, &pretty, valueOnly)
	},
}

func init() {
	GetCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "Pretty print the output as a table (optional)")
	GetCmd.Flags().BoolVar(&valueOnly, "value", false, "Display only raw sample value(s), one per line")
}
