package ls

import (
	"cwc/client"
	"cwc/handlers/user"
	"cwc/utils"

	"github.com/spf13/cobra"
)

var (
	adapterId string
	pretty    bool
	external  bool
)

var LsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List Default and public external AI adapters or ones created by you",
	Long:  `This command lists AI adapters (both default and public externals, or only externals created by you with --external flag). Use --adapter-id to get details of a specific adapter.`,
	Run: func(cmd *cobra.Command, args []string) {
		c, err := client.NewClient()
		utils.ExitIfError(err)

		if utils.IsNotBlank(adapterId) {
			adapter, err := c.GetAIAdapterById(adapterId)
			utils.ExitIfError(err)
			user.HandleGetAIAdapter(adapter, &pretty)
			return
		}

		if external {
			adapters, err := c.GetExternalAIAdapters()
			utils.ExitIfError(err)
			user.HandleGetExternalAIAdapters(adapters, &pretty)
		} else {
			adapters, err := c.GetAiAdapters()
			utils.ExitIfError(err)
			user.HandleGetAiAdapters(adapters, &pretty)
		}
	},
}

func init() {
	LsCmd.Flags().StringVarP(&adapterId, "adapter-id", "a", "", "ID of a specific AI adapter to get details (optional)")
	LsCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "Pretty print the output (optional)")
	LsCmd.Flags().BoolVarP(&external, "external", "e", false, "List only external adapters (optional, ignored when --adapter-id is used)")
}
