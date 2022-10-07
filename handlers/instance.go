package handlers

import (
	"cwc/client"
	"flag"
	"fmt"
	"os"
)

func HandleDeleteInstance(deleteCmd *flag.FlagSet, id *string) {

	deleteCmd.Parse(os.Args[3:])
	if *id == "" {
		fmt.Println("id is required to delete your instance")
		deleteCmd.PrintDefaults()
		os.Exit(1)
	}
	client := client.NewClient()
	err := client.DeleteInstance(*id)
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Instance %v successfully deleted\n", *id)
}
func ValidateInstance(createCmd *flag.FlagSet, name *string, env *string) {
	if *name == "" || *env == "" {
		createCmd.PrintDefaults()
		os.Exit(1)
	}
}

func HandleAttachInstance(attachCmd *flag.FlagSet, project_id *int, playbook *string, instance_type *string) {
	attachCmd.Parse(os.Args[3:])
	client := client.NewClient()
	created_instance, err := client.AttachInstance(*project_id, *playbook, *instance_type)
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("ID\tname\tstatus\tsize\tenvironment\tgitlab url\n")
	fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\n", created_instance.Id, created_instance.Name, created_instance.Status, created_instance.Instance_type, created_instance.Environment, created_instance.Project)

}
func HandleAddInstance(createCmd *flag.FlagSet, name *string, project_id *int, env *string, instance_type *string, zone *string) {
	createCmd.Parse(os.Args[3:])
	ValidateInstance(createCmd, name, env)
	client := client.NewClient()
	created_instance, err := client.AddInstance(*name, *project_id, *instance_type, *env, *zone)
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("ID\tname\tstatus\tsize\tenvironment\tgitlab url\n")
	fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\n", created_instance.Id, created_instance.Name, created_instance.Status, created_instance.Instance_type, created_instance.Environment, created_instance.Project)

}

func HandleUpdateInstance(updateCmd *flag.FlagSet, id *string, status *string) {
	updateCmd.Parse(os.Args[3:])
	if *id == "" || *status == "" {
		fmt.Println("id and status are required")
		updateCmd.PrintDefaults()
		os.Exit(1)
	}
	client := client.NewClient()
	err := client.UpdateInstance(*id, *status)
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Project %v successfully updated\n", *id)

}

func HandleGetInstance(getCmd *flag.FlagSet, all *bool, id *string) {

	getCmd.Parse(os.Args[3:])
	if !*all && *id == "" {
		fmt.Println("id is required or specify --all to get all instances.")
		getCmd.PrintDefaults()
		os.Exit(1)
	}
	client := client.NewClient()
	if *all {

		instances, err := client.GetAllInstances()

		if err != nil {
			fmt.Printf("failed: %s\n", err)
			os.Exit(1)
		}

		fmt.Printf("ID\tname\tstatus\tsize\tenvironment\tpublic ip\tgitlab url\n")
		for _, instance := range *instances {
			fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\t%v\n", instance.Id, instance.Name, instance.Status, instance.Instance_type, instance.Environment, instance.Ip_address, instance.Project)

		}

		return
	}

	if *id != "" {
		instance, err := client.GetInstance(*id)
		if err != nil {
			fmt.Printf("failed: %s\n", err)
			os.Exit(1)
		}
		fmt.Printf("ID\tname\tstatus\tsize\tenvironment\tpublic ip\tgitlab url\n")
		fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\t%v\n", instance.Id, instance.Name, instance.Status, instance.Instance_type, instance.Environment, instance.Ip_address, instance.Project)

		return
	}
}
