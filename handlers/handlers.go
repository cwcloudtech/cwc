package handlers

import (
	"cwc/client"
	"flag"
	"fmt"
	"os"
)

func HandleGet(getCmd *flag.FlagSet,all *bool, id *string ){

	getCmd.Parse(os.Args[2: ])
	if *all == false && *id == "" {
		fmt.Println("id is required or specify --all to get all instances.")
		getCmd.PrintDefaults()
		os.Exit(1)
	}
	region := "fr-par-1"
	token := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpZCI6MiwiZW1haWwiOiJhbWlyZ2hlZGlyYTA2QGdtYWlsLmNvbSIsInRpbWUiOiIwNS8wOC8yMDIyLCAxOTo0MDo0MCJ9.KVSLQGmlW2yyBVIfGMuMtr_RkNjqZRjlsEDSlPILj6k"
	client := client.NewClient(region, token)
	if *all{

		projects ,err := client.GetAll()

		if err != nil {
			fmt.Printf("an error occured")
			os.Exit(1)
		}

		fmt.Printf("ID\tproject name\tsize\tenvironment\tpublic ip\tgitlab url\n")
		for _,project := range *projects{
			fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v",project.Id,project.Name,project.Instance_type,project.Environment,project.Ip_address,project.Gitlab_url)

		}

		return
	}

	if *id != "" {
		project ,err := client.GetProject(*id)
		if err != nil {
			os.Exit(1)
		}
		fmt.Printf("ID\tproject name\tsize\tenvironment\tpublic ip\tgitlab url\n")
		fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v",project.Id,project.Name,project.Instance_type,project.Environment,project.Ip_address,project.Gitlab_url)

		return
	}
}

func ValidateProject(addCmd *flag.FlagSet,name *string, email *string,env *string,instance_type *string ){

	fmt.Println(*env)
	fmt.Println(*name)

	if *name == "" || *env == "" {
		addCmd.PrintDefaults()
		os.Exit(1)
	}

}

func HandleAdd(addCmd *flag.FlagSet,name *string,email *string,env *string,instance_type *string){
	addCmd.Parse(os.Args[2: ])
	ValidateProject(addCmd,name ,email ,env ,instance_type)
	region := "fr-par-1"
	token := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpZCI6MiwiZW1haWwiOiJhbWlyZ2hlZGlyYTA2QGdtYWlsLmNvbSIsInRpbWUiOiIwNS8wOC8yMDIyLCAxOTo0MDo0MCJ9.KVSLQGmlW2yyBVIfGMuMtr_RkNjqZRjlsEDSlPILj6k"
	client := client.NewClient(region, token)
	created_project,err := client.AddProject(*name, *instance_type,*env, *email)
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("ID\tproject name\tsize\tenvironment\tgitlab url\n")
	fmt.Printf("%v\t%v\t%v\t%v\t%v",created_project.Id,created_project.Name,created_project.Instance_type,created_project.Environment,created_project.Gitlab_url)


}