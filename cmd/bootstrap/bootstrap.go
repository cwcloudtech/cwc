package bootstrap

import (
	"github.com/spf13/cobra"

	"cwc/cmd/bootstrap/pfw"
	"cwc/cmd/bootstrap/uninstall"
	"cwc/handlers/user"
)

var (
	flagVerbose bool
	nameSpace   string
	otherValues []string
	releaseName string
	keepDir     bool
	recreateNs  bool
)

var BootstrapCmd = &cobra.Command{
	Use:   "bootstrap [flags]",
	Short: "CWCloud installation on Kubernetes",
	Long:  `CWCloud installation on Kubernetes.`,
	Run: func(cmd *cobra.Command, args []string) {
		user.HandleBootstrap(cmd, releaseName, nameSpace, otherValues, flagVerbose, keepDir, recreateNs)
	},
}

func init() {
	BootstrapCmd.DisableFlagsInUseLine = true
	BootstrapCmd.Flags().StringVarP(&releaseName, "release", "r", "release-0.1.0", "Release name for deployment (default: release-0.1.0)")
	BootstrapCmd.Flags().StringVarP(&nameSpace, "namespace", "n", "cwcloud", "Namespace to use for deployment (default: cwcloud)")
	BootstrapCmd.Flags().BoolVarP(&keepDir, "keep-dir", "k", false, "Keep the local helm directory")
	BootstrapCmd.Flags().BoolVarP(&recreateNs, "recreate-ns", "d", false, "Recreate the namespace")
	BootstrapCmd.Flags().StringArrayVarP(&otherValues, "value", "p", []string{}, `Values to override other configurations (e.g. --value key=value --value key2=value2)

Example:
  All applications are enabled by default. To disable some applications, use this format:

    -p applicationName.enabled=false


  Example:
    cwc bootstrap -p scheduler.enabled=false \
                  -p consumer.enabled=false
	`)

	BootstrapCmd.AddCommand(uninstall.UninstallCmd)
	BootstrapCmd.AddCommand(pfw.PfwCmd)
}
