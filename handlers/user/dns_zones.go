package user

import (
	"cwc/client"
	"cwc/config"
	"cwc/utils"
)

func HandleListDnsZones(dns_zones *client.Dns_zones, pretty *bool) {
	if config.IsPrettyFormatExpected(pretty) {
		utils.PrintPrettyArray("DNS zones available", dns_zones.Zones)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(dns_zones)
	} else {
		utils.PrintArray(dns_zones.Zones)
	}
}
