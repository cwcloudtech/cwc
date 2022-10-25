package user

import (
	"cwc/client"
	"cwc/utils"
	"fmt"
	"os"
)

func HandleGetEnvironments() {

	c, err := client.NewClient()

	environments, err := c.GetAllEnvironments()

	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	if client.GetDefaultFormat() == "json" {
		utils.PrintJson(environments)
	} else {
		utils.PrintMultiRow(client.Environment{}, *environments)
	}

	return

}

func HandleGetEnvironment(id *string) {
	c, err := client.NewClient()

	environment, err := c.GetEnvironment(*id)
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	if client.GetDefaultFormat() == "json" {
		utils.PrintJson(environment)
	} else {
		utils.PrintRow(*environment)
	}

	return

}
