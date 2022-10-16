package handlers

import (
	"cwc/client"
	"flag"
	"fmt"
	"os"
)

func HandleListProvider(providerCmd *flag.FlagSet) {

	providerCmd.Parse(os.Args[2:])
	providers, err := client.GetProviders()
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)

	}
	for _, available_provider := range providers.Providers {
		fmt.Printf("%v\n", available_provider.Name)

	}
	return
}
