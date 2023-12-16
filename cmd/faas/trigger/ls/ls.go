package ls

import (
	"cwc/handlers/user"

	"github.com/spf13/cobra"
)

var (
	triggerId string
	pretty bool = false
)

var LsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List available triggers",
	Long: `This command lets you list your available triggers in the cloud
This command takes no arguments`,
	Run: func(cmd *cobra.Command, args []string) {
		if *&triggerId == "" {
			user.HandleGetTriggers(&pretty)
		} else {
			user.HandleGetTrigger(&triggerId, &pretty)
		}
	},
}

func init() {
	LsCmd.Flags().StringVarP(&triggerId, "trigger", "t", "", "The trigger id")
	LsCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "Pretty print the output")
}