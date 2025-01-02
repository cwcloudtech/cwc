package user

import (
	"cwc/client"
	"cwc/config"
	"cwc/utils"
)

func HandleGetModels(models *client.ModelsResponse, pretty *bool) {
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

	response, err := c.SendPrompt(*model, *message)
	utils.ExitIfError(err)

	if config.GetDefaultFormat() == "json" {
		utils.PrintJson(response)
	} else {
		utils.PrintRow(*response)
	}
}

func HandleLoadModel(model *string) {
	c, err := client.NewClient()
	utils.ExitIfError(err)

	response, err := c.LoadModel(*model)
	utils.ExitIfError(err)

	if config.GetDefaultFormat() == "json" {
		utils.PrintJson(response)
	} else {
		utils.PrintRow(*response)
	}
}
