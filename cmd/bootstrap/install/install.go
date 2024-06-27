package install

import (
	"cwc/handlers/admin"

	"github.com/spf13/cobra"
)

var (
	flagVerbose bool
	clusterIP   string
	nameSpace   string
	dbPassword  string
	otherValues []string
	releaseName string
)

var InstallCmd = &cobra.Command{
	Use:   "install [RELEASE_NAME]",
	Short: "Automatic Comwork Cloud installation on Kubernetes",
	Long:  `Automatic Comwork Cloud installation on Kubernetes.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 1 {
			releaseName = args[0]
			admin.HandleBootstrap(cmd, clusterIP, releaseName, dbPassword, nameSpace, otherValues, flagVerbose)

		} else {
			cmd.Help()
		}
	},
}

func init() {

	InstallCmd.Flags().StringVarP(&clusterIP, "cluster-ip", "c", "", "Kubernetes Cluster IP address to connect")
	InstallCmd.Flags().StringVarP(&dbPassword, "db-password", "d", "", "Database password for authentication")
	InstallCmd.Flags().StringVarP(&nameSpace, "name-space", "n", "cwcloud", "Namespace to use for deployment (default: cwcloud)")
	InstallCmd.Flags().StringArrayVarP(&otherValues, "value", "p", []string{}, "Values to override other configurations (e.g. --value key=value --value key2=value2)")

	// Mark required flags
	_ = InstallCmd.MarkFlagRequired("cluster-ip")
	_ = InstallCmd.MarkFlagRequired("db-password")

}
