package handlers

import (
	"cwc/client"
	"fmt"
	"os"
)

func HandleAddProject(name *string, host *string, token *string, git_username *string, namespace *string) {

	client, _ := client.NewClient()
	created_project, err := client.AddProject(*name, *host, *token, *git_username, *namespace)
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("ID\tcreated_at\tname\turl\n")
	fmt.Printf("%v\t%v\t%v\t%v\n", created_project.Id, created_project.CreatedAt, created_project.Name, created_project.Url)

}

func HandleDeleteProject(id *string) {
	client, _ := client.NewClient()

	err := client.DeleteProject(*id)
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("project %v successfully deleted\n", *id)
}
func HandleGetProject() {
	client, _ := client.NewClient()

	projects, err := client.GetAllProjects()

	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("ID\tcreated_at\tname\turl\n")
	for _, project := range *projects {
		fmt.Printf("%v\t%v\t%v\t%v\n", project.Id, project.CreatedAt, project.Name, project.Url)
	}

	return
}
