package handlers

import (
	"cwc/client"
	"flag"
	"fmt"
	"os"
	"runtime"
)

func HandleVersion(versionCmd *flag.FlagSet, version string) {

	versionCmd.Parse(os.Args[2:])
	fmt.Printf("cwc-cli/%v %v %v\n", version, runtime.GOOS, runtime.GOARCH)
}

func HandleHelp(helpCmd *flag.FlagSet) {

	helpCmd.Parse(os.Args[2:])
	fmt.Printf("cwc: available commands:\n\n\n")
	fmt.Printf("- create \n")
	fmt.Printf("  create a new instance\n\n")
	fmt.Printf("- get \n")
	fmt.Printf("  get one or many instances\n\n")
	fmt.Printf("- delete \n")
	fmt.Printf("  delete an existing instance\n\n")
	fmt.Printf("- update \n")
	fmt.Printf("  update a particular instance state\n\n")
	fmt.Printf("- configure \n")
	fmt.Printf("  configure your default settings\n\n")

}

func HandleGetInstance(getCmd *flag.FlagSet, all *bool, id *string) {

	getCmd.Parse(os.Args[3:])
	if !*all && *id == "" {
		fmt.Println("id is required or specify --all to get all instances.")
		getCmd.PrintDefaults()
		os.Exit(1)
	}
	client := client.NewClient()
	if *all {

		instances, err := client.GetAllInstances()

		if err != nil {
			fmt.Printf("failed: %s\n", err)
			os.Exit(1)
		}

		fmt.Printf("ID\tname\tstatus\tsize\tenvironment\tpublic ip\tgitlab url\n")
		for _, instance := range *instances {
			fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\t%v\n", instance.Id, instance.Name, instance.Status, instance.Instance_type, instance.Environment, instance.Ip_address, instance.Project)

		}

		return
	}

	if *id != "" {
		instance, err := client.GetInstance(*id)
		if err != nil {
			fmt.Printf("failed: %s\n", err)
			os.Exit(1)
		}
		fmt.Printf("ID\tname\tstatus\tsize\tenvironment\tpublic ip\tgitlab url\n")
		fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\t%v\n", instance.Id, instance.Name, instance.Status, instance.Instance_type, instance.Environment, instance.Ip_address, instance.Project)

		return
	}
}

func HandleLogin(loginCmd *flag.FlagSet, email *string, password *string) {

	loginCmd.Parse(os.Args[2:])
	if *email == "" || *password == "" {
		fmt.Println("email and password are required to login")
		loginCmd.PrintDefaults()
		os.Exit(1)
	}
	client := client.NewClient()

	err := client.UserLogin(*email, *password)
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("You are successfully logged in\n")
}

func HandleConfigure(configureCmd *flag.FlagSet, region *bool) {

	configureCmd.Parse(os.Args[2:])
	if !*region {
		fmt.Println("cwc: flag is missing")
		configureCmd.PrintDefaults()
		os.Exit(1)
	}
	if len(os.Args) < 4 {
		fmt.Println("cwc: invalid arguments")
		fmt.Println("cwc configure <command> <subcommand> <value>")
		fmt.Println("<command>: -region")
		fmt.Println("<subcommand>: get / set")
		os.Exit(1)
	}
	subSubCommmand := os.Args[3]
	switch subSubCommmand {

	case "set":
		region_value := os.Args[4]
		if region_value == "" {
			fmt.Println("cwc: region value is missing")
			configureCmd.PrintDefaults()
			os.Exit(1)
		}
		possible_regions := make([]string, 4)
		possible_regions[0] = "fr-par-1"
		possible_regions[1] = "fr-par-2"
		possible_regions[2] = "nl-ams-1"
		possible_regions[3] = "pl-waw-1"

		if !stringInSlice(region_value, possible_regions) {
			fmt.Println("cwc: invalid region value")
			os.Exit(1)

		}

		client.SetDefaultRegion(region_value)
		fmt.Printf("Default region = %v\n", region_value)
	case "get":
		region := client.GetDefaultRegion()
		fmt.Printf("Default region = %v\n", region)
	default:
		fmt.Printf("cwc: option not found")
	}

}

func HandleDeleteInstance(deleteCmd *flag.FlagSet, id *string) {

	deleteCmd.Parse(os.Args[3:])
	if *id == "" {
		fmt.Println("id is required to delete your instance")
		deleteCmd.PrintDefaults()
		os.Exit(1)
	}
	client := client.NewClient()

	err := client.DeleteInstance(*id)
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Instance %v successfully deleted\n", *id)
}
func ValidateInstance(createCmd *flag.FlagSet, name *string, env *string) {
	fmt.Printf(*name)
	if *name == "" || *env == "" {
		createCmd.PrintDefaults()
		os.Exit(1)
	}
}
func HandleAddInstance(createCmd *flag.FlagSet, name *string, project_id *int, env *string, instance_type *string) {
	createCmd.Parse(os.Args[3:])
	ValidateInstance(createCmd, name, env)
	client := client.NewClient()
	created_instance, err := client.AddInstance(*name, *project_id, *instance_type, *env)
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("ID\tname\tstatus\tsize\tenvironment\tgitlab url\n")
	fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\n", created_instance.Id, created_instance.Name, created_instance.Status, created_instance.Instance_type, created_instance.Environment, created_instance.Project)

}

func HandleUpdateInstance(updateCmd *flag.FlagSet, id *string, status *string) {
	updateCmd.Parse(os.Args[3:])
	if *id == "" || *status == "" {
		fmt.Println("id and status are required")
		updateCmd.PrintDefaults()
		os.Exit(1)
	}
	client := client.NewClient()
	err := client.UpdateInstance(*id, *status)
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Project %v successfully updated\n", *id)

}

func HandleAddProject(createCmd *flag.FlagSet, name *string) {
	createCmd.Parse(os.Args[3:])
	client := client.NewClient()
	created_project, err := client.AddProject(*name)
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("ID\tname\tcreated_at\turl\n")
	fmt.Printf("%v\t%v\t%v\t%v\n", created_project.Id, created_project.Name, created_project.CreatedAt, created_project.Url)

}

func HandleDeleteProject(deleteCmd *flag.FlagSet, id *string) {

	deleteCmd.Parse(os.Args[3:])
	if *id == "" {
		fmt.Println("id is required to delete your project")
		deleteCmd.PrintDefaults()
		os.Exit(1)
	}
	client := client.NewClient()

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
	client := client.NewClient()
	if *all {

		projects, err := client.GetAllProjects()

		if err != nil {
			fmt.Printf("failed: %s\n", err)
			os.Exit(1)
		}

		fmt.Printf("ID\tname\tcreated_at\turl\n")
		for _, project := range *projects {
			fmt.Printf("%v\t%v\t%v\t%v\n", project.Id, project.Name, project.CreatedAt, project.Url)
		}

		return
	}

	if *id != "" {
		project, err := client.GetProject(*id)
		if err != nil {
			fmt.Printf("failed: %s\n", err)
			os.Exit(1)
		}
		fmt.Printf("ID\tname\tstatus\tsize\tenvironment\tpublic ip\tgitlab url\n")
		fmt.Printf("%v\t%v\t%v\t%v\n", project.Id, project.Name, project.CreatedAt, project.Url)

		return
	}
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
