package admin

import (
	"cwc/admin"
	"cwc/config"
	"cwc/utils"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

func HandleGetAIAdapters(adapters *admin.AdminAIAdaptersResponse, pretty *bool) {
	if config.IsPrettyFormatExpected(pretty) {
		displayAIAdaptersAsTable(adapters.Adapters)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(adapters)
	} else {
		utils.PrintMultiRow(admin.AdminAIAdapter{}, adapters.Adapters)
	}
}

func HandleGetAIAdapter(adapter *admin.AdminAIAdapterDetailResponse, pretty *bool) {
	if config.IsPrettyFormatExpected(pretty) {
		displayAdapter := struct {
			Id         string                  `json:"id"`
			Name       string                  `json:"name"`
			UserId     int                     `json:"user_id"`
			Url        string                  `json:"url"`
			Username   string                  `json:"username"`
			Headers    []admin.AIAdapterHeader `json:"headers"`
			Timeout    int                     `json:"timeout"`
			CheckTls   bool                    `json:"check_tls"`
			IsPublic   bool                    `json:"is_public"`
			CreatedAt  string                  `json:"created_at"`
			UpdatedAt  string                  `json:"updated_at"`
			OwnerId    int                     `json:"owner_id"`
			OwnerEmail string                  `json:"owner_email"`
		}{
			Id:         adapter.Adapter.Id,
			Name:       adapter.Adapter.Name,
			UserId:     adapter.Adapter.UserId,
			Url:        adapter.Adapter.Url,
			Username:   adapter.Adapter.Username,
			Headers:    adapter.Adapter.Headers,
			Timeout:    adapter.Adapter.Timeout,
			CheckTls:   adapter.Adapter.CheckTls,
			IsPublic:   adapter.Adapter.IsPublic,
			CreatedAt:  adapter.Adapter.CreatedAt,
			UpdatedAt:  adapter.Adapter.UpdatedAt,
			OwnerId:    adapter.Adapter.Owner.Id,
			OwnerEmail: adapter.Adapter.Owner.Email,
		}
		utils.PrintPretty("Found AI adapter", displayAdapter)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(adapter)
	} else {
		utils.PrintRow(adapter.Adapter)
	}
}

func displayAIAdaptersAsTable(adapters []admin.AdminAIAdapter) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "URL", "Owner Email", "Public", "Updated At"})

	if len(adapters) == 0 {
		table.Append([]string{"No AI adapters available", "404", "404", "404", "404", "404"})
	} else {
		for _, adapter := range adapters {
			table.Append([]string{
				adapter.Id,
				adapter.Name,
				adapter.Url,
				adapter.Owner.Email,
				fmt.Sprintf("%t", adapter.IsPublic),
				adapter.UpdatedAt,
			})
		}
	}
	table.Render()
}

func HandleCreateAIAdapter(response *admin.AdminAIAdapterResponse, pretty *bool) {
	if config.IsPrettyFormatExpected(pretty) {
		utils.PrintPretty("AI Adapter successfully created", *response)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(response)
	} else {
		fmt.Printf("AI Adapter created successfully with ID: %s\n", response.Id)
	}
}

func HandleUpdateAIAdapter(response *admin.AdminAIAdapterResponse, pretty *bool) {
	if config.IsPrettyFormatExpected(pretty) {
		utils.PrintPretty("AI Adapter successfully updated", *response)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(response)
	} else {
		fmt.Println("AI Adapter updated successfully")
	}
}

func HandleDeleteAIAdapter(response *admin.AdminAIAdapterResponse, pretty *bool) {
	if config.IsPrettyFormatExpected(pretty) {
		utils.PrintPretty("AI Adapter successfully deleted", *response)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(response)
	} else {
		fmt.Println("AI Adapter deleted successfully")
	}
}
