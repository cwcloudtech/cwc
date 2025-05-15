package user

import (
	"cwc/client"
	"cwc/config"
	"cwc/utils"
	"encoding/json"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

func PrepareCreateStorageKV(storageKV *client.StorageKVCreateRequest) (*client.StorageKVResponse, error) {
	if storageKV.Key == "" {
		return nil, fmt.Errorf("key cannot be empty")
	}

	c, err := client.NewClient()
	if err != nil {
		return nil, fmt.Errorf("error creating client: %v", err)
	}

	created, err := c.CreateStorageKV(*storageKV)
	if err != nil {
		return nil, fmt.Errorf("error creating storage key-value: %v", err)
	}

	return created, nil
}

func HandleCreateStorageKV(created *client.StorageKVResponse, pretty *bool) error {
	if created == nil {
		return fmt.Errorf("error: created storage KV is nil")
	}

	simplified := client.SimplifiedStorageKVResponse{
		Key:       created.Key,
		Message:   created.Message,
		Payload:   created.Payload,
		CreatedAt: created.CreatedAt,
		UpdatedAt: created.UpdatedAt,
		Source:    created.Source,
	}

	if config.IsPrettyFormatExpected(pretty) {
		utils.PrintPretty("Storage KV created", simplified)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(simplified)
	} else {
		payloadJSON, _ := json.Marshal(created.Payload)
		display := client.StorageKVDisplay{
			Key:       created.Key,
			Payload:   string(payloadJSON),
			CreatedAt: created.CreatedAt,
			UpdatedAt: created.UpdatedAt,
			Source:    created.Source,
		}
		utils.PrintRow(display)
	}

	return nil
}

func HandleListStorageKVs(search string, startIndex int, maxResults int, pretty *bool) error {
	c, err := client.NewClient()
	if err != nil {
		return fmt.Errorf("error creating client: %v", err)
	}

	kvs, err := c.ListStorageKVs(search, startIndex, maxResults)
	if err != nil {
		return fmt.Errorf("error listing storage key-values: %v", err)
	}

	if config.IsPrettyFormatExpected(pretty) {
		displayStorageKVsAsTable(kvs.Items)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(kvs.Items)
	} else {
		var kvDisplays []client.StorageKVDisplay
		for _, kv := range kvs.Items {
			payloadStr, _ := json.Marshal(kv.Payload)
			kvDisplays = append(kvDisplays, client.StorageKVDisplay{
				Key:       kv.Key,
				Payload:   string(payloadStr),
				CreatedAt: kv.CreatedAt,
				UpdatedAt: kv.UpdatedAt,
				Source:    kv.Source,
			})
		}
		utils.PrintMultiRow(client.StorageKVDisplay{}, kvDisplays)
	}

	return nil
}

func displayStorageKVsAsTable(kvs []client.StorageKV) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Key", "Payload", "Created At", "Updated At", "Source"})

	if len(kvs) == 0 {
		table.Append([]string{"No storage key-value entries available", "404", "404", "404", "404"})
	} else {
		for _, kv := range kvs {
			payloadJSON, _ := json.Marshal(kv.Payload)
			var payloadFormatted string
			if len(payloadJSON) > 50 {
				payloadFormatted = string(payloadJSON[:47]) + "..."
			} else {
				payloadFormatted = string(payloadJSON)
			}
			table.Append([]string{
				kv.Key,
				payloadFormatted,
				kv.CreatedAt,
				kv.UpdatedAt,
				kv.Source,
			})
		}
	}
	table.Render()
}

func PrepareUpdateStorageKV(key string, updateRequest *client.StorageKVUpdateRequest) (*client.StorageKVResponse, error) {
	if key == "" {
		return nil, fmt.Errorf("key cannot be empty")
	}

	c, err := client.NewClient()
	if err != nil {
		return nil, fmt.Errorf("error creating client: %v", err)
	}

	updated, err := c.UpdateStorageKV(key, *updateRequest)
	if err != nil {
		return nil, fmt.Errorf("error updating storage key-value: %v", err)
	}

	return updated, nil
}

func HandleUpdateStorageKV(updated *client.StorageKVResponse, pretty *bool) error {
	if updated == nil {
		return fmt.Errorf("error: updated storage KV is nil")
	}

	simplified := client.SimplifiedStorageKVResponse{
		Key:       updated.Key,
		Message:   updated.Message,
		Payload:   updated.Payload,
		CreatedAt: updated.CreatedAt,
		UpdatedAt: updated.UpdatedAt,
		Source:    updated.Source,
	}

	if config.IsPrettyFormatExpected(pretty) {
		utils.PrintPretty("Storage KV updated", simplified)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(simplified)
	} else {
		payloadJSON, _ := json.Marshal(updated.Payload)
		display := client.StorageKVDisplay{
			Key:       updated.Key,
			Payload:   string(payloadJSON),
			CreatedAt: updated.CreatedAt,
			UpdatedAt: updated.UpdatedAt,
			Source:    updated.Source,
		}
		utils.PrintRow(display)
	}

	return nil
}

func HandleDeleteStorageKV(key string, pretty *bool) error {
	c, err := client.NewClient()
	if err != nil {
		return fmt.Errorf("error creating client: %v", err)
	}

	_, err = c.DeleteStorageKV(key)
	if err != nil {
		return fmt.Errorf("error deleting storage key-value: %v", err)
	}

	if config.GetDefaultFormat() == "json" {
		simpleResponse := struct {
			Message string `json:"message"`
			Key     string `json:"key"`
		}{
			Message: "Storage KV deleted",
			Key:     key,
		}
		utils.PrintJson(simpleResponse)
	} else {
		fmt.Printf("Storage KV deleted: %s\n", key)
	}

	return nil
}
