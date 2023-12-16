package user

import (
	"cwc/client"
	"cwc/utils"
	"flag"
	"fmt"
	"os"
)

func HandleDeleteInstance(id *string) {
	client, err := client.NewClient()
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	err = client.DeleteInstance(*id)
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Instance %v successfully deleted\n", *id)
}

func HandleAttachInstance(attachCmd *flag.FlagSet, project_id *int, playbook *string, instance_type *string) {
	attachCmd.Parse(os.Args[3:])
	client, err := client.NewClient()
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	created_instance, err := client.AttachInstance(*project_id, *playbook, *instance_type)
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("ID\tcreated_at\tname\tstatus\tsize\tenvironment\tproject_id\n")
	fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\t%v\n", created_instance.Id, created_instance.CreatedAt, created_instance.Name, created_instance.Status, created_instance.Instance_type, created_instance.Environment, created_instance.Project)

}
func HandleAddInstance(name *string, project_id *int, project_name *string, project_url *string, env *string, instance_type *string, zone *string, dns_zone *string) {
	client, err := client.NewClient()
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	created_instance, err := client.AddInstance(*name, *project_id, *project_name, *project_url, *instance_type, *env, *zone, *dns_zone)
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("ID\tcreated_at\tname\tstatus\tsize\tenvironment\tproject_id\n")
	fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\t%v\n", created_instance.Id, created_instance.CreatedAt, created_instance.Name, created_instance.Status, created_instance.Instance_type, created_instance.Environment, created_instance.Project)

}

func HandleUpdateInstance(id *string, status *string) {

	client, err := client.NewClient()
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	err = client.UpdateInstance(*id, *status)
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Instance %v successfully updated\n", *id)

}

func HandleGetInstances() {

	c, err := client.NewClient()
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	instances, err := c.GetAllInstances()

	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	if client.GetDefaultFormat() == "json" {
		utils.PrintJson(instances)
	} else {
		utils.PrintMultiRow(client.Instance{}, *instances)
	}

	return
}
func HandleListInstancesTypes() {
	instancesTypes, err := client.GetInstancesTypes()
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)

	}
	for _, instance_type := range instancesTypes.Types {
		fmt.Printf("%v\n", instance_type)

	}
	return
}

func HandleGetInstance(id *string) {

	c, err := client.NewClient()
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	instance, err := c.GetInstance(*id)

	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	if client.GetDefaultFormat() == "json" {
		utils.PrintJson(instance)
	} else {
		utils.PrintRow(*instance)
	}
	return
}
