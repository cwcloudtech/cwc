package admin

import (
	"cwc/admin"
	"cwc/config"
	"cwc/utils"
	"fmt"
)

func HandleAddProject(user_email *string, name *string, host *string, token *string, git_username *string, namespace *string) {
	c, err := admin.NewClient()
	utils.ExitIfError(err)

	created_project, err := c.AdminAddProject(*user_email, *name, *host, *token, *git_username, *namespace)
	utils.ExitIfError(err)

	if config.GetDefaultFormat() == "json" {
		utils.PrintJson(created_project)
	} else {
		utils.PrintRow(*created_project)
	}
}

func HandleDeleteProject(id *string, name *string, url *string) {
	c, err := admin.NewClient()
	utils.ExitIfError(err)

	if utils.IsNotBlank(*id) {
		err := c.AdminDeleteProjectById(*id)
		utils.ExitIfError(err)
	} else if utils.IsNotBlank(*name) {
		err := c.AdminDeleteProjectByName(*name)
		utils.ExitIfError(err)
	} else {
		err := c.AdminDeleteProjectByUrl(*url)
		utils.ExitIfError(err)
	}

	fmt.Println("Project successfully deleted")
}

func HandleGetProjects(project_id *string, project_name *string, project_url *string) {
	var err error

	c, err := admin.NewClient()
	utils.ExitIfError(err)

	if utils.IsBlank(*project_id) && utils.IsBlank(*project_name) && utils.IsBlank(*project_url) {
		projects, err := c.AdminGetAllProjects()
		utils.ExitIfError(err)

		if config.GetDefaultFormat() == "json" {
			utils.PrintJson(projects)
		} else {
			utils.PrintMultiRow(admin.Project{}, *projects)
		}

		return
	}

	project := &admin.Project{}
	if utils.IsNotBlank(*project_id) {
		project, err = c.AdminGetProjectById(*project_id)
	} else if utils.IsNotBlank(*project_name) {
		project, err = c.AdminGetProjectByName(*project_name)
	} else {
		project, err = c.AdminGetProjectByUrl(*project_url)
	}

	utils.ExitIfError(err)

	if config.GetDefaultFormat() == "json" {
		utils.PrintJson(project)
	} else {
		utils.PrintRow(*project)
	}
}
