package truncate

import (
	"cwc/handlers/user"

	"github.com/spf13/cobra"
)

var TruncateCmd = &cobra.Command{
	Use:   "truncate",
	Short: "Truncate all triggers",
	Long: `This command lets you truncate all triggers.
This command will delete all triggers that you have.`,
	Run: func(cmd *cobra.Command, args []string) {
		user.HandleTruncateTriggers()
	},

}

func init() {}