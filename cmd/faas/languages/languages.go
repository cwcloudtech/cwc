package languages

import (
	"cwc/handlers/user"

	"github.com/spf13/cobra"
)

var pretty bool = false

var LanguagesCmd = &cobra.Command{
	Use:   "languages",
	Short: "List available languages",
	Long:  `This command lets you list your available languages in the cloud`,
	Run: func(cmd *cobra.Command, args []string) {
		user.HandleGetLanguages(&pretty)
	},
}

func init() {
	LanguagesCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "Pretty print the output (optional)")
}
