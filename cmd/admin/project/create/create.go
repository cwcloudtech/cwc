package create

import (
	"cwc/handlers/admin"
	"fmt"
	adminClient "cwc/admin"
	"github.com/spf13/cobra"
	"cwc/utils"
)

var (
	name       string
	host       string
	token      string
	git        string
	namespace  string
	user_email string
)

// createCmd represents the create command
var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a project in the cloud",
	Long: `This command lets you create a project in the cloud.
You have to provide the project name.
You can also provide your Gitlab host and access token and git username to save the project in another Gitlab instance`,
	Run: func(cmd *cobra.Command, args []string) {
		c, err := adminClient.NewClient()
		utils.ExitIfError(err)
		created_project, err := c.AdminAddProject(user_email, name, host, token, git, namespace)
		admin.HandleAddProject(created_project,&user_email, &name, &host, &token, &git, &namespace)
	},
}

func init() {
	CreateCmd.Flags().StringVarP(&name, "name", "n", "", "The project name")
	CreateCmd.Flags().StringVarP(&host, "host", "l", "", "Gitlab host")
	CreateCmd.Flags().StringVarP(&user_email, "user", "u", "", "user associeted with the project")
	CreateCmd.Flags().StringVarP(&token, "token", "t", "", "Gitlab Token")
	CreateCmd.Flags().StringVarP(&git, "git", "g", "", "Git username")
	CreateCmd.Flags().StringVarP(&namespace, "namespace", "s", "", "Gitlab Group ID")

	err := CreateCmd.MarkFlagRequired("name")
	if nil != err {
		fmt.Println(err)
	}

	err = CreateCmd.MarkFlagRequired("user")
	if nil != err {
		fmt.Println(err)
	}
}
