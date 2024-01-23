package set

import (
	"cwc/handlers/user"
	"cwc/utils"

	"github.com/spf13/cobra"
)

var (
	configFileName string
)

var SetConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Set the config file",
	Long:  `This command lets you select the config file to use`,
	Run: func(cmd *cobra.Command, args []string) {
		user.HandleSetConfigFile(&configFileName)
	},
}

func init() {
	SetConfigCmd.Flags().StringVarP(&configFileName, "config", "c", "", "The config file name to use")

	err := SetConfigCmd.MarkFlagRequired("config")
	utils.ExitIfError(err)
}