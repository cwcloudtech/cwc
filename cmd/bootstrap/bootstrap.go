package bootstrap

import (
	"github.com/spf13/cobra"

	"cwc/cmd/bootstrap/uninstall"
	"cwc/handlers/user"
)

var (
	flagVerbose bool
	nameSpace   string
	otherValues []string
	releaseName string
)

var BootstrapCmd = &cobra.Command{
	Use:   "bootstrap [flags]",
	Short: "Automatic Comwork Cloud installation on Kubernetes",
	Long:  `Automatic Comwork Cloud installation on Kubernetes.`,
	Run: func(cmd *cobra.Command, args []string) {
		user.HandleBootstrap(cmd, releaseName, nameSpace, otherValues, flagVerbose)
	},
}

func init() {
	BootstrapCmd.DisableFlagsInUseLine = true
	BootstrapCmd.Flags().StringVarP(&releaseName, "release", "r", "release-0.1.0", "Release name for deployment (default: release-0.1.0)")
	BootstrapCmd.Flags().StringVarP(&nameSpace, "name-space", "n", "cwcloud", "Namespace to use for deployment (default: cwcloud)")
	BootstrapCmd.Flags().StringArrayVarP(&otherValues, "value", "p", []string{}, "Values to override other configurations (e.g. --value key=value --value key2=value2)")

	BootstrapCmd.AddCommand(uninstall.UninstallCmd)
}
