package ls

import (
	"cwc/handlers/admin"

	"github.com/spf13/cobra"
)

var (
	invocationId string
	pretty       bool = false
)

var LsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List available invocations",
	Long: `This command lets you list your available invocations in the cloud
This command takes no arguments`,
	Run: func(cmd *cobra.Command, args []string) {
		if *&invocationId == "" {
			admin.HandleGetInvocations(&pretty)
		} else {
			admin.HandleGetInvocationInvoker(&invocationId, &pretty)
		}
	},
}

func init() {
	LsCmd.Flags().StringVarP(&invocationId, "invocation", "i", "", "The invocation id")
	LsCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "Pretty print the output (optional)")
}
