package ls

import (
	"cwc/handlers/user"

	"github.com/spf13/cobra"
)

var (
	pretty bool = false
)

// lsCmd represents the ls command
var LsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List available Dns Zones",
	Long: `This command lets you list the available Dns Zones in the cloud
This command takes no arguments`,
	Run: func(cmd *cobra.Command, args []string) {
		user.HandleListDnsZones(&pretty)
	},
}

func init() {
	LsCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "Pretty print the output (optional)")
}
