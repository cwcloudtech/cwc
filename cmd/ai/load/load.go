package load

import (
	"cwc/handlers/user"

	"github.com/spf13/cobra"
)

var LoadCmd = &cobra.Command{
	Use:   "load [model]",
	Short: "Load a model",
	Long:  `This command allows you to load a model using cwai api`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		model := args[0]
		user.HandleLoadModel(&model)
	},
}
