package handlers

import (
	"cwc/client"
	"flag"
	"fmt"
	"os"
)


func HandleHelp(getCmd *flag.FlagSet){
	fmt.Printf("cwc: available commands:\n\n\n")
	fmt.Printf("- create \n")
	fmt.Printf("  create a new instance\n\n")
	fmt.Printf("- get \n")
	fmt.Printf("  get one or many instances\n\n")
	fmt.Printf("- delete \n")
	fmt.Printf("  delete an existing instance\n\n")
	fmt.Printf("- update \n")
	fmt.Printf("  update a particular instance\n\n")
}

func HandleGet(getCmd *flag.FlagSet,all *bool, id *string ){

	getCmd.Parse(os.Args[2: ])
	if !*all && *id == "" {
		fmt.Println("id is required or specify --all to get all instances.")
		getCmd.PrintDefaults()
		os.Exit(1)
	}
	region := "fr-par-1"
	client := client.NewClient(region)
	if *all{

		projects ,err := client.GetAll()

		if err != nil {
			fmt.Printf("failed: %s\n",err)
			os.Exit(1)
		}

		fmt.Printf("ID\tname\tstatus\tsize\tenvironment\tgitlab url\n")
		for _,project := range *projects{
			fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\t%v\n",project.Id,project.Name,project.Status,project.Instance_type,project.Environment,project.Ip_address,project.Gitlab_url)

		}

		return
	}

	if *id != "" {
		project ,err := client.GetProject(*id)
		if err != nil {
			fmt.Printf("failed: %s\n",err)
			os.Exit(1)
		}
		fmt.Printf("ID\tname\tstatus\tsize\tenvironment\tgitlab url\n")
		fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\t%v\n",project.Id,project.Name,project.Status,project.Instance_type,project.Environment,project.Ip_address,project.Gitlab_url)

		return
	}
}

func HandleLogin(loginCmd *flag.FlagSet, email *string,password *string ){

	loginCmd.Parse(os.Args[2: ])
	if  *email == "" || *password=="" {
		fmt.Println("email and password are required to login")
		loginCmd.PrintDefaults()
		os.Exit(1)
	}
	region := "fr-par-1"
	client := client.NewClient(region)

		err := client.UserLogin(*email,*password)
		if err != nil {
			fmt.Printf("failed: %s\n",err)
			os.Exit(1)
		}
	fmt.Printf("You are successfully logged in\n")
}

func HandleDelete(deleteCmd *flag.FlagSet, id *string ){

	deleteCmd.Parse(os.Args[2: ])
	if  *id == "" {
		fmt.Println("id is required to delete your instance")
		deleteCmd.PrintDefaults()
		os.Exit(1)
	}
	region := "fr-par-1"
	client := client.NewClient(region)

		err := client.DeleteProject(*id)
		if err != nil {
			fmt.Printf("failed: %s\n",err)
			os.Exit(1)
		}
	fmt.Printf("Project %v successfully deleted\n",*id)
}
func ValidateProject(createCmd *flag.FlagSet,name *string, email *string,env *string,instance_type *string ){

	if *name == "" || *env == "" {
		createCmd.PrintDefaults()
		os.Exit(1)
	}
}
func HandleAdd(createCmd *flag.FlagSet,name *string,email *string,env *string,instance_type *string){
	createCmd.Parse(os.Args[2: ])
	ValidateProject(createCmd,name ,email ,env ,instance_type)
	region := "fr-par-1"
	client := client.NewClient(region)
	created_project,err := client.AddProject(*name, *instance_type,*env, *email)
	if err != nil {
		fmt.Printf("failed: %s\n",err)
		os.Exit(1)
	}
	fmt.Printf("ID\tname\tstatus\tsize\tenvironment\tgitlab url\n")
	fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\n",created_project.Id,created_project.Name,created_project.Status,created_project.Instance_type,created_project.Environment,created_project.Gitlab_url)


}

func HandleUpdate(updateCmd *flag.FlagSet,id *string,status *string,instance_type *string){
	updateCmd.Parse(os.Args[2: ])
	if *id == "" {
		fmt.Println("id is required")
		updateCmd.PrintDefaults()
		os.Exit(1)
	}
	if 	*status =="" && *instance_type==""{
		fmt.Println("You have to provide either the status or the instance size")
		updateCmd.PrintDefaults()
		os.Exit(1)
	}
	region := "fr-par-1"
	client := client.NewClient(region)
	err := client.UpdateProject(*id,*status,*instance_type)
	if err != nil {
		fmt.Printf("failed: %s\n",err)
		os.Exit(1)
	}
	fmt.Printf("Project %v successfully updated\n", *id)


}