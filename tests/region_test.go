package test

import (
    "cwc/client"
    "cwc/handlers/user"
    "testing"
)

func TestListRegions (t *testing.T) {
    mockRegions := &client.ProviderRegions{
        Regions: []client.ProviderRegion{
            {
                Name: "region1",
            },
            {
                Name: "region2",
            },
        },
    }
    pretty := true
    user.HandleListRegions(mockRegions, &pretty)
}
