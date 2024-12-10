package pfw

import (
	"cwc/handlers/user"

	"github.com/spf13/cobra"
)

var (
	nameSpace string
	openshift bool
)

var PfwCmd = &cobra.Command{
	Use:   "pfw",
	Short: "Open tunnels and display the graphical interface",
	Long:  `Open port forwarding for the GUI and API.`,
	Run: func(cmd *cobra.Command, args []string) {
		user.HandlePortForward(cmd, nameSpace, openshift)
	},
}

func init() {
	PfwCmd.Flags().StringVarP(&nameSpace, "namespace", "n", "cwcloud", "Namespace (default: cwcloud)")
	PfwCmd.Flags().BoolVarP(&openshift, "openshift", "o", false, "Use openshift cli instead of kubectl")
}
