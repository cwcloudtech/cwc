package models

import (
	"cwc/client"
	"cwc/handlers/user"
	"cwc/utils"

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
		c, err := client.NewClient()
		utils.ExitIfError(err)
		models, err := c.GetModels()
		utils.ExitIfError(err)
		user.HandleGetModels(models, &pretty)
	},
}

func init() {
	ModelsCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "Pretty print the output (optional)")
}
