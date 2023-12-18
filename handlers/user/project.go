package user

import (
	"cwc/client"
	"cwc/config"
	"cwc/utils"
	"fmt"
)

func HandleAddProject(name *string, host *string, token *string, git_username *string, namespace *string) {
	c, err := client.NewClient()
	utils.ExitIfError(err)

	created_project, err := c.AddProject(*name, *host, *token, *git_username, *namespace)
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

func HandleGetProjects(project_id *string, project_name *string, project_url *string) {
	var err error

	c, err := client.NewClient()
	utils.ExitIfError(err)

	if utils.IsBlank(*project_id) && utils.IsBlank(*project_name) && utils.IsBlank(*project_url) {
		projects, err := c.GetAllProjects()
		utils.ExitIfError(err)

		if config.GetDefaultFormat() == "json" {
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

		if config.GetDefaultFormat() == "json" {
			utils.PrintJson(project)
		} else {
			utils.PrintRow(*project)
		}
	}
}
