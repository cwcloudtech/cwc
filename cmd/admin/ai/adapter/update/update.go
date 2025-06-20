package update

import (
	adminClient "cwc/admin"
	"cwc/handlers/admin"
	"cwc/utils"

	"github.com/spf13/cobra"
)

var (
	adapterId string
	name      string
	userId    int
	url       string
	username  string
	password  string
	timeout   int
	checkTls  bool
	isPublic  bool
	pretty    bool
)

var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an AI adapter (admin)",
	Long:  `This command updates an existing AI adapter.`,
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

		response, err := c.UpdateAIAdapter(adapterId, adapter)
		utils.ExitIfError(err)

		admin.HandleUpdateAIAdapter(response, &pretty)
	},
}

func init() {
	UpdateCmd.Flags().StringVarP(&adapterId, "adapter-id", "a", "", "ID of the AI adapter to update (required)")
	UpdateCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the AI adapter (required)")
	UpdateCmd.Flags().StringVarP(&url, "url", "u", "", "URL of the AI adapter (required)")
	UpdateCmd.Flags().IntVarP(&userId, "user-id", "i", 0, "User ID who will own the adapter (required)")
	UpdateCmd.Flags().StringVarP(&username, "username", "s", "", "Username for authentication (optional)")
	UpdateCmd.Flags().StringVarP(&password, "password", "w", "", "Password for authentication (optional)")
	UpdateCmd.Flags().IntVarP(&timeout, "timeout", "t", 30, "Timeout in seconds (default: 30)")
	UpdateCmd.Flags().BoolVarP(&checkTls, "check-tls", "c", true, "Check TLS certificates (default: true)")
	UpdateCmd.Flags().BoolVarP(&isPublic, "public", "p", false, "Make adapter public (default: false)")
	UpdateCmd.Flags().BoolVarP(&pretty, "pretty", "", false, "Pretty print the output (optional)")

	UpdateCmd.MarkFlagRequired("adapter-id")
}
