package pfw

import (
	"cwc/handlers/user"

	"github.com/spf13/cobra"
)

var (
	nameSpace string
)

var PfwCmd = &cobra.Command{
	Use:   "pfw",
	Short: "Open tunnels and display the graphical interface",
	Long:  `Open port forwarding for the GUI and API.`,
	Run: func(cmd *cobra.Command, args []string) {
		user.HandlePortForward(cmd, nameSpace)
	},
}

func init() {
	PfwCmd.Flags().StringVarP(&nameSpace, "namespace", "n", "cwcloud", "Namespace (default: cwcloud)")
}
