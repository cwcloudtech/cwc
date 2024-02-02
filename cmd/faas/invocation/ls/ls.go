package ls

import (
	"cwc/client"
	"cwc/handlers/user"
	"cwc/utils"

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
		c, err := client.NewClient()
		utils.ExitIfError(err)
		if utils.IsBlank(invocationId) {
			invocations, err := c.GetAllInvocations()
			utils.ExitIfError(err)
			user.HandleGetInvocations(invocations, &pretty)
		} else {
			invocation, err := c.GetInvocationById(invocationId)
			utils.ExitIfError(err)
			user.HandleGetInvocation(invocation, &pretty)
		}
	},
}

func init() {
	LsCmd.Flags().StringVarP(&invocationId, "invocation", "i", "", "The invocation id")
	LsCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "Pretty print the output (optional)")
}
