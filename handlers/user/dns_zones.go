package user

import (
	"cwc/client"
	"cwc/utils"
	"fmt"
)

func HandleListDnsZones() {
	dns_zones, err := client.GetDnsZones()
	utils.ExitIfError(err)

	for _, dns_zone := range dns_zones.Zones {
		fmt.Printf("%v\n", dns_zone)
	}
}
