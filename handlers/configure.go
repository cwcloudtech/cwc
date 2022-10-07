package handlers

import (
	"cwc/client"
	"cwc/utils"
	"flag"
	"fmt"
	"os"
)

func HandleConfigure(configureCmd *flag.FlagSet, region *bool, endpoint *bool, provider *bool) {

	configureCmd.Parse(os.Args[2:])
	if !*region && !*endpoint && !*provider {
		fmt.Println("cwc: flag is missing")
		configureCmd.PrintDefaults()
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
		value := os.Args[4]
		if value == "" {
			fmt.Println("cwc: value is missing")
			configureCmd.PrintDefaults()
			os.Exit(1)
		}
		if *region {

			possible_regions := make([]string, 9)
			possible_regions[0] = "fr-par"
			possible_regions[1] = "nl-ams"
			possible_regions[2] = "pl-waw"
			possible_regions[3] = "SBG5"
			possible_regions[4] = "GRA11"
			possible_regions[5] = "UK1"
			possible_regions[6] = "DE1"
			possible_regions[7] = "BHS5"
			possible_regions[8] = "WAW1"

			if !utils.StringInSlice(value, possible_regions) {
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
