package user

import (
	"cwc/client"
	"cwc/config"
	"cwc/utils"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

func HandleGetDeployments(deployments *[]client.Deployment, pretty *bool) {
	if config.IsPrettyFormatExpected(pretty) {
		displayDeploymentsAsTable(*deployments)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(deployments)
	} else {
		var deploymentsDisplay []client.DeploymentDisplay
		for i, deployment := range *deployments {
			deploymentsDisplay = append(deploymentsDisplay, client.DeploymentDisplay{
				Name:       deployment.Name,
				Namespace:  deployment.Namespace,
				Created_at: deployment.Created_at,
			})
			deploymentsDisplay[i].Id = deployment.Id
		}
		utils.PrintMultiRow(client.DeploymentDisplay{}, deploymentsDisplay)
	}
}

func HandleGetDeployment(deployment *client.DeploymentByIdResponse, pretty *bool) {
	var deploymentDisplay client.DeploymentDisplay
	deploymentDisplay.Name = deployment.Name
	deploymentDisplay.Namespace = deployment.Namespace

	if config.IsPrettyFormatExpected(pretty) {
		utils.PrintPretty("Found deployment", deployment)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(deployment)
	} else {
		utils.PrintRow(deploymentDisplay)
	}
}

func PrepareAddDeployment(deployment *client.CreationDeployment) (*client.CreationDeployment, error) {
	c, err := client.NewClient()
	utils.ExitIfError(err)

	created_deployment, err := c.CreateDeployment(*deployment)
	utils.ExitIfError(err)

	return created_deployment, nil
}

func HandleAddDeployment(created_deployment *client.CreationDeployment, pretty *bool) {
	if created_deployment == nil {
		fmt.Println("Error: Created deployment is nil")
		return
	}
	if config.IsPrettyFormatExpected(pretty) {
		utils.PrintPretty("Created deployment", created_deployment)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(created_deployment)
	} else {
		utils.PrintRow(created_deployment)
	}
}

func HandleDeleteDeployment(deploymentId *string) {
	c, err := client.NewClient()
	utils.ExitIfError(err)

	err = c.DeleteDeploymentById(*deploymentId)
	utils.ExitIfError(err)

	fmt.Println("Deployment deleted successfully")
}

func displayDeploymentsAsTable(deployments []client.Deployment) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "Namespace", "Created_at"})

	if len(deployments) == 0 {
		table.Append([]string{"No deployments found", "404", "404", "404"})
	} else {
		for _, deployment := range deployments {
			table.Append([]string{
				deployment.Id,
				deployment.Name,
				deployment.Namespace,
				deployment.Created_at,
			})
		}
	}
	table.Render()
}
