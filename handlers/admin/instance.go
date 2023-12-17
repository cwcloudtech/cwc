package admin

import (
	"cwc/admin"
	"cwc/utils"
	"fmt"
)

func HandleDeleteInstance(id *string) {
	c, err := admin.NewClient()
	utils.ExitIfError(err)

	err = c.AdminDeleteInstance(*id)
	utils.ExitIfError(err)

	fmt.Printf("Instance %v successfully deleted\n", *id)
}

func HandleRefreshInstance(id *string) {
	c, err := admin.NewClient()
	utils.ExitIfError(err)

	err = c.AdminRefreshInstance(*id)
	utils.ExitIfError(err)

	fmt.Printf("Instance %v state successfully refreshed\n", *id)
}

func HandleAddInstance(user_email *string, name *string, project_id *int, project_name *string, project_url *string, env *string, instance_type *string, zone *string, dns_zone *string) {
	c, err := admin.NewClient()
	utils.ExitIfError(err)

	created_instance, err := c.AdminAddInstance(*user_email, *name, *project_id, *project_name, *project_url, *instance_type, *env, *zone, *dns_zone)
	utils.ExitIfError(err)

	if admin.GetDefaultFormat() == "json" {
		utils.PrintJson(created_instance)
	} else {
		utils.PrintRow(*created_instance)
	}
}

func HandleUpdateInstance(id *string, status *string) {
	c, err := admin.NewClient()
	utils.ExitIfError(err)

	err = c.AdminUpdateInstance(*id, *status)
	utils.ExitIfError(err)

	fmt.Printf("Instance %v successfully updated\n", *id)
}

func HandleGetInstances() {
	c, err := admin.NewClient()
	utils.ExitIfError(err)

	instances, err := c.AdminGetAllInstances()
	utils.ExitIfError(err)

	if admin.GetDefaultFormat() == "json" {
		utils.PrintJson(instances)
	} else {
		utils.PrintMultiRow(admin.Instance{}, *instances)
	}
}

func HandleGetInstance(id *string) {
	c, err := admin.NewClient()
	utils.ExitIfError(err)

	instance, err := c.GetInstance(*id)
	utils.ExitIfError(err)

	if admin.GetDefaultFormat() == "json" {
		utils.PrintJson(instance)
	} else {
		utils.PrintRow(*instance)
	}
}
