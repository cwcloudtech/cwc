package handlers

import (
	"cwc/client"
	"fmt"
	"os"
)

func HandleGetEnvironment() {

	client, err := client.NewClient()
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
