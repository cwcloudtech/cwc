package user

import (
	"cwc/client"
	"cwc/utils"
	"fmt"
)

func HandleListRegions() {
	provider_regions, err := client.GetProviderRegions()
	utils.ExitIfError(err)

	for _, available_region := range provider_regions.Regions {
		fmt.Printf("%v\n", available_region.Name)
	}
}

func HandlerGetDefaultRegion() {
	region := client.GetDefaultRegion()
	fmt.Printf("Default region = %v\n", region)
}

func HandlerSetDefaultRegion(value string) {
	provider_regions, err := client.GetProviderRegions()
	utils.ExitIfError(err)

	available_regions := []string{}
	for _, available_region := range provider_regions.Regions {
		available_regions = append(
			available_regions,
			available_region.Name,
		)
	}

	utils.ExitIfNeeded("Invalid region", !utils.StringInSlice(value, available_regions))

	client.SetDefaultRegion(value)
	fmt.Printf("Default region = %v\n", value)
}
