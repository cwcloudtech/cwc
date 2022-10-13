package handlers

import (
	"cwc/client"
	"flag"
	"fmt"
	"os"
)

func HandleAddProject(createCmd *flag.FlagSet, name *string, host *string, token *string, git_username *string, namespace *string) {
	createCmd.Parse(os.Args[3:])
	if *name == "" {
		fmt.Println("name is required to add a new project")
		createCmd.PrintDefaults()
		os.Exit(1)
	}
	client, _ := client.NewClient()
	created_project, err := client.AddProject(*name, *host, *token, *git_username, *namespace)
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("ID\tcreated_at\tname\turl\n")
	fmt.Printf("%v\t%v\t%v\t%v\n", created_project.Id, created_project.CreatedAt, created_project.Name, created_project.Url)

}

func HandleDeleteProject(deleteCmd *flag.FlagSet, id *string) {

	deleteCmd.Parse(os.Args[3:])
	if *id == "" {
		fmt.Println("id is required to delete your project")
		deleteCmd.PrintDefaults()
		os.Exit(1)
	}
	client, _ := client.NewClient()

	err := client.DeleteProject(*id)
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("project %v successfully deleted\n", *id)
}
func HandleGetProject(getCmd *flag.FlagSet, all *bool, id *string) {

	getCmd.Parse(os.Args[3:])
	if !*all && *id == "" {
		fmt.Println("id is required or specify --all to get all projects.")
		getCmd.PrintDefaults()
		os.Exit(1)
	}
	client, _ := client.NewClient()
	if *all {

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

	if *id != "" {
		project, err := client.GetProject(*id)
		if err != nil {
			fmt.Printf("failed: %s\n", err)
			os.Exit(1)
		}
		fmt.Printf("ID\tcreated_at\tname\tcreated_at\turl\n")
		fmt.Printf("%v\t%v\t%v\t%v\n", project.Id, project.CreatedAt, project.Name, project.Url)

		return
	}
}
