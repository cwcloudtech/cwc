package admin

import (
	"cwc/admin"
	"cwc/config"
	"cwc/utils"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

func HandleAddDnsRecord(recordName *string, dnsZone *string, dnsType *string, ttl *int, data *string) {
	c, err := admin.NewClient()
	utils.ExitIfError(err)

	created_record, err := c.AdminAddDnsRecord(*recordName, *dnsZone, *dnsType, *ttl, *data)
	utils.ExitIfError(err)

	if config.GetDefaultFormat() == "json" {
		utils.PrintJson(created_record)
	} else {
		utils.PrintRow(*created_record)
	}
}

func HandleGetDnsRecords(dnsRecords *[]admin.DnsRecord, pretty *bool) {
	if config.IsPrettyFormatExpected(pretty) {
		displayDnsRecordsAsTable(*dnsRecords)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(dnsRecords)
	}
}

func HandleDeleteDnsRecord(recordId *string, recordName *string, dnsZone *string) {
	c, err := admin.NewClient()
	utils.ExitIfError(err)

	err = c.DeleteDnsRecord(*recordId, *recordName, *dnsZone)
	utils.ExitIfError(err)

	fmt.Printf("DNS record with id %v successfully deleted\n", *recordId)
}

func displayDnsRecordsAsTable(dnsRecords []admin.DnsRecord) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Zone", "Record", "TTL", "Type", "Data"})

	if len(dnsRecords) == 0 {
		fmt.Println("No DNS records found")
	} else {
		for _, record := range dnsRecords {
			table.Append([]string{
				record.Id,
				record.Zone,
				record.Record,
				fmt.Sprintf("%d", record.Ttl), 
				record.Type, 
				record.Data,
			})
		}
		table.Render()
	}
}
