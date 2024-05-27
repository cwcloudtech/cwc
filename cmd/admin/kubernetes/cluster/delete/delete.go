package delete

import (
	"cwc/handlers/admin"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	clusterId string
)

var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a particular Kubernetes cluster",
	Long: `This command lets you delete a particular Kubernetes cluster.
To use this command you have to provide the cluster ID that you want to delete`,
	Run: func(cmd *cobra.Command, args []string) {
		admin.HandleDeleteCluster(&clusterId)
	},
}

func init() {
	DeleteCmd.Flags().StringVarP(&clusterId, "cluster", "c", "", "The cluster id")

	err := DeleteCmd.MarkFlagRequired("cluster")
	if nil != err {
		fmt.Println(err)
	}
}
