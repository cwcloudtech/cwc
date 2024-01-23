package ls

import (
	"cwc/handlers/user"

	"github.com/spf13/cobra"
)

var LsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List available config files",
	Long: `This command lets you list your available config files on your machine`,
	Run: func(cmd *cobra.Command, args []string) {
		user.HandleGetConfigFiles()
	},
}
