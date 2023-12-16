package user

import (
	"cwc/client"
	"cwc/utils"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

func HandleGetEnvironments(pretty *bool) {
	c, err := client.NewClient()
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	environments, err := c.GetAllEnvironments()
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	if client.GetDefaultFormat() == "json" {
		utils.PrintJson(environments)
	} else if *pretty {
		displayEnvironmentsAsTable(*environments)
	} else {
		utils.PrintMultiRow(client.Environment{}, *environments)
	}
}

func HandleGetEnvironment(id *string, pretty *bool) {
	c, err := client.NewClient()
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	environment, err := c.GetEnvironment(*id)
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	if client.GetDefaultFormat() == "json" {
		utils.PrintJson(environment)
	} else if *pretty {
		utils.PrintPretty("Environment found", *environment)
	} else {
		utils.PrintRow(environment)
	}
}

func displayEnvironmentsAsTable(environments []client.Environment) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "Path", "Description"})

	if len(environments) == 0 {
		fmt.Println("No environments found")
		return
	} else {
		for _, environment := range environments {
			table.Append([]string{
				fmt.Sprintf("%d", environment.Id),
				environment.Name,
				environment.Path,
				environment.Description,
			})
		}
	}

	table.Render()
}
