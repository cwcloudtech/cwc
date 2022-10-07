package handlers

import (
	"cwc/client"
	"flag"
	"fmt"
	"os"
)

func HandleGetEnvironment(getCmd *flag.FlagSet, all *bool, id *string) {

	getCmd.Parse(os.Args[3:])
	if !*all && *id == "" {
		fmt.Println("id is required or specify --all to get all projects.")
		getCmd.PrintDefaults()
		os.Exit(1)
	}
	client := client.NewClient()
	if *all {

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

	if *id != "" {
		environment, err := client.GetEnvironment(*id)
		if err != nil {
			fmt.Printf("failed: %s\n", err)
			os.Exit(1)
		}
		fmt.Printf("ID\tname\tpath\tdescription\n")
		fmt.Printf("%v\t%v\t%v\t%v\n", environment.Id, environment.Name, environment.Path, environment.Description)

		return
	}
}
