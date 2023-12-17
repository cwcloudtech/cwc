package user

import (
	"cwc/client"
	"cwc/utils"
	"fmt"
	"os"
)

func HandleAddProject(name *string, host *string, token *string, git_username *string, namespace *string) {
	c, err := client.NewClient()
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	created_project, err := c.AddProject(*name, *host, *token, *git_username, *namespace)
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	if client.GetDefaultFormat() == "json" {
		utils.PrintJson(created_project)
	} else {
		utils.PrintRow(*created_project)
	}
}

func HandleDeleteProject(id *string, name *string, url *string) {
	c, err := client.NewClient()
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	if *id != "" {
		err := c.DeleteProjectById(*id)
		if nil != err {
			fmt.Printf("failed: %s\n", err)
			os.Exit(1)
		}
	} else if *name != "" {
		err := c.DeleteProjectByName(*name)
		if nil != err {
			fmt.Printf("failed: %s\n", err)
			os.Exit(1)
		}
	} else {
		err := c.DeleteProjectByUrl(*url)
		if nil != err {
			fmt.Printf("failed: %s\n", err)
			os.Exit(1)
		}
	}

	fmt.Printf("project successfully deleted\n")
}

func HandleGetProjects(project_id *string, project_name *string, project_url *string) {
	var err error

	c, err := client.NewClient()
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	if *project_id == "" && *project_name == "" && *project_url == "" {
		projects, err := c.GetAllProjects()

		if nil != err {
			fmt.Printf("failed: %s\n", err)
			os.Exit(1)
		}

		if client.GetDefaultFormat() == "json" {
			utils.PrintJson(projects)
		} else {
			utils.PrintMultiRow(client.Project{}, *projects)
		}
	} else {
		project := &client.Project{}
		if *project_id != "" {
			project, err = c.GetProjectById(*project_id)
		} else if *project_name != "" {
			project, err = c.GetProjectByName(*project_name)
		} else {
			project, err = c.GetProjectByUrl(*project_url)
		}

		if nil != err {
			fmt.Printf("failed: %s\n", err)
			os.Exit(1)
		}

		if client.GetDefaultFormat() == "json" {
			utils.PrintJson(project)
		} else {
			utils.PrintRow(*project)
		}
	}
}
