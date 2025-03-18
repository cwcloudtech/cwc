package prompt

import (
	"cwc/handlers/user"
	"fmt"

	"github.com/spf13/cobra"
)

var DetailsCmd = &cobra.Command{
	Use:   "details",
	Short: "Get details about a prompt list",
	Long:  `This command retrieves detailed information about a specific prompt list`,
	Run: func(cmd *cobra.Command, args []string) {
		user.HandleGetPromptDetails(&listId, &pretty)
	},
}

func init() {
	DetailsCmd.Flags().StringVarP(&listId, "list", "l", "", "List ID to get details for")
	DetailsCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "Pretty print the output (optional)")

	err := DetailsCmd.MarkFlagRequired("list")
	if nil != err {
		fmt.Println(err)
	}
}