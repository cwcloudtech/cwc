package handlers

import (
	"cwc/client"
	"flag"
	"fmt"
	"os"
)

func HandleListDnsZones(dnsZonesCmd *flag.FlagSet) {

	dnsZonesCmd.Parse(os.Args[2:])
	dns_zones, err := client.GetDnsZones()
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)

	}
	fmt.Printf("Dns Zones \n\n")
	for _, dns_zone := range dns_zones.Zones {
		fmt.Printf("%v\n", dns_zone)

	}
	return
}
