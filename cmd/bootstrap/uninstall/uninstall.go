package uninstall

import (
	"cwc/handlers/admin"
	"fmt"

	"github.com/spf13/cobra"
)
var (
	releaseName string
	nameSpace   string
)

var UninstallCmd = &cobra.Command{
	Use:   "uninstall [RELEASE_NAME]",
	Short: "Uninstall the Helm release for cwcloud application",
	Long:  `Uninstall the Helm release from Kubernetes.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			releaseName = args[0]
			fmt.Println(" releaseName: ", releaseName)
			fmt.Println(" nameSpace: ", nameSpace)
			admin.HandleUninstall(cmd, releaseName, nameSpace)
		} else {
			cmd.Help()
		}
	},
}
func init() {
	UninstallCmd.Flags().StringVarP(&nameSpace, "namespace", "n", "cwcloud", "Namespace to use for uninstalling deployment (default: cwcloud)")
}