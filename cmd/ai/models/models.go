package models

import (
	"cwc/handlers/user"

	"github.com/spf13/cobra"
)

var (
	model   string
	message string
)

// createCmd represents the create command
var ModelsCmd = &cobra.Command{
	Use:   "models",
	Short: "Get the available models",
	Long:  `This command allows you to list the available models of the cwai api`,
	Run: func(cmd *cobra.Command, args []string) {
		user.HandleGetModels()
	},
}

func init() {
}
