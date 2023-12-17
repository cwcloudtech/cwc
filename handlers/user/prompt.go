package user

import (
	"cwc/client"
	"cwc/utils"
)

func HandleGetModels() {
	c, err := client.NewClient()
	utils.ExitIfError(err)

	models, err := c.GetModels()
	utils.ExitIfError(err)

	if client.GetDefaultFormat() == "json" {
		utils.PrintJson(models)
	} else {
		utils.PrintRow(*models)
	}
}

func HandleSendPrompt(model *string, message *string) {
	c, err := client.NewClient()
	utils.ExitIfError(err)

	email, err := c.SendPrompt(*model, *message)
	utils.ExitIfError(err)

	if client.GetDefaultFormat() == "json" {
		utils.PrintJson(email)
	} else {
		utils.PrintRow(*email)
	}
}
