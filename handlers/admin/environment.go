package admin

import (
	"cwc/admin"
	"cwc/utils"
	"fmt"
	"os"
	"github.com/olekukonko/tablewriter"
	"os/exec"
    "io/ioutil"
)

func HandleAddEnvironment(name *string, path *string, roles *string, is_private *bool, description *string, subdomains *string, logo_url *string) {
	added_environment := &admin.Environment{
		Name:        *name,
		Path:        *path,
		Roles:       *roles,
		IsPrivate:   *is_private,
		Description: *description,
		SubDomains:  *subdomains,
		LogUrl:     *logo_url,
		EnvironmentTemplate: "",
		DocTemplate: "",
	}
	// prompt for environment template
	fmt.Print("Do you want to add environment template? [Y/N]: ")
	var add_env_template string
	fmt.Scanln(&add_env_template)

	if add_env_template == "Y" || add_env_template == "y" {
		var editorCommand string
		editorCommand = os.Getenv("EDITOR")
		if editorCommand == "" {
			editorCommand = "vi"
		}

		// Create a temporary file with a specific name and path
		tempFileName := "temp-code-editor.txt"
		_, err := os.Create(tempFileName)
		if err != nil {
			fmt.Printf("Error creating temporary file: %s\n", err)
			os.Exit(1)
		}
		defer os.Remove(tempFileName)
	
		// Prompt the user to write code in the editor
		fmt.Printf("Please write your code in the text editor that opens. Save and close the editor when done.\n")
	
		cmd := exec.Command(editorCommand, tempFileName)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	
		if err := cmd.Run(); err != nil {
			fmt.Printf("Error opening the text editor: %s\n", err)
			os.Exit(1)
		}
	
		// Read the code from the temporary file
		codeBytes, err := ioutil.ReadFile(tempFileName)
		if err != nil {
			fmt.Printf("Error reading code from the text editor: %s\n", err)
			os.Exit(1)
		}

		added_environment.EnvironmentTemplate = string(codeBytes)
	}

	// prompt for doc template
	fmt.Print("Do you want to add doc template? [Y/N]: ")
	var add_doc_template string
	fmt.Scanln(&add_doc_template)

	if add_doc_template == "Y" || add_doc_template == "y" {
		var editorCommand string
		editorCommand = os.Getenv("EDITOR")
		if editorCommand == "" {
			editorCommand = "vi"
		}

		// Create a temporary file with a specific name and path
		tempFileName := "temp-code-editor.txt"
		_, err := os.Create(tempFileName)
		if err != nil {
			fmt.Printf("Error creating temporary file: %s\n", err)
			os.Exit(1)
		}
		defer os.Remove(tempFileName)

		// Prompt the user to write code in the editor
		fmt.Printf("Please write your code in the text editor that opens. Save and close the editor when done.\n")

		cmd := exec.Command(editorCommand, tempFileName)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			fmt.Printf("Error opening the text editor: %s\n", err)
			os.Exit(1)
		}

		// Read the code from the temporary file
		codeBytes, err := ioutil.ReadFile(tempFileName)
		if err != nil {
			fmt.Printf("Error reading code from the text editor: %s\n", err)
			os.Exit(1)
		}

		added_environment.DocTemplate = string(codeBytes)
	}

	client, err := admin.NewClient()
	created_env, err := client.AdminAddEnvironment(*added_environment)
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	if admin.GetDefaultFormat() == "json" {
		utils.PrintJson(created_env)
	} else {
		fmt.Printf("Environment %s successfully created\n", created_env.Name)
		fmt.Printf("➤ ID: %d\n", created_env.Id)
		fmt.Printf("➤ Name: %s\n", created_env.Name)
		fmt.Printf("➤ Path: %s\n", created_env.Path)
		fmt.Printf("➤ Description: %s\n", created_env.Description)
		fmt.Printf("➤ Subdomains: %s\n", created_env.SubDomains)
		fmt.Printf("➤ Is Private: %t\n", created_env.IsPrivate)
		fmt.Printf("➤ Roles: %s\n", created_env.Roles)
		fmt.Printf(("➤ Environment Logo URL: %s\n"), created_env.LogUrl)
	}

}

func HandleDeleteEnvironment(id *string) {

	client, err := admin.NewClient()
	err = client.AdminDeleteEnvironment(*id)
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Environment %v successfully deleted\n", *id)
}

func HandleGetEnvironments(pretty *bool) {

	client, err := admin.NewClient()
	environments, err := client.GetAllEnvironments()

	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	if admin.GetDefaultFormat() == "json" {
		utils.PrintJson(environments)
	} else {
		if *pretty {
			displayEnvironmentsAsTable(*environments)
		} else {
			utils.PrintMultiRow(admin.Environment{}, *environments)
		}
	}
}

func HandleGetEnvironment(id *string, pretty *bool) {
	client, err := admin.NewClient()

	environment, err := client.GetEnvironment(*id)
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	if admin.GetDefaultFormat() == "json" {
		utils.PrintJson(environment)
	} else {
		if *pretty {
			fmt.Printf("➤ ID: %d\n", environment.Id)
			fmt.Printf("➤ Name: %s\n", environment.Name)
			fmt.Printf("➤ Path: %s\n", environment.Path)
			fmt.Printf("➤ Description: %s\n", environment.Description)
			fmt.Printf("➤ Subdomains: %s\n", environment.SubDomains)
			fmt.Printf("➤ Is Private: %t\n", environment.IsPrivate)
			fmt.Printf("➤ Roles: %s\n", environment.Roles)
		} else {
			utils.PrintRow(environment)
		}
	}

}


func displayEnvironmentsAsTable(environments []admin.Environment) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "Path", "Description", "Subdomains", "Is Private" })

	if len(environments) == 0 {
		fmt.Println("No environments found")
	} else {
		for _, environment := range environments {
			table.Append([]string{
				fmt.Sprintf("%d", environment.Id), 
				environment.Name,
				environment.Path,
				environment.Description,
				environment.SubDomains,
				fmt.Sprintf("%t", environment.IsPrivate),
			})
		}
	}

	table.Render()
}