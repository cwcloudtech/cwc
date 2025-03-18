package user

import (
	"cwc/client"
	"cwc/config"
	"cwc/utils"
	"fmt"
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

func HandleGetPromptDetails(listId *string, pretty *bool) {
	c, err := client.NewClient()
	utils.ExitIfError(err)

	response, err := c.GetPromptDetails(*listId)
	utils.ExitIfError(err)

	if config.IsPrettyFormatExpected(pretty) {
		displayPromptDetails(response)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(response)
	} else {
		displayPromptDetailsSummary(response)
	}
}

func displayPromptDetails(response *client.PromptDetailsResponse) {
	fmt.Printf("Prompt List: %s\n", response.PromptList.Title)
	fmt.Printf("ID: %s\n", response.PromptList.Id)
	fmt.Printf("Created: %s\n", response.PromptList.CreatedAt)
	fmt.Printf("Updated: %s\n", response.PromptList.UpdatedAt)
	
	fmt.Printf("Prompts (%d):\n", len(response.Prompts))
	
	for i, prompt := range response.Prompts {
		fmt.Printf("\n--- Prompt %d ---\n", i+1)
		fmt.Printf("ID: %s\n", prompt.Id)
		fmt.Printf("Adapter: %s\n", prompt.Adapter)
		fmt.Printf("Created: %s\n", prompt.CreatedAt)
		fmt.Printf("Message: %s\n", prompt.Prompt.Message)
		fmt.Printf("Response: %s\n", prompt.Answer.Response)
		fmt.Printf("Usage: %d tokens (Prompt: %d, Completion: %d)\n", 
			prompt.Answer.Usage.Total,
			prompt.Answer.Usage.Prompt,
			prompt.Answer.Usage.Completion)
	}
}

func displayPromptDetailsSummary(response *client.PromptDetailsResponse) {
    for _, prompt := range response.Prompts {
        fmt.Printf("%s\t%s\t%s\n", prompt.Id, prompt.Prompt.Message, prompt.Answer.Response)
    }
}
