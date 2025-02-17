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

func HandleSendPrompt(adapter *string, message *string) {
	c, err := client.NewClient()
	utils.ExitIfError(err)

	response, err := c.SendPrompt(*adapter, *message)
	utils.ExitIfError(err)

	if config.GetDefaultFormat() == "json" {
		utils.PrintJson(response)
	} else {
		utils.PrintRow(*response)
	}
}
