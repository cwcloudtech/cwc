package user

import (
	"cwc/client"
	"fmt"
	"os"
)

func HandleListDnsZones() {
	dns_zones, err := client.GetDnsZones()
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)

	}
	for _, dns_zone := range dns_zones.Zones {
		fmt.Printf("%v\n", dns_zone)

	}
	return
}
