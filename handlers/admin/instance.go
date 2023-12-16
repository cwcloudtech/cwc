package admin

import (
	"cwc/admin"
	"cwc/utils"
	"fmt"
	"os"
)

func HandleDeleteInstance(id *string) {
	client, err := admin.NewClient()
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	err = client.AdminDeleteInstance(*id)
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Instance %v successfully deleted\n", *id)
}

func HandleRefreshInstance(id *string) {
	client, err := admin.NewClient()
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	err = client.AdminRefreshInstance(*id)
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Instance %v state successfully refreshed\n", *id)
}

func HandleAddInstance(user_email *string, name *string, project_id *int, project_name *string, project_url *string, env *string, instance_type *string, zone *string, dns_zone *string) {
	client, err := admin.NewClient()
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	created_instance, err := client.AdminAddInstance(*user_email, *name, *project_id, *project_name, *project_url, *instance_type, *env, *zone, *dns_zone)
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	if admin.GetDefaultFormat() == "json" {
		utils.PrintJson(created_instance)
	} else {
		utils.PrintRow(*created_instance)
	}
}

func HandleUpdateInstance(id *string, status *string) {
	client, err := admin.NewClient()
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	err = client.AdminUpdateInstance(*id, *status)
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Instance %v successfully updated\n", *id)
}

func HandleGetInstances() {
	client, err := admin.NewClient()
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	instances, err := client.AdminGetAllInstances()
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	if admin.GetDefaultFormat() == "json" {
		utils.PrintJson(instances)
	} else {
		utils.PrintMultiRow(admin.Instance{}, *instances)
	}
}

func HandleGetInstance(id *string) {
	client, err := admin.NewClient()
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	instance, err := client.GetInstance(*id)
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	if admin.GetDefaultFormat() == "json" {
		utils.PrintJson(instance)
	} else {
		utils.PrintRow(*instance)
	}
}
