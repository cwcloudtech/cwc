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

func HandleGet(getCmd *flag.FlagSet, all *bool, id *string) {

	getCmd.Parse(os.Args[2:])
	if !*all && *id == "" {
		fmt.Println("id is required or specify --all to get all instances.")
		getCmd.PrintDefaults()
		os.Exit(1)
	}
	client := client.NewClient()
	if *all {

		projects, err := client.GetAll()

		if err != nil {
			fmt.Printf("failed: %s\n", err)
			os.Exit(1)
		}

		fmt.Printf("ID\tname\tstatus\tsize\tenvironment\tpublic ip\tgitlab url\n")
		for _, project := range *projects {
			fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\t%v\n", project.Id, project.Name, project.Status, project.Instance_type, project.Environment, project.Ip_address, project.Gitlab_url)

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
		fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\t%v\n", project.Id, project.Name, project.Status, project.Instance_type, project.Environment, project.Ip_address, project.Gitlab_url)

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

func HandleDelete(deleteCmd *flag.FlagSet, id *string) {

	deleteCmd.Parse(os.Args[2:])
	if *id == "" {
		fmt.Println("id is required to delete your instance")
		deleteCmd.PrintDefaults()
		os.Exit(1)
	}
	client := client.NewClient()

	err := client.DeleteProject(*id)
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Project %v successfully deleted\n", *id)
}
func ValidateProject(createCmd *flag.FlagSet, name *string, env *string) {

	if *name == "" || *env == "" {
		createCmd.PrintDefaults()
		os.Exit(1)
	}
}
func HandleAdd(createCmd *flag.FlagSet, name *string, email *string, env *string, instance_type *string) {
	createCmd.Parse(os.Args[2:])
	ValidateProject(createCmd, name, env)
	client := client.NewClient()
	created_project, err := client.AddProject(*name, *instance_type, *env, *email)
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("ID\tname\tstatus\tsize\tenvironment\tgitlab url\n")
	fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\n", created_project.Id, created_project.Name, created_project.Status, created_project.Instance_type, created_project.Environment, created_project.Gitlab_url)

}

func HandleUpdate(updateCmd *flag.FlagSet, id *string, status *string) {
	updateCmd.Parse(os.Args[2:])
	if *id == "" || *status == "" {
		fmt.Println("id and status are required")
		updateCmd.PrintDefaults()
		os.Exit(1)
	}
	client := client.NewClient()
	err := client.UpdateProject(*id, *status)
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Project %v successfully updated\n", *id)

}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
