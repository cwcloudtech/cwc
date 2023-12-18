package admin

import (
	"cwc/admin"
	"cwc/config"
	"cwc/utils"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
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

	if config.GetDefaultFormat() == "json" {
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

func HandleGetInstances(pretty *bool) {
	c, err := admin.NewClient()
	utils.ExitIfError(err)

	instances, err := c.AdminGetAllInstances()
	utils.ExitIfError(err)

	if config.IsPrettyFormatExpected(pretty) {
		displayInstanceAsTable(*instances)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(instances)
	} else {
		utils.PrintMultiRow(admin.Instance{}, *instances)
	}
}

func HandleGetInstance(id *string, pretty *bool) {
	c, err := admin.NewClient()
	utils.ExitIfError(err)

	instance, err := c.GetInstance(*id)
	utils.ExitIfError(err)

	if config.IsPrettyFormatExpected(pretty) {
		utils.PrintPretty("Instance's informations", *instance)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(instance)
	} else {
		utils.PrintRow(*instance)
	}
}

func displayInstanceAsTable(instances []admin.Instance) {
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
