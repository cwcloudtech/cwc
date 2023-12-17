package user

import (
	"cwc/client"
	"cwc/utils"
	"fmt"
)

func HandleListProviders() {
	providers, err := client.GetProviders()
	utils.ExitIfError(err)

	for _, available_provider := range providers.Providers {
		fmt.Printf("%v\n", available_provider.Name)
	}
}

func HandlerGetDefaultProvider() {
	provider := client.GetDefaultProvider()
	fmt.Printf("Default provider = %v\n", provider)
}

func HandlerSetDefaultProvider(value string) {
	providers, err := client.GetProviders()
	utils.ExitIfError(err)

	available_providers := []string{}
	for _, available_provider := range providers.Providers {
		available_providers = append(
			available_providers,
			available_provider.Name,
		)
	}

	utils.ExitIfNeeded("Invalid provider value", !utils.StringInSlice(value, available_providers))

	client.SetDefaultProvider(value)
	fmt.Printf("Default provider = %v\n", value)
}
