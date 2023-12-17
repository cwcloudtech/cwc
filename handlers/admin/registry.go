package admin

import (
	"cwc/admin"
	"cwc/utils"
	"fmt"
	"os"
)

func HandleAddRegistry(user_email *string, name *string, reg_type *string) {
	c, err := admin.NewClient()
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	created_registry, err := c.AdminAddRegistry(*user_email, *name, *reg_type)
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	if admin.GetDefaultFormat() == "json" {
		utils.PrintJson(created_registry)
	} else {
		utils.PrintRow(*created_registry)
	}
}

func HandleDeleteRegistry(id *string) {
	c, err := admin.NewClient()
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	err = c.DeleteRegistry(*id)
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Registry %v successfully deleted\n", *id)
}

func HandleUpdateRegistry(id *string, email *string) {
	c, err := admin.NewClient()
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	err = c.UpdateRegistry(*id, *email)
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Registry %v successfully updated\n", *id)
}

func HandleGetRegistries() {
	c, err := admin.NewClient()
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	registries, err := c.GetAllRegistries()
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	if admin.GetDefaultFormat() == "json" {
		utils.PrintJson(registries)
	} else {
		utils.PrintMultiRow(admin.Project{}, *registries)
	}
}

func HandleGetRegistry(id *string) {
	c, err := admin.NewClient()
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	registry, err := c.GetRegistry(*id)
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	if admin.GetDefaultFormat() == "json" {
		utils.PrintJson(registry)
	} else {
		utils.PrintRow(*registry)
	}
}
