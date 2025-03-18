package user

import (
	"cwc/client"
	"cwc/config"
	"cwc/utils"
)

func HandleGetAiAdapters(adapters *client.AiAdaptersResponse, pretty *bool) {
	if config.IsPrettyFormatExpected(pretty) {
		utils.PrintPrettyArray("Available adapters", adapters.Adapters)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(adapters)
	} else {
		utils.PrintArray(adapters.Adapters)
	}
}

func HandleSendPrompt(adapter *string, message *string, listId *string, pretty *bool) {
	c, err := client.NewClient()
	utils.ExitIfError(err)

	response, err := c.SendPrompt(*adapter, *message, *listId)
	utils.ExitIfError(err)

	if config.IsPrettyFormatExpected(pretty) {
		utils.PrintPretty("AI response", *response)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(response)
	} else {
		utils.PrintRow(*response)
	}
}
