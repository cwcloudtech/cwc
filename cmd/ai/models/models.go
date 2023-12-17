package models

import (
	"cwc/handlers/user"

	"github.com/spf13/cobra"
)

var (
	pretty bool
)

// createCmd represents the create command
var ModelsCmd = &cobra.Command{
	Use:   "models",
	Short: "Get the available models",
	Long:  `This command allows you to list the available models of the cwai api`,
	Run: func(cmd *cobra.Command, args []string) {
		user.HandleGetModels(&pretty)
	},
}

func init() {
	ModelsCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "Pretty print the output (optional)")
}
