package main

import (
	// "fmt"
	"cwc/handlers"
	"flag"
	"fmt"
	"os"
)



func main(){
	getCmd := flag.NewFlagSet("get",flag.ExitOnError)
	getAll:= getCmd.Bool("all",false,"Get all instances")
	getById:= getCmd.String("id", "" ,"Get instance by ID")


	addCmd := flag.NewFlagSet("add",flag.ExitOnError)
	addProjectName:= addCmd.String("name", "" ,"The Project name")
	addEmail:= addCmd.String("email", "" ,"The email address associeted to the project")
	addEnvironment:= addCmd.String("env", "" ,"The environment of the project")
	addInstanceType:= addCmd.String("instance_type", "" ,"The instance size")


	if len(os.Args)<2 {
		fmt.Println("Expected get or add sub command")
		os.Exit(1)
	}

	switch os.Args[1]{

	case "get":
		handlers.HandleGet(getCmd,getAll,getById)
	case "add":
		handlers.HandleAdd(addCmd,addProjectName,addEmail,addEnvironment,addInstanceType)
	default: 
	}


}

