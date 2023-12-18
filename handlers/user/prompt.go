package user

import (
	"cwc/client"
	"cwc/config"
	"cwc/utils"
)

func HandleGetModels(pretty *bool) {
	c, err := client.NewClient()
	utils.ExitIfError(err)

	models, err := c.GetModels()
	utils.ExitIfError(err)

	if config.IsPrettyFormatExpected(pretty) {
		utils.PrintPrettyArray("Available models", models.Models)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(models)
	} else {
		utils.PrintArray(models.Models)
	}
}

func HandleSendPrompt(model *string, message *string) {
	c, err := client.NewClient()
	utils.ExitIfError(err)

	email, err := c.SendPrompt(*model, *message)
	utils.ExitIfError(err)

	if config.GetDefaultFormat() == "json" {
		utils.PrintJson(email)
	} else {
		utils.PrintRow(*email)
	}
}
