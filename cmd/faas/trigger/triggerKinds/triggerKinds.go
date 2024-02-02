package triggerKinds

import (
	"cwc/client"
	"cwc/handlers/user"
	"cwc/utils"

	"github.com/spf13/cobra"
)

var pretty bool = false

var TriggerKindsCmd = &cobra.Command{
	Use:   "kinds",
	Short: "List available trigger kinds",
	Long:  `This command lets you list your available trigger kinds in the cloud`,
	Run: func(cmd *cobra.Command, args []string) {
		triggerKinds, err := client.GetTriggerKinds()
		utils.ExitIfError(err)
		user.HandleGetTriggerKinds(triggerKinds, &pretty)
	},
}

func init() {
	TriggerKindsCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "Pretty print the output (optional)")
}
