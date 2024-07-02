package uninstall

import (
	"cwc/handlers/user"

	"github.com/spf13/cobra"
)
var (
	releaseName string
	nameSpace   string
)

var UninstallCmd = &cobra.Command{
	Use:   "uninstall [flags]",
	Short: "Uninstall the Helm release for cwcloud application",
	Long:  `Uninstall the Helm release from Kubernetes.`,
	Run: func(cmd *cobra.Command, args []string) {

		user.HandleUninstall(cmd, releaseName, nameSpace)
	},
}
func init() {

	UninstallCmd.Flags().StringVarP(&nameSpace, "namespace", "n", "cwcloud", "Namespace to use for uninstalling deployment (default: cwcloud)")
	UninstallCmd.Flags().StringVarP(&releaseName, "release", "r", "release-0.1.0", "Release name for deployment (default: release-0.1.0)")
}