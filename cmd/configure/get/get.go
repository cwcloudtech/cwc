package get

import (
	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve informations about your default configurations",
	Long:  `This command lets you retrieve informations about your default configurations`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	GetCmd.DisableFlagsInUseLine = true
	GetCmd.AddCommand(GetEndpointCmd)
	GetCmd.AddCommand(GetFormatCmd)
	GetCmd.AddCommand(GetProviderCmd)
	GetCmd.AddCommand(GetRegionCmd)
}
