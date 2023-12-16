package user

import (
	"cwc/client"
	"cwc/utils"
	"fmt"
	"os"
)

func HandleDeleteRegistry(id *string) {
	c, err := client.NewClient()
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

func HandleUpdateRegistry(id *string) {
	c, err := client.NewClient()
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	err = c.UpdateRegistry(*id)
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Registry %v successfully updated\n", *id)
}

func HandleGetRegistries() {
	c, err := client.NewClient()
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	registries, err := c.GetAllRegistries()
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	if client.GetDefaultFormat() == "json" {
		utils.PrintJson(registries)
	} else {
		utils.PrintMultiRow(client.Registry{}, *registries)
	}
}

func HandleGetRegistry(id *string) {
	c, err := client.NewClient()
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	registry, err := c.GetRegistry(*id)
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	if client.GetDefaultFormat() == "json" {
		utils.PrintJson(registry)
	} else {
		utils.PrintRow(*registry)
	}
}
