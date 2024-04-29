package create

import (
	adminClient "cwc/admin"
	"cwc/handlers/admin"
	"cwc/utils"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	objectType  adminClient.ObjectType
	interactive bool = false
	pretty      bool = false
)

var CreateCmd *cobra.Command

func init() {
	CreateCmd = &cobra.Command{
		Use:   "create",
		Short: "Create an object type in the cloud",
		Long:  "This command lets you create an object type in the cloud.",
		Run:   createCmdRun,
	}

	CreateCmd.Flags().StringVarP(&objectType.Content.Name, "name", "n", "", "Name of the object type")
	CreateCmd.Flags().BoolVar(&objectType.Content.Public, "public", false, "Is the object type public?")
	CreateCmd.Flags().StringVarP(&objectType.Content.DecodingFunction, "decoding_function", "d", "", "Decoding function of the object type")
	CreateCmd.Flags().StringSliceVarP(&objectType.Content.Triggers, "triggers", "t", []string{}, "Triggers of the object type")
	CreateCmd.Flags().IntVarP(&objectType.User_id, "user_id", "u", 0, "Owner id")
	CreateCmd.Flags().BoolVarP(&interactive, "interactive", "i", false, "Interactive mode (optional)")
}

func createCmdRun(cmd *cobra.Command, args []string) {
	if !interactive {
		err := CreateCmd.MarkFlagRequired("decoding_function")
		if nil != err {
			fmt.Println(err)
		}
		err = CreateCmd.MarkFlagRequired("user_id")
		if nil != err {
			fmt.Println(err)
		}
	}
	created_objectType, err := admin.PrepareAddObjectType(&objectType, &interactive)
	utils.ExitIfError(err)
	admin.HandleAddObjectType(created_objectType, &pretty)
}
