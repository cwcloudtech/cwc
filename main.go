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

	getCmd := flag.NewFlagSet("get", flag.ExitOnError)
	getAll := getCmd.Bool("all", false, "Get all instances")
	getById := getCmd.String("id", "", "Get instance by ID")

	createCmd := flag.NewFlagSet("create", flag.ExitOnError)
	addProjectName := createCmd.String("name", "", "The Project name")
	addEmail := createCmd.String("email", "", "The email address associeted to the project")
	addEnvironment := createCmd.String("env", "", "The environment of the project (code, wpaas)")
	addInstanceType := createCmd.String("instance_type", "", "The instance size (DEV1-S, DEV1-M, DEV1-L, DEV1-XL)")

	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	deleteById := deleteCmd.String("id", "", "Target instance ID")

	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
	updateId := updateCmd.String("id", "", "Target instance ID")
	updateStatus := updateCmd.String("status", "", "Instance status (poweroff, poweron, reboot)")

	loginCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	loginEmail := loginCmd.String("u", "", "Account email")
	loginPassword := loginCmd.String("p", "", "Account password")

	configureCmd := flag.NewFlagSet("configure", flag.ExitOnError)
	configureRegionCmd := configureCmd.Bool("region", false, "Configure the default region")

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
		handlers.HandleGet(getCmd, getAll, getById)
	case "create":
		handlers.HandleAdd(createCmd, addProjectName, addEmail, addEnvironment, addInstanceType)
	case "delete":
		handlers.HandleDelete(deleteCmd, deleteById)
	case "update":
		handlers.HandleUpdate(updateCmd, updateId, updateStatus)

	case "login":
		handlers.HandleLogin(loginCmd, loginEmail, loginPassword)
	case "configure":
		handlers.HandleConfigure(configureCmd, configureRegionCmd)
	case "help":
		handlers.HandleHelp(helpCmd)
	case "--version":
		handlers.HandleVersion(versionCmd, Version)
	default:
		fmt.Printf("cwc: command not found")
	}

}
