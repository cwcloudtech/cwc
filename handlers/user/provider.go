package user

import (
	"cwc/client"
	"cwc/config"
	"cwc/utils"
	"fmt"
)

func HandleListProviders(providers *client.AvailableProviders, pretty *bool) {
	var names []string
	for _, available_provider := range providers.Providers {
		names = append(names, available_provider.Name)
	}

	if config.IsPrettyFormatExpected(pretty) {
		utils.PrintPrettyArray("Available providers", names)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(providers.Providers)
	} else {
		utils.PrintArray(names)
	}
}

func HandlerGetDefaultProvider() {
	provider := config.GetDefaultProvider()
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

	config.SetDefaultProvider(value)
	fmt.Printf("Default provider = %v\n", value)
}
