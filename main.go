package main

import (
	// "fmt"
	"cwc/handlers"
	"flag"
	"fmt"
	"os"
)

var Version = "dev"

func main() {

	helpCmd := flag.NewFlagSet("help", flag.ExitOnError)
	versionCmd := flag.NewFlagSet("version", flag.ExitOnError)

	// instance handlers
	getInstanceCmd := flag.NewFlagSet("get instance", flag.ExitOnError)
	getAllInstances := getInstanceCmd.Bool("all", false, "Get all instances")
	getInstanceById := getInstanceCmd.String("id", "", "Get instance by ID")

	createInstanceCmd := flag.NewFlagSet("create instance", flag.ExitOnError)
	addInstanceName := createInstanceCmd.String("name", "", "The instance name")
	addInstanceProjectId := createInstanceCmd.Int("project_id", 0, "The project id that you want to associete with the instance")
	addInstanceEnvironment := createInstanceCmd.String("env", "", "The environment of the instance (code, wpaas)")
	addInstanceType := createInstanceCmd.String("instance_type", "", "The instance size (DEV1-S, DEV1-M, DEV1-L, DEV1-XL)")

	attachInstanceCmd := flag.NewFlagSet("attach instance", flag.ExitOnError)
	attachInstancePlaybook := attachInstanceCmd.String("name", "", "The playbook name you want to run")
	attachInstanceProjectId := attachInstanceCmd.Int("project_id", 0, "The project id that you want to attach with the instance")
	attachInstanceType := attachInstanceCmd.String("instance_type", "", "The instance size (DEV1-S, DEV1-M, DEV1-L, DEV1-XL)")

	deleteInstanceCmd := flag.NewFlagSet("delete instance", flag.ExitOnError)
	deleteInstanceById := deleteInstanceCmd.String("id", "", "Target instance ID")

	updateInstanceCmd := flag.NewFlagSet("update instance", flag.ExitOnError)
	updateInstanceId := updateInstanceCmd.String("id", "", "Target instance ID")
	updateInstanceStatus := updateInstanceCmd.String("status", "", "Instance status (poweroff, poweron, reboot)")

	// project handlers

	GetProjectCmd := flag.NewFlagSet("get project", flag.ExitOnError)
	getAllProjects := GetProjectCmd.Bool("all", false, "Get all projects")
	GetProjectById := GetProjectCmd.String("id", "", "Get project by ID")

	createProjectCmd := flag.NewFlagSet("create project", flag.ExitOnError)
	AddProjectName := createProjectCmd.String("name", "", "The Project name")
	AddProjectHost := createProjectCmd.String("host", "", "Gitlab host")
	AddProjectToken := createProjectCmd.String("token", "", "Gitlab Token")
	AddProjectGitUsername := createProjectCmd.String("git", "", "Git username")
	AddProjectNamespace := createProjectCmd.String("group", "", "Gitlab Group ID")

	DeleteInstanceCmd := flag.NewFlagSet("delete project", flag.ExitOnError)
	DeleteInstanceById := DeleteInstanceCmd.String("id", "", "Target project ID")

	// environnment handlers

	GetEnvCmd := flag.NewFlagSet("get environment", flag.ExitOnError)
	getAllEnv := GetEnvCmd.Bool("all", false, "Get all environments")
	getEnvById := GetEnvCmd.String("id", "", "Get environment by ID")

	// login handlers
	loginCmd := flag.NewFlagSet("login", flag.ExitOnError)
	loginEmail := loginCmd.String("a", "", "access key")
	loginPassword := loginCmd.String("s", "", "secret key")

	// configuration handlers
	configureCmd := flag.NewFlagSet("configure", flag.ExitOnError)
	configureRegionCmd := configureCmd.Bool("region", false, "Configure the default region")
	configureEndpointCmd := configureCmd.Bool("endpoint", false, "Configure the cloud api endpoint")
	configureProviderCmd := configureCmd.Bool("provider", false, "Configure the default provider")

	if len(os.Args) < 2 {
		fmt.Println("usage: cwc <command> [parameters]")
		fmt.Printf("To see help text, you can run:\n\n")
		// fmt.Println("cwc help")
		fmt.Printf("cwc <command> help\n\n")
		fmt.Println("cwc: error: the following arguments are required: command")
		os.Exit(1)
		return
	}

	switch os.Args[1] {

	case "get":
		switch os.Args[2] {
		case "project":
			handlers.HandleGetProject(GetProjectCmd, getAllProjects, GetProjectById)
		case "instance":
			handlers.HandleGetInstance(getInstanceCmd, getAllInstances, getInstanceById)

		case "environment":
			handlers.HandleGetEnvironment(GetEnvCmd, getAllEnv, getEnvById)
		}

	case "create":

		switch os.Args[2] {
		case "project":
			handlers.HandleAddProject(createProjectCmd, AddProjectName,AddProjectHost,AddProjectToken,AddProjectGitUsername,AddProjectNamespace)
		case "instance":
			handlers.HandleAddInstance(createInstanceCmd, addInstanceName, addInstanceProjectId, addInstanceEnvironment, addInstanceType)
		}
	case "attach":
		switch os.Args[2] {
		case "instance":
			handlers.HandleAttachInstance(attachInstanceCmd, attachInstanceProjectId, attachInstancePlaybook, attachInstanceType)
		}
	case "delete":
		switch os.Args[2] {
		case "project":
			handlers.HandleDeleteProject(DeleteInstanceCmd, DeleteInstanceById)
		case "instance":
			handlers.HandleDeleteInstance(deleteInstanceCmd, deleteInstanceById)
		}
	case "update":
		switch os.Args[2] {
		case "instance":
			handlers.HandleUpdateInstance(updateInstanceCmd, updateInstanceId, updateInstanceStatus)
		}

	case "login":
		handlers.HandleLogin(loginCmd, loginEmail, loginPassword)
	case "configure":
		handlers.HandleConfigure(configureCmd, configureRegionCmd,configureEndpointCmd,configureProviderCmd)
	case "help":
		handlers.HandleHelp(helpCmd)
	case "--version":
		handlers.HandleVersion(versionCmd, Version)
	default:
		fmt.Printf("cwc: command not found")
	}

}
