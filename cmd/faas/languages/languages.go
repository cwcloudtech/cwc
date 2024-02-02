package languages

import (
	"cwc/client"
	"cwc/handlers/user"
	"cwc/utils"

	"github.com/spf13/cobra"
)

var pretty bool = false

var LanguagesCmd = &cobra.Command{
	Use:   "languages",
	Short: "List available languages",
	Long:  `This command lets you list your available languages in the cloud`,
	Run: func(cmd *cobra.Command, args []string) {
		languages, err := client.GetLanguages()
		utils.ExitIfError(err)
		user.HandleGetLanguages(languages, &pretty)
	},
}

func init() {
	LanguagesCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "Pretty print the output (optional)")
}
