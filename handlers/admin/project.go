package admin

import (
	"cwc/admin"
	"cwc/utils"
	"fmt"
)

func HandleAddProject(user_email *string, name *string, host *string, token *string, git_username *string, namespace *string) {
	c, err := admin.NewClient()
	utils.ExitIfError(err)

	created_project, err := c.AdminAddProject(*user_email, *name, *host, *token, *git_username, *namespace)
	utils.ExitIfError(err)

	if admin.GetDefaultFormat() == "json" {
		utils.PrintJson(created_project)
	} else {
		utils.PrintRow(*created_project)
	}
}

func HandleDeleteProject(id *string, name *string, url *string) {
	c, err := admin.NewClient()
	utils.ExitIfError(err)

	if *id != "" {
		err := c.AdminDeleteProjectById(*id)
		utils.ExitIfError(err)
	} else if *name != "" {
		err := c.AdminDeleteProjectByName(*name)
		utils.ExitIfError(err)
	} else {
		err := c.AdminDeleteProjectByUrl(*url)
		utils.ExitIfError(err)
	}

	fmt.Printf("project successfully deleted\n")
}

func HandleGetProjects(project_id *string, project_name *string, project_url *string) {
	var err error

	c, err := admin.NewClient()
	utils.ExitIfError(err)

	if *project_id == "" && *project_name == "" && *project_url == "" {
		projects, err := c.AdminGetAllProjects()
		utils.ExitIfError(err)

		if admin.GetDefaultFormat() == "json" {
			utils.PrintJson(projects)
		} else {
			utils.PrintMultiRow(admin.Project{}, *projects)
		}

		return
	}

	project := &admin.Project{}
	if *project_id != "" {
		project, err = c.AdminGetProjectById(*project_id)
	} else if *project_name != "" {
		project, err = c.AdminGetProjectByName(*project_name)
	} else {
		project, err = c.AdminGetProjectByUrl(*project_url)
	}

	utils.ExitIfError(err)

	if admin.GetDefaultFormat() == "json" {
		utils.PrintJson(project)
	} else {
		utils.PrintRow(*project)
	}
}
