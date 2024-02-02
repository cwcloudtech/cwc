package ls

import (
	"cwc/handlers/admin"
	"cwc/utils"
	adminClient "cwc/admin"
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
		c, err := adminClient.NewClient()
		utils.ExitIfError(err)
		if utils.IsBlank(invocationId) {
			invocations, err := c.GetAllInvocations()
			utils.ExitIfError(err)
			admin.HandleGetInvocations(invocations,&pretty)
		} else {
			invoker, err := c.GetInvocationInvokerById(invocationId)
			utils.ExitIfError(err)
			admin.HandleGetInvocationInvoker(invoker, &pretty)
		}
	},
}

func init() {
	LsCmd.Flags().StringVarP(&invocationId, "invocation", "i", "", "The invocation id")
	LsCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "Pretty print the output (optional)")
}
