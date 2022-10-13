package handlers

import (
	"cwc/client"
	"flag"
	"fmt"
	"os"
)

func HandleDeleteRegistry(deleteCmd *flag.FlagSet, id *string) {

	deleteCmd.Parse(os.Args[3:])
	if *id == "" {
		fmt.Println("id is required to delete your registry")
		deleteCmd.PrintDefaults()
		os.Exit(1)
	}
	client, err := client.NewClient()
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	err = client.DeleteRegistry(*id)
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Registry %v successfully deleted\n", *id)
}

func HandleUpdateRegistry(updateCmd *flag.FlagSet, id *string) {
	updateCmd.Parse(os.Args[3:])
	if *id == "" {
		fmt.Println("id are required")
		updateCmd.PrintDefaults()
		os.Exit(1)
	}
	client, err := client.NewClient()
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	err = client.UpdateRegistry(*id)
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Registry %v successfully updated\n", *id)

}

func HandleGetRegistry(getCmd *flag.FlagSet, all *bool, id *string) {

	getCmd.Parse(os.Args[3:])
	if !*all && *id == "" {
		fmt.Println("id is required or specify --all to get all registries.")
		getCmd.PrintDefaults()
		os.Exit(1)
	}
	client, err := client.NewClient()
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	if *all {

		registries, err := client.GetAllRegistries()

		if err != nil {
			fmt.Printf("failed: %s\n", err)
			os.Exit(1)
		}

		fmt.Printf("ID\tcreated_at\tname\tstatus\taccess_key\tsecret_key\tendpoint\n")
		for _, registry := range *registries {
			fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\t%v\n", registry.Id, registry.CreatedAt, registry.Name, registry.Status, registry.AccessKey, registry.SecretKey, registry.Endpoint)

		}

		return
	}

	if *id != "" {
		registry, err := client.GetRegistry(*id)
		if err != nil {
			fmt.Printf("failed: %s\n", err)
			os.Exit(1)
		}
		fmt.Printf("ID\tcreated_at\tname\tstatus\taccess_key\tsecret_key\tendpoint\n")
		fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\t%v\n", registry.Id, registry.CreatedAt, registry.Name, registry.Status, registry.AccessKey, registry.SecretKey, registry.Endpoint)

		return
	}
}
