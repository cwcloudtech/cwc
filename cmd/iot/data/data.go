package data

import (
	"cwc/cmd/iot/data/create"

	"github.com/spf13/cobra"
)

var DataCmd = &cobra.Command{
	Use:   "data",
	Short: "Manage your data in the cloud",
	Long: `This command lets you Manage your data in the cloud.
Several actions are associated with this command such as creating a data`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	DataCmd.DisableFlagsInUseLine = true
	DataCmd.AddCommand(create.CreateCmd)
}
