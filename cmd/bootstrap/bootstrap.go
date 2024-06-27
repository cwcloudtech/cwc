package bootstrap

import (
	"cwc/cmd/bootstrap/uninstall"
	
	"cwc/cmd/bootstrap/install"

	"github.com/spf13/cobra"
)



var BootstrapCmd = &cobra.Command{
	Use:   "bootstrap",
	Short: "Automatic Comwork Cloud installation on Kubernetes",
	Long:  `Automatic Comwork Cloud installation on Kubernetes.`,
	Run: func(cmd *cobra.Command, args []string) {

			cmd.Help()

	},
}

func init() {
	

	BootstrapCmd.DisableFlagsInUseLine = true
	BootstrapCmd.AddCommand(install.InstallCmd)
	BootstrapCmd.AddCommand(uninstall.UninstallCmd)
}
