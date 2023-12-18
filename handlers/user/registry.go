package user

import (
	"cwc/client"
	"cwc/config"
	"cwc/utils"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

func HandleDeleteRegistry(id *string) {
	c, err := client.NewClient()
	utils.ExitIfError(err)

	err = c.DeleteRegistry(*id)
	utils.ExitIfError(err)

	fmt.Printf("Registry %v successfully deleted\n", *id)
}

func HandleUpdateRegistry(id *string) {
	c, err := client.NewClient()
	utils.ExitIfError(err)

	err = c.UpdateRegistry(*id)
	utils.ExitIfError(err)

	fmt.Printf("Registry %v successfully updated\n", *id)
}

func HandleGetRegistries(pretty *bool) {
	c, err := client.NewClient()
	utils.ExitIfError(err)

	registries, err := c.GetAllRegistries()
	utils.ExitIfError(err)

	if config.IsPrettyFormatExpected(pretty) {
		displayRegistriesAsTable(*registries)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(registries)
	} else {
		utils.PrintMultiRow(client.Registry{}, *registries)
	}
}

func HandleGetRegistry(id *string, pretty *bool) {
	c, err := client.NewClient()
	utils.ExitIfError(err)

	registry, err := c.GetRegistry(*id)
	utils.ExitIfError(err)

	if config.IsPrettyFormatExpected(pretty) {
		utils.PrintPretty("Registry's informations", *registry)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(registry)
	} else {
		utils.PrintRow(*registry)
	}
}

func displayRegistriesAsTable(registries []client.Registry) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "Type", "Endpoint", "Region", "Created at"})

	if len(registries) == 0 {
		fmt.Println("No registry found")
	} else {
		for _, registry := range registries {
			table.Append([]string{
				fmt.Sprintf("%d", registry.Id),
				registry.Name,
				registry.Type,
				registry.Endpoint,
				registry.Region,
				registry.CreatedAt,
			})
		}
	}

	table.Render()
}
