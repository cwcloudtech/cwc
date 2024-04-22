package update

import (
	adminClient "cwc/admin"
	"cwc/handlers/admin"

	"github.com/spf13/cobra"
)

var (
	objectTypeId  string
	interactive bool = false
	objectType    adminClient.ObjectType
)

var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "update a particular object type",
	Long: "This command lets you update a particular object type. To use this command you have to provide the object type ID",
	Run: func(cmd *cobra.Command, args []string) {
		admin.HandleUpdateObjectType(&objectTypeId, &objectType, &interactive)
	},
}

func init() {
	UpdateCmd.Flags().StringVarP(&objectTypeId, "id", "o", "", "The object type ID")
	UpdateCmd.Flags().BoolVar(&objectType.Content.Public, "public", false, "Is the object type public?")
	UpdateCmd.Flags().StringVarP(&objectType.Content.Name, "name", "n", "", "Name of the object type")
	UpdateCmd.Flags().StringVarP(&objectType.Content.DecodingFunction, "decoding_function", "d", "", "Decoding function of the object type")
	UpdateCmd.Flags().StringSliceVarP(&objectType.Content.Triggers, "triggers", "t", []string{}, "Triggers of the object type")
	UpdateCmd.Flags().IntVarP(&objectType.User_id, "user_id", "u", 0, "Owner id")
	UpdateCmd.Flags().BoolVarP(&interactive, "interactive", "i", false, "Interactive mode (optional)")

	err := UpdateCmd.MarkFlagRequired("id")
	if nil != err {
		panic(err)
	}
}

