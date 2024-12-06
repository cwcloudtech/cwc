package dns_records

import (
	"cwc/cmd/admin/dns_records/create"
	"cwc/cmd/admin/dns_records/delete"
	"cwc/cmd/admin/dns_records/ls"

	"github.com/spf13/cobra"
)

// dnsRecordsCmd represents the dnsRecords command

var DnsRecordsCmd = &cobra.Command{
	Use:   "dnsRecord",
	Short: "Manage your DNS records in the cloud",
	Long: `This command lets you manage your DNS records in the cloud.
Several actions are associated with this command such as update a DNS record, deleting a DNS record
and listing your available DNS records`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	DnsRecordsCmd.DisableFlagsInUseLine = true
	DnsRecordsCmd.AddCommand(ls.LCmd)
	DnsRecordsCmd.AddCommand(create.CreateCmd)
	DnsRecordsCmd.AddCommand(delete.DeleteCmd)
}
