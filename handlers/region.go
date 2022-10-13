package handlers

import (
	"cwc/client"
	"flag"
	"fmt"
	"os"
)

func HandleListRegions(regionCmd *flag.FlagSet) {

	regionCmd.Parse(os.Args[2:])
	provider_regions, err := client.GetProviderRegions()
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)

	}
	fmt.Printf("regions \n\n")
	for _, available_region := range provider_regions.Regions {
		fmt.Printf("%v\n", available_region.Name)

	}
	return
}
