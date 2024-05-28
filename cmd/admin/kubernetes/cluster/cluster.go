package cluster

import (
	"cwc/cmd/admin/kubernetes/cluster/delete"
	"cwc/cmd/admin/kubernetes/cluster/ls"

	"github.com/spf13/cobra"
)

var ClusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Manage your kubernetes clusters in the cloud",
	Long: `This command lets you Manage your kubernetes clusters in the cloud.
Several actions are associated with this command such listing your available clusters`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	ClusterCmd.DisableFlagsInUseLine = true
	ClusterCmd.AddCommand(ls.LsCmd)
	ClusterCmd.AddCommand(delete.DeleteCmd)
}
