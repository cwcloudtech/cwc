package ls

import (
	adminClient "cwc/admin"
	"cwc/handlers/admin"
	"cwc/utils"
	"github.com/spf13/cobra"
)

var (
	pretty bool = false
)

var LsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List available environments",
	Long: `This command lets you list your available environments in the cloud
This command takes no arguments`,
	Run: func(cmd *cobra.Command, args []string) {
		c, err := adminClient.NewClient()
		utils.ExitIfError(err)
		environments, err := c.GetAllEnvironments(true)
		utils.ExitIfError(err)
		admin.HandleGetEnvironments(environments, &pretty)
	},
}

func init() {
	LsCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "Pretty print the output (optional)")
}
