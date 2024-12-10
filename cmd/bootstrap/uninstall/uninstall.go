package uninstall

import (
	"cwc/handlers/user"

	"github.com/spf13/cobra"
)

var (
	releaseName string
	nameSpace   string
	force       bool
	openshift   bool
)

var UninstallCmd = &cobra.Command{
	Use:   "uninstall [flags]",
	Short: "Uninstall the Helm release for cwcloud application",
	Long:  `Uninstall the Helm release from Kubernetes.`,
	Run: func(cmd *cobra.Command, args []string) {
		user.HandleUninstall(cmd, releaseName, nameSpace, force, openshift)
	},
}

func init() {
	UninstallCmd.Flags().StringVarP(&nameSpace, "namespace", "n", "cwcloud", "Namespace to use for uninstalling deployment (default: cwcloud)")
	UninstallCmd.Flags().StringVarP(&releaseName, "release", "r", "release-0.1.0", "Release name for deployment (default: release-0.1.0)")
	UninstallCmd.Flags().BoolVarP(&force, "force", "f", false, "Force remove every resources on the namespace")
	UninstallCmd.Flags().BoolVarP(&openshift, "openshift", "o", false, "Use openshift cli instead of kubectl")
}
