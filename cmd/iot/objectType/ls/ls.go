package ls

import (
	"cwc/client"
	"cwc/handlers/user"
	"cwc/utils"

	"github.com/spf13/cobra"
)

var (
	objectTypeId string
	pretty       bool = false
)

var LsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List available object types",
	Long: `This command lets you list your available object types in the cloud
This command takes no arguments`,
	Run: func(cmd *cobra.Command, args []string) {
		c, err := client.NewClient()
		utils.ExitIfError(err)
		if utils.IsBlank(objectTypeId) {
			objectTypes, err := c.GetAllObjectTypes()
			utils.ExitIfError(err)
			user.HandleGetObjectTypes(objectTypes, &pretty)
		} else {
			objectType, err := c.GetObjectTypeById(*&objectTypeId)
			utils.ExitIfError(err)
			user.HandleGetObjectType(objectType, &pretty)
		}
	},
}

func init() {
	LsCmd.Flags().StringVarP(&objectTypeId, "id", "o", "", "The object type id")
	LsCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "Pretty print the output (optional)")
}