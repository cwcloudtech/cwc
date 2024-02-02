package ls

import (
	"cwc/client"
	"cwc/handlers/user"
	"cwc/utils"

	"github.com/spf13/cobra"
)

var (
	triggerId string
	pretty    bool = false
)

var LsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List available triggers",
	Long: `This command lets you list your available triggers in the cloud
This command takes no arguments`,
	Run: func(cmd *cobra.Command, args []string) {
		c, err := client.NewClient()
		utils.ExitIfError(err)
		if utils.IsBlank(triggerId) {
			triggers, err := c.GetAllTriggers()
			utils.ExitIfError(err)
			user.HandleGetTriggers(triggers, &pretty)
		} else {
			trigger, err := c.GetTriggerById(triggerId)
			utils.ExitIfError(err)
			user.HandleGetTrigger(trigger, &pretty)
		}
	},
}

func init() {
	LsCmd.Flags().StringVarP(&triggerId, "trigger", "t", "", "The trigger id")
	LsCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "Pretty print the output (optional)")
}
