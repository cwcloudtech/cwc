package delete

import (
	"cwc/handlers/user"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	monitorId string
)

var DeleteCmd = &cobra.Command{
	Use:  "delete",
	Short: "Delete a particular monitor",
	Long: `This command lets you delete a particular monitor.
To use this command you have to provide the monitor ID that you want to delete.`,
	Run: func(cmd *cobra.Command, args []string) {
		user.HandleDeleteMonitor(&monitorId)
	},
}

func init() {
	DeleteCmd.Flags().StringVarP(&monitorId, "monitor", "m", "", "The monitor id")

	err := DeleteCmd.MarkFlagRequired("monitor")
	if nil != err {
		fmt.Println(err)
	}
}
