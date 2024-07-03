package user

import (
	"cwc/client"
	"cwc/config"
	"cwc/utils"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

func HandleAddProject(name *string, host *string, token *string, git_username *string, namespace *string, project_type *string) {
	c, err := client.NewClient()
	utils.ExitIfError(err)

	created_project, err := c.AddProject(*name, *host, *token, *git_username, *namespace, *project_type)
	utils.ExitIfError(err)

	if config.GetDefaultFormat() == "json" {
		utils.PrintJson(created_project)
	} else {
		utils.PrintRow(*created_project)
	}
}

func HandleDeleteProject(id *string, name *string, url *string) {
	c, err := client.NewClient()
	utils.ExitIfError(err)

	if utils.IsNotBlank(*id) {
		err := c.DeleteProjectById(*id)
		utils.ExitIfError(err)
	} else if utils.IsNotBlank(*name) {
		err := c.DeleteProjectByName(*name)
		utils.ExitIfError(err)
	} else {
		err := c.DeleteProjectByUrl(*url)
		utils.ExitIfError(err)
	}

	fmt.Println("project successfully deleted")
}

func HandleGetProjects(project_id *string, project_name *string, project_url *string, pretty *bool, project_type *string) {
	var err error

	c, err := client.NewClient()
	utils.ExitIfError(err)

	if utils.IsBlank(*project_id) && utils.IsBlank(*project_name) && utils.IsBlank(*project_url) {
		projects, err := c.GetAllProjects(*project_type)
		utils.ExitIfError(err)

		if config.IsPrettyFormatExpected(pretty) {
			displayProjectsAsTable(*projects)
		} else if config.GetDefaultFormat() == "json" {
			utils.PrintJson(projects)
		} else {
			utils.PrintMultiRow(client.Project{}, *projects)
		}
	} else {
		project := &client.Project{}
		if utils.IsNotBlank(*project_id) {
			project, err = c.GetProjectById(*project_id)
		} else if utils.IsNotBlank(*project_name) {
			project, err = c.GetProjectByName(*project_name)
		} else {
			project, err = c.GetProjectByUrl(*project_url)
		}

		utils.ExitIfError(err)

		if config.IsPrettyFormatExpected(pretty) {
			utils.PrintPretty("Project's informations", *project)
		} else if config.GetDefaultFormat() == "json" {
			utils.PrintJson(project)
		} else {
			utils.PrintRow(*project)
		}
	}
}

func displayProjectsAsTable(projects []client.Project) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "Type", "URL", "Created at"})

	if len(projects) == 0 {
		fmt.Println("No projects found")
	} else {
		for _, project := range projects {
			table.Append([]string{
				fmt.Sprintf("%d", project.Id),
				project.Name,
				project.Type,
				project.Url,
				project.CreatedAt,
			})
		}
		table.Render()
	}
}
