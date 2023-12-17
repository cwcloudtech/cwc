package admin

import (
	"cwc/admin"
	"cwc/utils"
	"fmt"
)

func HandleAddRegistry(user_email *string, name *string, reg_type *string) {
	c, err := admin.NewClient()
	utils.ExitIfError(err)

	created_registry, err := c.AdminAddRegistry(*user_email, *name, *reg_type)
	utils.ExitIfError(err)

	if admin.GetDefaultFormat() == "json" {
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

func HandleUpdateRegistry(id *string, email *string) {
	c, err := admin.NewClient()
	utils.ExitIfError(err)

	err = c.UpdateRegistry(*id, *email)
	utils.ExitIfError(err)

	fmt.Printf("Registry %v successfully updated\n", *id)
}

func HandleGetRegistries() {
	c, err := admin.NewClient()
	utils.ExitIfError(err)

	registries, err := c.GetAllRegistries()
	utils.ExitIfError(err)

	if admin.GetDefaultFormat() == "json" {
		utils.PrintJson(registries)
	} else {
		utils.PrintMultiRow(admin.Project{}, *registries)
	}
}

func HandleGetRegistry(id *string) {
	c, err := admin.NewClient()
	utils.ExitIfError(err)

	registry, err := c.GetRegistry(*id)
	utils.ExitIfError(err)

	if admin.GetDefaultFormat() == "json" {
		utils.PrintJson(registry)
	} else {
		utils.PrintRow(*registry)
	}
}
