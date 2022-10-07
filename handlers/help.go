package handlers

import (
	"flag"
	"fmt"
	"os"
)

func HandleHelp(helpCmd *flag.FlagSet) {

	helpCmd.Parse(os.Args[2:])
	fmt.Printf("cwc: available commands:\n\n\n")
	fmt.Printf("- create instance \n")
	fmt.Printf("  create a new instance\n\n")
	fmt.Printf("- get instance \n")
	fmt.Printf("  get one or many instances\n\n")
	fmt.Printf("- delete instance \n")
	fmt.Printf("  delete an existing instance\n\n")
	fmt.Printf("- update instance \n")
	fmt.Printf("  update a particular instance state\n\n")

	fmt.Printf("- create project \n")
	fmt.Printf("  create a new project\n\n")
	fmt.Printf("- get project \n")
	fmt.Printf("  get one or many projects\n\n")
	fmt.Printf("- delete project \n")
	fmt.Printf("  delete an existing project\n\n")

	fmt.Printf("- get environement \n")
	fmt.Printf("  get one or many environments\n\n")

	fmt.Printf("- login \n")
	fmt.Printf("  login to your account\n\n")
	fmt.Printf("- configure \n")
	fmt.Printf("  configure your default settings like region\n\n")

}
