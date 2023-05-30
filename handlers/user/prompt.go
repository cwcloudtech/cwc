package user

import (
	"cwc/client"
	"cwc/utils"
	"fmt"
	"os"
)

func HandleGetModels() {
	c, err := client.NewClient()
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	email, err := c.GetModels()
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	if client.GetDefaultFormat() == "json" {
		utils.PrintJson(email)
	} else {
		utils.PrintRow(*email)
	}

}

func HandleSendPrompt(model *string, message *string) {
	c, err := client.NewClient()
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	email, err := c.SendPrompt(*model, *message)
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	if client.GetDefaultFormat() == "json" {
		utils.PrintJson(email)
	} else {
		utils.PrintRow(*email)
	}

}
