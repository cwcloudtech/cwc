package update

import (
	"cwc/client"
	"cwc/handlers/user"
	"cwc/utils"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	key                  string
	ttl                  int
	pretty               bool = false
	payload              string
	updateRequestPayload interface{}
)

var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a key-value entry in your storage",
	Long:  "This command updates an existing key-value entry in your cloud storage.",
	Run: func(cmd *cobra.Command, args []string) {
		updateRequest := client.StorageKVUpdateRequest{
			Payload: updateRequestPayload,
		}

		if ttl > 0 {
			updateRequest.TTL = &ttl
		}

		updated, err := user.PrepareUpdateStorageKV(key, &updateRequest)
		utils.ExitIfError(err)

		err = user.HandleUpdateStorageKV(updated, &pretty)
		utils.ExitIfError(err)
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {
		parsedPayload, err := utils.ParseJSONPayload(payload)
		if err != nil {
			return err
		}

		updateRequestPayload = parsedPayload
		return nil
	},
}

func init() {
	UpdateCmd.Flags().StringVarP(&key, "key", "k", "", "Key of the storage entry to update")
	UpdateCmd.Flags().StringVarP(&payload, "payload", "p", "", "New value payload for the storage entry (JSON string)")
	UpdateCmd.Flags().IntVarP(&ttl, "ttl", "t", 0, "Time-to-live in hours (optional)")
	UpdateCmd.Flags().BoolVar(&pretty, "pretty", false, "Pretty print the output (optional)")

	err := UpdateCmd.MarkFlagRequired("key")
	if nil != err {
		fmt.Println(err)
	}

	err = UpdateCmd.MarkFlagRequired("payload")
	if nil != err {
		fmt.Println(err)
	}
}
