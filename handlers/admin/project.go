package admin

import (
	"cwc/admin"

	"fmt"
	"os"
)

func HandleAddProject(user_email *string, name *string, host *string, token *string, git_username *string, namespace *string) {

	admin, _ := admin.NewClient()
	created_project, err := admin.AdminAddProject(*user_email, *name, *host, *token, *git_username, *namespace)
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("ID\tcreated_at\tname\turl\n")
	fmt.Printf("%v\t%v\t%v\t%v\n", created_project.Id, created_project.CreatedAt, created_project.Name, created_project.Url)

}

func HandleDeleteProject(id *string, name *string, url *string) {
	admin, _ := admin.NewClient()
	if *id != "" {
		err := admin.AdminDeleteProjectById(*id)
		if err != nil {
			fmt.Printf("failed: %s\n", err)
			os.Exit(1)
		}
	} else if *name != "" {
		err := admin.AdminDeleteProjectByName(*name)
		if err != nil {
			fmt.Printf("failed: %s\n", err)
			os.Exit(1)
		}
	} else {
		err := admin.AdminDeleteProjectByUrl(*url)
		if err != nil {
			fmt.Printf("failed: %s\n", err)
			os.Exit(1)
		}
	}
	fmt.Printf("project successfully deleted\n")
}
func HandleGetProjects(project_id *string, project_name *string, project_url *string) {
	c, _ := admin.NewClient()
	var err error
	if *project_id == "" && *project_name == "" && *project_url == "" {
		projects, err := c.AdminGetAllProjects()

		if err != nil {
			fmt.Printf("failed: %s\n", err)
			os.Exit(1)
		}

		fmt.Printf("ID\tcreated_at\tname\turl\n")
		for _, project := range *projects {
			fmt.Printf("%v\t%v\t%v\t%v\n", project.Id, project.CreatedAt, project.Name, project.Url)
		}

	} else {
		project := &admin.Project{}
		if *project_id != "" {
			project, err = c.AdminGetProjectById(*project_id)
		} else if *project_name != "" {
			project, err = c.AdminGetProjectByName(*project_name)

		} else {
			project, err = c.AdminGetProjectByUrl(*project_url)
		}

		if err != nil {
			fmt.Printf("failed: %s\n", err)
			os.Exit(1)
		}
		fmt.Printf("ID\tcreated_at\tname\tcreated_at\turl\n")
		fmt.Printf("%v\t%v\t%v\t%v\n", project.Id, project.CreatedAt, project.Name, project.Url)

	}

	return
}
