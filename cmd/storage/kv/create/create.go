package create

import (
	"cwc/client"
	"cwc/handlers/user"
	"cwc/utils"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	storageKV client.StorageKVCreateRequest
	ttl       int
	pretty    bool = false
)

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a key-value entry in your storage",
	Long:  "This command creates a new key-value entry in your cloud storage.",
	Run: func(cmd *cobra.Command, args []string) {
		if ttl > 0 {
			storageKV.TTL = &ttl
		}

		created, err := user.PrepareCreateStorageKV(&storageKV)
		utils.ExitIfError(err)

		err = user.HandleCreateStorageKV(created, &pretty)
		utils.ExitIfError(err)
	},
}

func init() {
	CreateCmd.Flags().StringVarP(&storageKV.Key, "key", "k", "", "Key for the storage entry")
	CreateCmd.Flags().StringP("payload", "p", "", "Value payload for the storage entry (JSON string)")
	CreateCmd.Flags().IntVarP(&ttl, "ttl", "t", 0, "Time-to-live in hours (optional)")
	CreateCmd.Flags().BoolVar(&pretty, "pretty", false, "Pretty print the output (optional)")

	err := CreateCmd.MarkFlagRequired("key")
	if nil != err {
		fmt.Println(err)
	}

	err = CreateCmd.MarkFlagRequired("payload")
	if nil != err {
		fmt.Println(err)
	}

	CreateCmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		payloadStr, _ := cmd.Flags().GetString("payload")
		if utils.IsNotBlank(payloadStr) {
			payload, err := utils.ParseJSONPayload(payloadStr)
			if err != nil {
				return err
			}
			storageKV.Payload = payload
		}
		return nil
	}
}
