package adapters

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
var AiAdaptersCmd = &cobra.Command{
	Use:   "adapters",
	Short: "Get the available adapters",
	Long:  `This command allows you to list the available adapters of the cwai api`,
	Run: func(cmd *cobra.Command, args []string) {
		c, err := client.NewClient()
		utils.ExitIfError(err)
		adapters, err := c.GetAiAdapters()
		utils.ExitIfError(err)
		user.HandleGetAiAdapters(adapters, &pretty)
	},
}

func init() {
	AiAdaptersCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "Pretty print the output (optional)")
}
