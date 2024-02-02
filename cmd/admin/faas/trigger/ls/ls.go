package ls

import (
	adminClient "cwc/admin"
	"cwc/handlers/admin"
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
		c, err := adminClient.NewClient()
		utils.ExitIfError(err)
		if utils.IsBlank(triggerId) {
			triggers, err := c.GetAllTriggers()
			utils.ExitIfError(err)
			admin.HandleGetTriggers(triggers, &pretty)
		} else {
			owner, err := c.GetTriggerOwnerById(triggerId)
			utils.ExitIfError(err)
			admin.HandleGetTriggerOwner(owner, &pretty)
		}
	},
}

func init() {
	LsCmd.Flags().StringVarP(&triggerId, "trigger", "t", "", "The trigger id")
	LsCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "Pretty print the output (optional)")
}
