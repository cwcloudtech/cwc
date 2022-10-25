package admin

import (
	"cwc/admin"
	"cwc/utils"
	"fmt"
	"os"
)

func HandleAddEnvironment(name *string, path *string, roles *[]string, main_role *string, is_private *bool, description *string, subdomains *[]string) {
	client, err := admin.NewClient()
	created_env, err := client.AdminAddEnvironment(*name, *path, *roles, *main_role, *is_private, *description, *subdomains)
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	if admin.GetDefaultFormat() == "json" {
		utils.PrintJson(created_env)
	} else {
		utils.PrintRow(*created_env)
	}

}

func HandleDeleteEnvironment(id *string) {

	client, err := admin.NewClient()
	err = client.AdminDeleteEnvironment(*id)
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Environment %v successfully deleted\n", *id)
}

func HandleGetEnvironments() {

	client, err := admin.NewClient()
	environments, err := client.GetAllEnvironments()

	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	if admin.GetDefaultFormat() == "json" {
		utils.PrintJson(environments)
	} else {
		utils.PrintMultiRow(admin.Environment{}, *environments)
	}

	return

}

func HandleGetEnvironment(id *string) {
	client, err := admin.NewClient()

	environment, err := client.GetEnvironment(*id)
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	if admin.GetDefaultFormat() == "json" {
		utils.PrintJson(environment)
	} else {
		utils.PrintRow(*environment)
	}

	return

}
