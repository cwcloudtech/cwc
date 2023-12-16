package ls

import (
	"cwc/handlers/admin"

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
		if *&functionId == "" {
			admin.HandleGetFunctions(&pretty)
		} else {
			admin.HandleGetFunctionOwner(&functionId, &pretty)
		}
	},
}

func init() {
	LsCmd.Flags().StringVarP(&functionId, "id", "f", "", "The function id")
	LsCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "Pretty print the output (optional)")
}
