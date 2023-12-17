package user

import (
	"cwc/client"
	"cwc/utils"
	"flag"
	"fmt"
	"os"
)

func HandleDeleteInstance(id *string) {
	c, err := client.NewClient()
	utils.ExitIfError(err)

	err = c.DeleteInstance(*id)
	utils.ExitIfError(err)

	fmt.Printf("Instance %v successfully deleted\n", *id)
}

func HandleAttachInstance(attachCmd *flag.FlagSet, project_id *int, playbook *string, instance_type *string) {
	attachCmd.Parse(os.Args[3:])
	c, err := client.NewClient()
	utils.ExitIfError(err)

	created_instance, err := c.AttachInstance(*project_id, *playbook, *instance_type)
	utils.ExitIfError(err)

	fmt.Printf("ID\tcreated_at\tname\tstatus\tsize\tenvironment\tproject_id\n")
	fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\t%v\n", created_instance.Id, created_instance.CreatedAt, created_instance.Name, created_instance.Status, created_instance.Instance_type, created_instance.Environment, created_instance.Project)
}

func HandleAddInstance(name *string, project_id *int, project_name *string, project_url *string, env *string, instance_type *string, zone *string, dns_zone *string) {
	c, err := client.NewClient()
	utils.ExitIfError(err)

	created_instance, err := c.AddInstance(*name, *project_id, *project_name, *project_url, *instance_type, *env, *zone, *dns_zone)
	utils.ExitIfError(err)

	fmt.Printf("ID\tcreated_at\tname\tstatus\tsize\tenvironment\tproject_id\n")
	fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\t%v\n", created_instance.Id, created_instance.CreatedAt, created_instance.Name, created_instance.Status, created_instance.Instance_type, created_instance.Environment, created_instance.Project)
}

func HandleUpdateInstance(id *string, status *string) {
	c, err := client.NewClient()
	utils.ExitIfError(err)

	err = c.UpdateInstance(*id, *status)
	utils.ExitIfError(err)

	fmt.Printf("Instance %v successfully updated\n", *id)
}

func HandleGetInstances() {
	c, err := client.NewClient()
	utils.ExitIfError(err)

	instances, err := c.GetAllInstances()
	utils.ExitIfError(err)

	if client.GetDefaultFormat() == "json" {
		utils.PrintJson(instances)
	} else {
		utils.PrintMultiRow(client.Instance{}, *instances)
	}
}

func HandleListInstancesTypes() {
	instancesTypes, err := client.GetInstancesTypes()
	utils.ExitIfError(err)

	for _, instance_type := range instancesTypes.Types {
		fmt.Printf("%v\n", instance_type)
	}
}

func HandleGetInstance(id *string) {
	c, err := client.NewClient()
	utils.ExitIfError(err)

	instance, err := c.GetInstance(*id)
	utils.ExitIfError(err)

	if client.GetDefaultFormat() == "json" {
		utils.PrintJson(instance)
	} else {
		utils.PrintRow(*instance)
	}
}
