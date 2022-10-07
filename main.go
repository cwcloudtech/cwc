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
	addInstanceZone := createInstanceCmd.String("zone", "", "instance zone")
	attachInstanceCmd := flag.NewFlagSet("attach instance", flag.ExitOnError)
	attachInstancePlaybook := attachInstanceCmd.String("name", "", "The playbook name you want to run")
	attachInstanceProjectId := attachInstanceCmd.Int("project_id", 0, "The project id that you want to attach with the instance")
	attachInstanceType := attachInstanceCmd.String("instance_type", "", "The instance size (DEV1-S, DEV1-M, DEV1-L, DEV1-XL)")

	deleteInstanceCmd := flag.NewFlagSet("delete instance", flag.ExitOnError)
	deleteInstanceById := deleteInstanceCmd.String("id", "", "Target instance ID")

	updateInstanceCmd := flag.NewFlagSet("update instance", flag.ExitOnError)
	updateInstanceId := updateInstanceCmd.String("id", "", "Target instance ID")
	updateInstanceStatus := updateInstanceCmd.String("status", "", "Instance status (poweroff, poweron, reboot)")

	// bucket handlers
	getBucketCmd := flag.NewFlagSet("get bucket", flag.ExitOnError)
	getAllBuckets := getBucketCmd.Bool("all", false, "Get all buckets")
	getBucketById := getBucketCmd.String("id", "", "Get bucket by ID")

	deleteBucketCmd := flag.NewFlagSet("delete bucket", flag.ExitOnError)
	deleteBucketById := deleteBucketCmd.String("id", "", "Target bucket ID")

	updateBucketCmd := flag.NewFlagSet("update bucket", flag.ExitOnError)
	updateBucketById := updateBucketCmd.String("id", "", "Target bucket ID")

	// registry handlers
	getRegistryCmd := flag.NewFlagSet("get registry", flag.ExitOnError)
	getAllRegistries := getRegistryCmd.Bool("all", false, "Get all registries")
	getRegistryById := getRegistryCmd.String("id", "", "Get registry by ID")

	deleteRegistryCmd := flag.NewFlagSet("delete registry", flag.ExitOnError)
	deleteRegistryById := deleteRegistryCmd.String("id", "", "Target registry ID")

	updateRegistryCmd := flag.NewFlagSet("update registry", flag.ExitOnError)
	updateRegistryById := updateRegistryCmd.String("id", "", "Target registry ID")

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
		case "bucket":
			handlers.HandleGetBucket(getBucketCmd, getAllBuckets, getBucketById)
		case "registry":
			handlers.HandleGetRegistry(getRegistryCmd, getAllRegistries, getRegistryById)
		case "environment":
			handlers.HandleGetEnvironment(GetEnvCmd, getAllEnv, getEnvById)
		}

	case "create":

		switch os.Args[2] {
		case "project":
			handlers.HandleAddProject(createProjectCmd, AddProjectName, AddProjectHost, AddProjectToken, AddProjectGitUsername, AddProjectNamespace)
		case "instance":
			handlers.HandleAddInstance(createInstanceCmd, addInstanceName, addInstanceProjectId, addInstanceEnvironment, addInstanceType, addInstanceZone)
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
		case "registry":
			handlers.HandleDeleteRegistry(deleteRegistryCmd, deleteRegistryById)
		case "bucket":
			handlers.HandleDeleteBucket(deleteBucketCmd, deleteBucketById)
		}
	case "update":
		switch os.Args[2] {
		case "instance":
			handlers.HandleUpdateInstance(updateInstanceCmd, updateInstanceId, updateInstanceStatus)
		case "bucket":
			handlers.HandleUpdateBucket(updateBucketCmd, updateBucketById)
		case "registry":
			handlers.HandleUpdateRegistry(updateRegistryCmd, updateRegistryById)
		}

	case "login":
		handlers.HandleLogin(loginCmd, loginEmail, loginPassword)
	case "configure":
		handlers.HandleConfigure(configureCmd, configureRegionCmd, configureEndpointCmd, configureProviderCmd)
	case "help":
		handlers.HandleHelp(helpCmd)
	case "--version":
		handlers.HandleVersion(versionCmd, Version)
	default:
		fmt.Printf("cwc: command not found")
	}

}
