package admin

import (
	"cwc/admin"
	"cwc/utils"
	"fmt"
	"os"
)

func HandleDeleteInstance(id *string) {
	c, err := admin.NewClient()
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	err = c.AdminDeleteInstance(*id)
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Instance %v successfully deleted\n", *id)
}

func HandleRefreshInstance(id *string) {
	c, err := admin.NewClient()
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	err = c.AdminRefreshInstance(*id)
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Instance %v state successfully refreshed\n", *id)
}

func HandleAddInstance(user_email *string, name *string, project_id *int, project_name *string, project_url *string, env *string, instance_type *string, zone *string, dns_zone *string) {
	c, err := admin.NewClient()
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	created_instance, err := c.AdminAddInstance(*user_email, *name, *project_id, *project_name, *project_url, *instance_type, *env, *zone, *dns_zone)
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
	c, err := admin.NewClient()
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	err = c.AdminUpdateInstance(*id, *status)
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Instance %v successfully updated\n", *id)
}

func HandleGetInstances() {
	c, err := admin.NewClient()
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	instances, err := c.AdminGetAllInstances()
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
	c, err := admin.NewClient()
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	instance, err := c.GetInstance(*id)
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
