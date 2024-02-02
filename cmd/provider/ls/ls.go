package ls

import (
	"cwc/client"
	"cwc/handlers/user"
	"cwc/utils"

	"github.com/spf13/cobra"
)

var (
	pretty bool
)

// lsCmd represents the ls command
var LsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List available providers",
	Long: `This command lets you list the available providers in the cloud
This command takes no arguments`,
	Run: func(cmd *cobra.Command, args []string) {
		providers, err := client.GetProviders()
		utils.ExitIfError(err)
		user.HandleListProviders(providers, &pretty)
	},
}

func init() {
	LsCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "Pretty print the output (optional)")
}
