package ls

import (
	"cwc/client"
	"cwc/handlers/user"
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
		c, err := client.NewClient()
		utils.ExitIfError(err)
		if utils.IsBlank(functionId) {
			functions, err := c.GetAllFunctions()
			utils.ExitIfError(err)
			user.HandleGetFunctions(functions, &pretty)
		} else {
			function, err := c.GetFunctionById(*&functionId)
			utils.ExitIfError(err)
			user.HandleGetFunction(function, &pretty)
		}
	},
}

func init() {
	LsCmd.Flags().StringVarP(&functionId, "id", "f", "", "The function id")
	LsCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "Pretty print the output (optional)")
}
