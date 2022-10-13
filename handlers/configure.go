package handlers

import (
	"cwc/client"
	"cwc/utils"
	"flag"
	"fmt"
	"os"
)

func HandleConfigure(configureCmd *flag.FlagSet, region *bool, endpoint *bool, provider *bool) {
	if len(os.Args) <= 2 {
		configureCmd.PrintDefaults()
		os.Exit(1)
	}
	configureCmd.Parse(os.Args[2:])
	if !*region && !*endpoint && !*provider {
		if os.Args[2] == "help" {
			configureCmd.PrintDefaults()

		} else {
			fmt.Println("cwc: flag is missing")
		}
		os.Exit(1)
	}

	if len(os.Args) < 4 {
		fmt.Println("cwc: invalid arguments")
		fmt.Println("cwc configure <command> <subcommand> <value>")
		fmt.Println("<command>: -region")
		fmt.Println("<subcommand>: get / set")
		fmt.Println("<command>: -endpoint")
		fmt.Println("<subcommand>: get / set")
		os.Exit(1)
	}
	subSubCommmand := os.Args[3]
	switch subSubCommmand {
	case "set":

		if len(os.Args) <= 4 {
			fmt.Println("cwc: missing mandatory parameter, please check cwc --help.")
			os.Exit(1)
		}
		value := os.Args[4]
		if value == "" {
			fmt.Println("cwc: value is missing")
			configureCmd.PrintDefaults()
			os.Exit(1)
		}
		if *region {
			provider_regions, err := client.GetProviderRegions()
			if err != nil {
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
		if *endpoint {
			client.SetDefaultEndpoint(value)
			fmt.Printf("Default endpoint = %v\n", value)
		}
		if *provider {
			providers, err := client.GetProviders()
			if err != nil {
				fmt.Printf("failed: %s\n", err)
				os.Exit(1)

			}
			available_providers := []string{}
			for _, available_provider := range providers.Providers {
				available_providers = append(
					available_providers,
					available_provider.Name,
				)

			}
			if !utils.StringInSlice(value, available_providers) {
				fmt.Println("cwc: invalid provider value")
				os.Exit(1)

			}
			client.SetDefaultProvider(value)
			fmt.Printf("Default provider = %v\n", value)
		}

	case "get":
		if *region {
			region := client.GetDefaultRegion()
			fmt.Printf("Default region = %v\n", region)

		}
		if *endpoint {
			endpoint := client.GetDefaultEndpoint()
			fmt.Printf("Default endpoint = %v\n", endpoint)
		}
		if *provider {
			provider := client.GetDefaultProvider()
			fmt.Printf("Default provider = %v\n", provider)
		}
	default:
		fmt.Printf("cwc: option not found")
	}

}
