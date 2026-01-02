package ls

import (
	adminClient "cwc/admin"
	"cwc/handlers/admin"
	"cwc/utils"

	"github.com/spf13/cobra"
)

var (
	pretty bool
)

var LCmd = &cobra.Command{
	Use:   "ls",
	Short: "List available DNS records",
	Long: `This command lets you list your available DNS records with cwcloud
This command takes no arguments`,
	Run: func(cmd *cobra.Command, args []string) {
		c, err := adminClient.NewClient()
		utils.ExitIfError(err)
		dnsRecords, err := c.GetAllDnsRecords()
		utils.ExitIfError(err)
		admin.HandleGetDnsRecords(dnsRecords, &pretty)
	},
}

func init() {
	LCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "Pretty print the output (optional)")
}
