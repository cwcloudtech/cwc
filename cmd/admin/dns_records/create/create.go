package create 

import (
	"cwc/handlers/admin"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	recordName string
	dnsZone   string
	dnsType 	string
	ttl 		int
	data 		string
)

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a DNS record in the cloud",
	Long:  `This command lets you create a DNS record in the cloud`,
	Run: func(cmd *cobra.Command, args []string) {
		admin.HandleAddDnsRecord(&recordName, &dnsZone, &dnsType, &ttl, &data)
	},
}

func init() {
	CreateCmd.Flags().StringVarP(&recordName, "name", "n", "", "The record name")
	CreateCmd.Flags().StringVarP(&dnsZone, "zone", "z", "", "The DNS zone")
	CreateCmd.Flags().StringVarP(&dnsType, "type", "t", "", "The record type (A, CNAME, TXT, MX, NS, SOA, SRV, PTR, AAAA)")
	CreateCmd.Flags().IntVarP(&ttl, "ttl", "l", 0, "The record TTL")
	CreateCmd.Flags().StringVarP(&data, "data", "d", "", "The record data")

	err := CreateCmd.MarkFlagRequired("name")
	if nil != err {
		fmt.Println(err)
	}

	err = CreateCmd.MarkFlagRequired("zone")
	if nil != err {
		fmt.Println(err)
	}

	err = CreateCmd.MarkFlagRequired("type")
	if nil != err {
		fmt.Println(err)
	}

	err = CreateCmd.MarkFlagRequired("ttl")
	if nil != err {
		fmt.Println(err)
	}

	err = CreateCmd.MarkFlagRequired("data")
	if nil != err {
		fmt.Println(err)
	}
}