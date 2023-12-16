package truncate

import (
	"cwc/handlers/user"

	"github.com/spf13/cobra"
)

var TruncateCmd = &cobra.Command{
	Use:   "truncate",
	Short: "Truncate all invocations",
	Long: `This command lets you truncate all invocations.
This command will delete all invocations that you have.`,
	Run: func(cmd *cobra.Command, args []string) {
		user.HandleTruncateInvocations()
	},
}

func init() {
}
