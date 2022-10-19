package handlers

import (
	"cwc/client"
	"flag"
	"fmt"
	"os"
)

func HandleDeleteInstance(id *string) {

	client, err := client.NewClient()
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	err = client.DeleteInstance(*id)
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Instance %v successfully deleted\n", *id)
}

func HandleAttachInstance(attachCmd *flag.FlagSet, project_id *int, playbook *string, instance_type *string) {
	attachCmd.Parse(os.Args[3:])
	client, err := client.NewClient()
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	created_instance, err := client.AttachInstance(*project_id, *playbook, *instance_type)
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("ID\tcreated_at\tname\tstatus\tsize\tenvironment\tproject_id\n")
	fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\t%v\n", created_instance.Id, created_instance.CreatedAt, created_instance.Name, created_instance.Status, created_instance.Instance_type, created_instance.Environment, created_instance.Project)

}
func HandleAddInstance(name *string, project_id *int, project_name *string, env *string, instance_type *string, zone *string, dns_zone *string) {
	client, err := client.NewClient()
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	created_instance, err := client.AddInstance(*name, *project_id, *project_name, *instance_type, *env, *zone, *dns_zone)
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("ID\tcreated_at\tname\tstatus\tsize\tenvironment\tproject_id\n")
	fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\t%v\n", created_instance.Id, created_instance.CreatedAt, created_instance.Name, created_instance.Status, created_instance.Instance_type, created_instance.Environment, created_instance.Project)

}

func HandleUpdateInstance(id *string, status *string) {

	client, err := client.NewClient()
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	err = client.UpdateInstance(*id, *status)
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Instance %v successfully updated\n", *id)

}

func HandleGetInstance() {

	client, err := client.NewClient()
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	instances, err := client.GetAllInstances()

	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("ID\tcreated_at\tname\tstatus\tsize\tenvironment\tpublic ip\tproject_id\n")
	for _, instance := range *instances {
		fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\n", instance.Id, instance.CreatedAt, instance.Name, instance.Status, instance.Instance_type, instance.Environment, instance.Ip_address, instance.Project)

	}

	return
}
func HandleListInstancesTypes() {
	instancesTypes, err := client.GetInstancesTypes()
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)

	}
	for _, instance_type := range instancesTypes.Types {
		fmt.Printf("%v\n", instance_type)

	}
	return
}
