package restart

import (
	"cwc/handlers/user"
	"fmt"

	"github.com/spf13/cobra"
)

var instanceId string

// RestartCmd represents the restart command
var RestartCmd = &cobra.Command{
	Use:     "restart",
	Aliases: []string{"reboot"},
	Short:   "Restart a particular instance",
	Long: `This command lets you restart a particular instance.
To use this command you have to provide the instance ID`,
	Run: func(cmd *cobra.Command, args []string) {
		status := "reboot"
		user.HandleUpdateInstance(&instanceId, &status)
	},
}

func init() {
	RestartCmd.Flags().StringVarP(&instanceId, "instance", "i", "", "The instance id")

	err := RestartCmd.MarkFlagRequired("instance")
	if nil != err {
		fmt.Println(err)
	}
}
