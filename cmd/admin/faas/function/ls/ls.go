package ls

import (
	adminClient "cwc/admin"
	"cwc/handlers/admin"
	"cwc/utils"
	"github.com/spf13/cobra"
)

var (
	functionId string
	pretty     bool = false
)

var LsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List available functions",
	Long: `This command lets you list your available functions in the cloud
This command takes no arguments`,

	Run: func(cmd *cobra.Command, args []string) {
		c, err := adminClient.NewClient()
		utils.ExitIfError(err)
		if utils.IsBlank(functionId) {
			functions, err := c.GetAllFunctions()
			utils.ExitIfError(err)
			admin.HandleGetFunctions(functions, &pretty)
		} else {
			owner, err := c.GetFunctionOwnerById(functionId)
			utils.ExitIfError(err)
			admin.HandleGetFunctionOwner(owner, &pretty)
		}
	},
}

func init() {
	LsCmd.Flags().StringVarP(&functionId, "id", "f", "", "The function id")
	LsCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "Pretty print the output (optional)")
}
