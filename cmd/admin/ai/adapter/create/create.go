package create

import (
	adminClient "cwc/admin"
	"cwc/handlers/admin"
	"cwc/utils"

	"github.com/spf13/cobra"
)

var (
	name     string
	userId   int
	url      string
	username string
	password string
	timeout  int
	checkTls bool
	isPublic bool
	pretty   bool
)

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new external AI adapter",
	Long:  `This command creates a new external AI adapter and assign it to a specific user.`,
	Run: func(cmd *cobra.Command, args []string) {
		c, err := adminClient.NewClient()
		utils.ExitIfError(err)

		adapter := adminClient.AdminAIAdapterRequest{
			Name:     name,
			UserId:   userId,
			Url:      url,
			Username: username,
			Password: password,
			Headers:  []adminClient.AIAdapterHeader{},
			Timeout:  timeout,
			CheckTls: checkTls,
			IsPublic: isPublic,
		}

		response, err := c.CreateAIAdapter(adapter)
		utils.ExitIfError(err)

		admin.HandleCreateAIAdapter(response, &pretty)
	},
}

func init() {
	CreateCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the AI adapter (required)")
	CreateCmd.Flags().StringVarP(&url, "url", "u", "", "URL of the AI adapter (required)")
	CreateCmd.Flags().IntVarP(&userId, "user-id", "i", 0, "User ID who will own the adapter (required)")
	CreateCmd.Flags().StringVarP(&username, "username", "s", "", "Username for authentication (optional)")
	CreateCmd.Flags().StringVarP(&password, "password", "w", "", "Password for authentication (optional)")
	CreateCmd.Flags().IntVarP(&timeout, "timeout", "t", 30, "Timeout in seconds (default: 30)")
	CreateCmd.Flags().BoolVarP(&checkTls, "check-tls", "c", true, "Check TLS certificates (default: true)")
	CreateCmd.Flags().BoolVarP(&isPublic, "public", "p", false, "Make adapter public (default: false)")
	CreateCmd.Flags().BoolVarP(&pretty, "pretty", "", false, "Pretty print the output (optional)")

	CreateCmd.MarkFlagRequired("name")
	CreateCmd.MarkFlagRequired("user-id")
	CreateCmd.MarkFlagRequired("url")
}
