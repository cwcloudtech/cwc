package create

import (
	"cwc/client"
	"cwc/handlers/user"
	"cwc/utils"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	deployment client.CreationDeployment
	pretty     bool = false
)

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a deployment in the cloud",
	Long:  "This command lets you create a deployment in the cloud.",
	Run: func(cmd *cobra.Command, args []string) {
		created_deployment, err := user.PrepareAddDeployment(&deployment)
		utils.ExitIfError(err)
		user.HandleAddDeployment(created_deployment, &pretty)
	},
}

func init() {
	CreateCmd.Flags().StringVarP(&deployment.Name, "name", "n", "", "Name of the deployment")
	CreateCmd.Flags().StringVarP(&deployment.Description, "description", "d", "", "Description of the deployment")
	CreateCmd.Flags().IntVarP(&deployment.Cluster_id, "cluster", "c", 0, "Cluster ID of the deployment")
	CreateCmd.Flags().IntVarP(&deployment.Project_id, "project", "i", 0, "Project ID of the deployment")
	CreateCmd.Flags().IntVarP(&deployment.Env_id, "environment", "e", 0, "Environment ID of the deployment")
	CreateCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "Pretty print the output (optional)")

	err := CreateCmd.MarkFlagRequired("name")
	if nil != err {
		fmt.Println(err)
	}

	err = CreateCmd.MarkFlagRequired("description")
	if nil != err {
		fmt.Println(err)
	}

	err = CreateCmd.MarkFlagRequired("cluster")
	if nil != err {
		fmt.Println(err)
	}

	err = CreateCmd.MarkFlagRequired("project")
	if nil != err {
		fmt.Println(err)
	}

	err = CreateCmd.MarkFlagRequired("environment")
	if nil != err {
		fmt.Println(err)
	}

}
