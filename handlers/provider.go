package handlers

import (
	"cwc/client"
	"cwc/utils"
	"fmt"
	"os"
)

func HandleListProviders() {

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

func HandlerGetDefaultProvider() {

	provider := client.GetDefaultProvider()
	fmt.Printf("Default provider = %v\n", provider)

}

func HandlerSetDefaultProvider(value string) {
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

}
