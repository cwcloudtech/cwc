package handlers

import (
	"flag"
	"fmt"
	"os"
)

func HandleCreateHelp(helpCmd *flag.FlagSet) {
	helpCmd.Parse(os.Args[2:])
	fmt.Printf("Usage: cwc <commands>\n\n")
	fmt.Printf("List of available create commands:\n")
	fmt.Printf("- create instance \n")
	fmt.Printf("  create a new instance\n")
	fmt.Printf("- create project\n")
	fmt.Printf("  create a new project\n")
	fmt.Printf("\nFor getting more informations, see this tutorial: https://doc.cloud.comwork.io/docs/tutorials/api/cli\n\n")
}

func HandleDeleteHelp(helpCmd *flag.FlagSet) {
	helpCmd.Parse(os.Args[2:])
	fmt.Printf("Usage: cwc <commands>\n\n")
	fmt.Printf("List of available delete commands:\n")
	fmt.Printf("- delete instance \n")
	fmt.Printf("  delete an existing instance\n")
	fmt.Printf("- delete project \n")
	fmt.Printf("  delete an existing project\n")
	fmt.Printf("- delete bucket \n")
	fmt.Printf("  delete an existing bucket\n")
	fmt.Printf("- delete registry \n")
	fmt.Printf("  delete an existing registry\n")
	fmt.Printf("\nFor getting more informations, see this tutorial: https://doc.cloud.comwork.io/docs/tutorials/api/cli\n\n")
}

func HandleUpdateHelp(helpCmd *flag.FlagSet) {
	helpCmd.Parse(os.Args[2:])
	fmt.Printf("Usage: cwc <commands>\n\n")
	fmt.Printf("List of available update commands:\n")
	fmt.Printf("- update instance \n")
	fmt.Printf("  update a particular instance state\n")
	fmt.Printf("\nFor getting more informations, see this tutorial: https://doc.cloud.comwork.io/docs/tutorials/api/cli\n\n")
}

func HandleGetHelp(helpCmd *flag.FlagSet) {
	helpCmd.Parse(os.Args[2:])
	fmt.Printf("Usage: cwc <commands>\n\n")
	fmt.Printf("List of available get commands:\n")
	fmt.Printf("- get instance \n")
	fmt.Printf("  get one or many instances\n")
	fmt.Printf("- get project\n")
	fmt.Printf("  get one or many projects\n")
	fmt.Printf("- get registry\n")
	fmt.Printf("  get one or many registries\n")
	fmt.Printf("- get bucket\n")
	fmt.Printf("  get one or many buckets\n")
	fmt.Printf("- get environement\n")
	fmt.Printf("  get one or many environments\n")
	fmt.Printf("\nFor getting more informations, see this tutorial: https://doc.cloud.comwork.io/docs/tutorials/api/cli\n\n")
}

func HandleHelp(helpCmd *flag.FlagSet) {
	helpCmd.Parse(os.Args[2:])
	fmt.Printf("Usage: cwc <commands>\n\n")
	fmt.Printf("List of available commands:\n")
	fmt.Printf("- help (or -h|--help)\n")
	fmt.Printf("  getting help details\n")
	fmt.Printf("- version (or -v|--version)\n")
	fmt.Printf("  getting the CLI version\n")
	fmt.Printf("- login\n")
	fmt.Printf("  login to your account\n")
	fmt.Printf("- configure\n")
	fmt.Printf("  configure your default settings like region\n")
	fmt.Printf("- create instance \n")
	fmt.Printf("  create a new instance\n")
	fmt.Printf("- get instance \n")
	fmt.Printf("  get one or many instances\n")
	fmt.Printf("- delete instance \n")
	fmt.Printf("  delete an existing instance\n")
	fmt.Printf("- update instance \n")
	fmt.Printf("  update a particular instance state\n")
	fmt.Printf("- create project\n")
	fmt.Printf("  create a new project\n")
	fmt.Printf("- get project\n")
	fmt.Printf("  get one or many projects\n")
	fmt.Printf("- delete project \n")
	fmt.Printf("  delete an existing project\n")
	fmt.Printf("- get bucket\n")
	fmt.Printf("  get one or many buckets\n")
	fmt.Printf("- delete bucket \n")
	fmt.Printf("  delete an existing bucket\n")
	fmt.Printf("- get registry\n")
	fmt.Printf("  get one or many registries\n")
	fmt.Printf("- delete registry \n")
	fmt.Printf("  delete an existing registry\n")
	fmt.Printf("- get environement\n")
	fmt.Printf("  get one or many environments\n")
	fmt.Printf("\nFor getting more informations, see this tutorial: https://doc.cloud.comwork.io/docs/tutorials/api/cli\n\n")
}
