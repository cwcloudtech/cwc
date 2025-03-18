package user

import (
	"cwc/client"
	"cwc/config"
	"cwc/utils"
	"os"

	"github.com/olekukonko/tablewriter"
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

func HandleGetPromptHistory(startIndex *int, maxResults *int, pretty *bool) {
	c, err := client.NewClient()
	utils.ExitIfError(err)

	startIndexValue := 0
	if startIndex != nil {
		startIndexValue = *startIndex
	}

	maxResultsValue := 7
	if maxResults != nil {
		maxResultsValue = *maxResults
	}

	response, err := c.GetPromptHistory(startIndexValue, maxResultsValue)
	utils.ExitIfError(err)

	if config.IsPrettyFormatExpected(pretty) {
		displayPromptListsAsTable(response.PromptLists)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(response)
	} else {
		utils.PrintMultiRow(client.PromptList{}, response.PromptLists)
	}
}

func displayPromptListsAsTable(promptLists []client.PromptList) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Title", "Updated At", "Created At"})

	if len(promptLists) == 0 {
		table.Append([]string{"No prompt lists available", "404", "404", "404"})
	} else {
		for _, list := range promptLists {
			table.Append([]string{
				list.Id,
				list.Title,
				list.UpdatedAt,
				list.CreatedAt,
			})
		}
	}
	table.Render()
}
