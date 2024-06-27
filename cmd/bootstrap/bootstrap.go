package bootstrap

import (
	"cwc/cmd/bootstrap/uninstall"
	"cwc/cmd/bootstrap/install"
	// "cwc/handlers/admin"

	"github.com/spf13/cobra"
)

// var (
// 	flagVerbose   bool
// 	clusterIP     string
// 	nameSpace     string
// 	dbPassword    string
// 	otherValues   []string
// 	releaseName   string 
// )

var BootstrapCmd = &cobra.Command{
	Use:   "bootstrap",
	Short: "Automatic Comwork Cloud installation on Kubernetes",
	Long:  `Automatic Comwork Cloud installation on Kubernetes.`,
	// Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// releaseName = args[0]
		// if len(args) == 0{
			cmd.Help()
		// } else {
		// 	admin.HandleBootstrap(cmd, clusterIP,releaseName, dbPassword, nameSpace, otherValues, flagVerbose)
		// }
	},
}

func init() {
	
	// BootstrapCmd.Flags().StringVarP(&clusterIP, "cluster-ip", "c", "", "Kubernetes Cluster IP address to connect")
	// BootstrapCmd.Flags().StringVarP(&dbPassword, "db-password", "d", "", "Database password for authentication")
	// BootstrapCmd.Flags().StringVarP(&nameSpace, "name-space", "n", "cwcloud", "Namespace to use for deployment (default: cwcloud)")
	// BootstrapCmd.Flags().StringArrayVarP(&otherValues, "values", "p", []string{}, "Array of values to override other configurations")

	// // Mark required flags
	// _ = BootstrapCmd.MarkFlagRequired("cluster-ip")
	// _ = BootstrapCmd.MarkFlagRequired("db-password")

	// BootstrapCmd.Flags().BoolVarP(&flagVerbose, "verbose", "v", false, "Enable verbose output")
	BootstrapCmd.DisableFlagsInUseLine = true
	BootstrapCmd.AddCommand(install.InstallCmd)
	BootstrapCmd.AddCommand(uninstall.UninstallCmd)
}
