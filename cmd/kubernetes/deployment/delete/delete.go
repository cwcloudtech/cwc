package delete

import (
	"cwc/handlers/user"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	deploymentId string
)

var DeleteCmd = &cobra.Command{
	Use:   "deleted",
	Short: "Delete a particular deployment",
	Long: `This command lets you delete a particular deployment.
To use this command you have to provide the deployment ID that you want to delete.`,

	Run: func(cmd *cobra.Command, args []string) {
		user.HandleDeleteDeployment(&deploymentId)
	},
}

func init() {
	DeleteCmd.Flags().StringVarP(&deploymentId, "id", "d", "", "The deployment id")

	err := DeleteCmd.MarkFlagRequired("deployment")
	if nil != err {
		fmt.Println(err)
	}
}
