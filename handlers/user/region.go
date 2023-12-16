package user

import (
	"cwc/client"
	"cwc/utils"
	"fmt"
	"os"
)

func HandleListRegions() {
	provider_regions, err := client.GetProviderRegions()
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	for _, available_region := range provider_regions.Regions {
		fmt.Printf("%v\n", available_region.Name)
	}

	return
}

func HandlerGetDefaultRegion() {
	region := client.GetDefaultRegion()
	fmt.Printf("Default region = %v\n", region)
}

func HandlerSetDefaultRegion(value string) {
	provider_regions, err := client.GetProviderRegions()
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)

	}

	available_regions := []string{}
	for _, available_region := range provider_regions.Regions {
		available_regions = append(
			available_regions,
			available_region.Name,
		)

	}

	if !utils.StringInSlice(value, available_regions) {
		fmt.Println("cwc: invalid region value")
		os.Exit(1)

	}

	client.SetDefaultRegion(value)
	fmt.Printf("Default region = %v\n", value)
}
