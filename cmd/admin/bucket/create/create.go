package create

import (
	"cwc/handlers/admin"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	name       string
	reg_type   string
	user_email string
)

// createCmd represents the create command
var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a bucket in the cloud",
	Long:  `This command lets you create a bucket in the cloud`,
	Run: func(cmd *cobra.Command, args []string) {
		admin.HandleAddBucket(&user_email, &name, &reg_type)
	},
}

func init() {
	CreateCmd.Flags().StringVarP(&name, "name", "n", "", "The bucket name")
	CreateCmd.Flags().StringVarP(&user_email, "user", "u", "", "user associeted with the project")
	CreateCmd.Flags().StringVarP(&reg_type, "type", "t", "", "The bucket type (private/public-read)")

	err := CreateCmd.MarkFlagRequired("name")
	if nil != err {
		fmt.Println(err)
	}

	err = CreateCmd.MarkFlagRequired("user")
	if nil != err {
		fmt.Println(err)
	}
}
