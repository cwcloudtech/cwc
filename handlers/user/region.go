package user

import (
	"cwc/client"
	"cwc/config"
	"cwc/utils"
	"fmt"
)

func HandleListRegions(pretty *bool) {
	provider_regions, err := client.GetProviderRegions()
	utils.ExitIfError(err)

	var region_names []string
	for _, available_region := range provider_regions.Regions {
		region_names = append(region_names, available_region.Name)
	}

	if config.IsPrettyFormatExpected(pretty) {
		utils.PrintPrettyArray("Available regions", region_names)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(provider_regions.Regions)
	} else {
		utils.PrintArray(region_names)
	}
}

func HandlerGetDefaultRegion() {
	region := config.GetDefaultRegion()
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

	config.SetDefaultRegion(value)
	fmt.Printf("Default region = %v\n", value)
}
