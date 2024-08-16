package delete

import (
	"cwc/handlers/admin"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	dnsRecordId string
	dnsRecordName string
	dnsZone string
)

var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a particular DNS record",
	Long: `This command lets you delete a particular DNS record.
To use this command you have to provide the DNS record ID that you want to delete`,
	Run: func(cmd *cobra.Command, args []string) {
		admin.HandleDeleteDnsRecord(&dnsRecordId, &dnsRecordName, &dnsZone)
	},
}

func init() {
	DeleteCmd.Flags().StringVarP(&dnsRecordId, "record", "r", "", "The DNS record id")
	DeleteCmd.Flags().StringVarP(&dnsRecordName, "name", "n", "", "The DNS record name")
	DeleteCmd.Flags().StringVarP(&dnsZone, "zone", "z", "", "The DNS zone")

	err := DeleteCmd.MarkFlagRequired("record")
	if nil != err {
		fmt.Println(err)
	}

	err = DeleteCmd.MarkFlagRequired("name")
	if nil != err {
		fmt.Println(err)
	}

	err = DeleteCmd.MarkFlagRequired("zone")
	if nil != err {
		fmt.Println(err)
	}
}