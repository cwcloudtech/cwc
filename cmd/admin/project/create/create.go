package create

import (
	adminClient "cwc/admin"
	"cwc/handlers/admin"
	"cwc/utils"
	"fmt"
	"github.com/spf13/cobra"
)

var (
	name         string
	host         string
	token        string
	git          string
	namespace    string
	user_email   string
	project_type string
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
		created_project, add_project_err := c.AdminAddProject(user_email, name, host, token, git, namespace, project_type)
		utils.ExitIfError(add_project_err)
		admin.HandleAddProject(created_project, &user_email, &name, &host, &token, &git, &namespace)
	},
}

func init() {
	CreateCmd.Flags().StringVarP(&name, "name", "n", "", "The project name")
	CreateCmd.Flags().StringVarP(&host, "host", "l", "", "Gitlab host")
	CreateCmd.Flags().StringVarP(&user_email, "user", "u", "", "user associeted with the project")
	CreateCmd.Flags().StringVarP(&token, "token", "t", "", "Gitlab Token")
	CreateCmd.Flags().StringVarP(&git, "git", "g", "", "Git username")
	CreateCmd.Flags().StringVarP(&namespace, "namespace", "s", "", "Gitlab Group ID")
	CreateCmd.Flags().StringVarP(&project_type, "type", "p", "", "Project Type (vm, k8s)")

	err := CreateCmd.MarkFlagRequired("name")
	if nil != err {
		fmt.Println(err)
	}

	err = CreateCmd.MarkFlagRequired("user")
	if nil != err {
		fmt.Println(err)
	}
}
