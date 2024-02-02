package test

import (
	"cwc/client"
	"cwc/handlers/user"
	"testing"
)

func TestListProviders(t *testing.T) {
	mockProviders := &client.AvailableProviders{
		Providers: []client.Provider{
			{
				Name: "provider1",
			},
			{
				Name: "provider2",
			},
		},
	}
	pretty := true
	user.HandleListProviders(mockProviders, &pretty)
}
