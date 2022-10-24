package admin

import (
	"cwc/admin"
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
	fmt.Printf("ID\tname\tpath\tmain_role\troles\tdescription\n")
	fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\n", created_env.Id, created_env.Name, created_env.Path, created_env.MainRole, created_env.Roles, created_env.Description)

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


func HandleGetEnvironment() {

	client, err := admin.NewClient()
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	environments, err := client.GetAllEnvironments()

	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("ID\tname\tpath\tdescription\n")
	for _, environment := range *environments {
		fmt.Printf("%v\t%v\t%v\t%v\n", environment.Id, environment.Name, environment.Path, environment.Description)
	}

	return

}