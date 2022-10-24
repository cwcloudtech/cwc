package user

import (
	"cwc/client"
	"fmt"
	"os"
)

func HandleGetEnvironments() {

	client, err := client.NewClient()

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

func HandleGetEnvironment(id *string) {
	client, err := client.NewClient()

	environment, err := client.GetEnvironment(*id)
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("ID\tname\tpath\tdescription\n")
	fmt.Printf("%v\t%v\t%v\t%v\n", environment.Id, environment.Name, environment.Path, environment.Description)

	return
	
}

