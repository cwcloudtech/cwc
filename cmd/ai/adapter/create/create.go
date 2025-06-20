package create

import (
	"cwc/client"
	"cwc/handlers/user"
	"cwc/utils"

	"github.com/spf13/cobra"
)

var (
	adapter client.AIAdapterRequest
	pretty  bool
)

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an AI adapter",
	Long:  `This command creates a new AI adapter`,
	Run: func(cmd *cobra.Command, args []string) {
		c, err := client.NewClient()
		utils.ExitIfError(err)

		response, err := c.CreateAIAdapter(adapter)
		utils.ExitIfError(err)

		user.HandleCreateAIAdapter(response, &pretty)
	},
}

func init() {
	CreateCmd.Flags().StringVarP(&adapter.Name, "name", "n", "", "Name of the AI adapter (required)")
	CreateCmd.Flags().StringVarP(&adapter.Url, "url", "u", "", "URL of the AI adapter (required)")
	CreateCmd.Flags().StringVarP(&adapter.Username, "username", "s", "", "Username for authentication (optional)")
	CreateCmd.Flags().StringVarP(&adapter.Password, "password", "w", "", "Password for authentication (optional)")
	CreateCmd.Flags().IntVarP(&adapter.Timeout, "timeout", "t", 30, "Timeout in seconds (optional)")
	CreateCmd.Flags().BoolVarP(&adapter.CheckTls, "check-tls", "k", true, "Check TLS certificate (optional)")
	CreateCmd.Flags().BoolVarP(&adapter.IsPublic, "public", "p", false, "Make adapter public (optional)")
	CreateCmd.Flags().BoolVarP(&pretty, "pretty", "", false, "Pretty print the output (optional)")

	CreateCmd.MarkFlagRequired("name")
	CreateCmd.MarkFlagRequired("url")
}
