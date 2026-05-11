package admin

import (
	"cwc/admin"
	"cwc/config"
	"cwc/utils"
	"encoding/json"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

func HandleListStorageKVs(search string, userId string, startIndex int, maxResults int, pretty *bool) error {
	c, err := admin.NewClient()
	if err != nil {
		return fmt.Errorf("error creating admin client: %v", err)
	}

	var kvs *admin.StorageKVListResponse

	if utils.IsNotBlank(userId) {
		kvs, err = c.GetUserStorageKVs(userId, search, startIndex, maxResults)
		if err != nil {
			return fmt.Errorf("error listing user's storage key-values: %v", err)
		}
	} else {
		kvs, err = c.GetAllStorageKVs(search, startIndex, maxResults, "")
		if err != nil {
			return fmt.Errorf("error listing storage key-values: %v", err)
		}
	}

	if config.IsPrettyFormatExpected(pretty) {
		displayAdminStorageKVsAsTable(kvs.Items)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(kvs.Items)
	} else {
		displayAdminStorageKVsAsTable(kvs.Items)
	}

	return nil
}

func displayAdminStorageKVsAsTable(kvs []admin.StorageKV) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"User ID", "Key", "Payload", "Created At", "Updated At", "Source"})

	if len(kvs) == 0 {
		table.Append([]string{"No storage key-value entries available", "404", "404", "404", "404", "404"})
	} else {
		for _, kv := range kvs {
			payloadJSON, _ := json.Marshal(kv.Payload)
			var payloadFormatted string
			if len(payloadJSON) > 50 {
				payloadFormatted = string(payloadJSON[:47]) + "..."
			} else {
				payloadFormatted = string(payloadJSON)
			}
			userID := fmt.Sprintf("%d", kv.UserId)
			table.Append([]string{
				userID,
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

func HandleDeleteUserStorageKV(userId string, key string, pretty *bool) error {
	c, err := admin.NewClient()
	if err != nil {
		return fmt.Errorf("error creating admin client: %v", err)
	}

	_, err = c.DeleteUserStorageKV(userId, key)
	if err != nil {
		return fmt.Errorf("error deleting user's storage key-value: %v", err)
	}

	if config.GetDefaultFormat() == "json" {
		simpleResponse := struct {
			Message string `json:"message"`
			Key     string `json:"key"`
			UserId  string `json:"user_id"`
		}{
			Message: "Storage KV deleted",
			Key:     key,
			UserId:  userId,
		}
		utils.PrintJson(simpleResponse)
	} else {
		fmt.Printf("Storage KV deleted for user %s: %s\n", userId, key)
	}

	return nil
}
