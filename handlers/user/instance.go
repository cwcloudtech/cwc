package user

import (
	"cwc/client"
	"cwc/config"
	"cwc/utils"
	"flag"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
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

func HandleGetInstances(instances *[]client.Instance, pretty *bool) {
	if config.IsPrettyFormatExpected(pretty) {
		displayInstancesAsTable(*instances)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(instances)
	} else {
		utils.PrintMultiRow(client.Instance{}, *instances)
	}
}

func HandleListInstancesTypes(instances_types *client.InstancesTypes, pretty *bool) {
	if config.IsPrettyFormatExpected(pretty) {
		utils.PrintPrettyArray("Types of instances available", instances_types.Types)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(instances_types)
	} else {
		utils.PrintArray(instances_types.Types)
	}
}

func HandleGetInstance(instance *client.Instance, pretty *bool) {
	if config.IsPrettyFormatExpected(pretty) {
		utils.PrintPretty("Instance's informations", *instance)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(instance)
	} else {
		utils.PrintRow(*instance)
	}
}

func displayInstancesAsTable(instances []client.Instance) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "IP", "Type", "Created at"})

	if len(instances) == 0 {
		fmt.Println("No instances found")
	} else {
		for _, instance := range instances {
			table.Append([]string{
				fmt.Sprintf("%d", instance.Id),
				instance.Name,
				instance.Ip_address,
				instance.Instance_type,
				instance.CreatedAt,
			})
		}
	}

	table.Render()
}
