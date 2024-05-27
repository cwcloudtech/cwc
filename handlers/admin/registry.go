package admin

import (
	"cwc/admin"
	"cwc/config"
	"cwc/utils"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

func HandleAddRegistry(user_email *string, name *string, reg_type *string) {
	c, err := admin.NewClient()
	utils.ExitIfError(err)

	created_registry, err := c.AdminAddRegistry(*user_email, *name, *reg_type)
	utils.ExitIfError(err)

	if config.GetDefaultFormat() == "json" {
		utils.PrintJson(created_registry)
	} else {
		utils.PrintRow(*created_registry)
	}
}

func HandleDeleteRegistry(id *string) {
	c, err := admin.NewClient()
	utils.ExitIfError(err)

	err = c.DeleteRegistry(*id)
	utils.ExitIfError(err)

	fmt.Printf("Registry %v successfully deleted\n", *id)
}

func HandleTransferRegistryOwnership(id *string, email *string) {
	c, err := admin.NewClient()
	utils.ExitIfError(err)

	err = c.UpdateRegistry(*id, *email)
	utils.ExitIfError(err)

	fmt.Printf("Registry with id %v successfully transferred to this email owner: %v\n", *id, *email)
}

func HandleRenewRegistryCredentials(id *string) {
	c, err := admin.NewClient()
	utils.ExitIfError(err)

	err = c.UpdateRegistry(*id)
	utils.ExitIfError(err)

	fmt.Printf("Registry with id %v successfully renewed credentials\n", *id)
}

func HandleGetRegistries(registries *[]admin.Registry, pretty *bool) {
	if config.IsPrettyFormatExpected(pretty) {
		displayRegistriesAsTable(*registries)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(registries)
	} else {
		utils.PrintMultiRow(admin.Project{}, *registries)
	}
}

func HandleGetRegistry(registry *admin.Registry, pretty *bool) {
	if config.IsPrettyFormatExpected(pretty) {
		utils.PrintPretty("Registry's informations", *registry)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(registry)
	} else {
		utils.PrintRow(*registry)
	}
}

func displayRegistriesAsTable(registries []admin.Registry) {
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
		table.Render()
	}
}
