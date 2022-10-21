package admin

import (
	"cwc/admin"
	"fmt"
	"os"
)

func HandleAddRegistry(user_email *string, name *string, reg_type *string) {
	client, err := admin.NewClient()
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	created_registry, err := client.AdminAddRegistry(*user_email, *name, *reg_type)
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("ID\tcreated_at\tname\tstatus\taccess_key\tsecret_key\tendpoint\n")
	fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\t%v\n", created_registry.Id, created_registry.CreatedAt, created_registry.Name, created_registry.Status, created_registry.AccessKey, created_registry.SecretKey, created_registry.Endpoint)
}

func HandleDeleteRegistry(id *string) {
	client, err := admin.NewClient()
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

func HandleUpdateRegistry(id *string) {
	client, err := admin.NewClient()
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

func HandleGetRegistry() {

	client, err := admin.NewClient()
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

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
