package test

import (
	"cwc/client"
	"cwc/handlers/user"
	"testing"
)

func TestHandleListDnsZonesWithPretty(t *testing.T) {
	mockDnsZones := &client.Dns_zones{
		Zones: []string{"example.com", "example.org"},
	}
	mockPretty := true
	user.HandleListDnsZones(mockDnsZones, &mockPretty)
}
