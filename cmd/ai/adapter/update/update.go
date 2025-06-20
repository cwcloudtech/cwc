package update

import (
	"cwc/client"
	"cwc/handlers/user"
	"cwc/utils"

	"github.com/spf13/cobra"
)

var (
	adapterId   string
	name        string
	url         string
	username    string
	password    string
	timeout     int
	checkTls    bool
	isPublic    bool
	nameSet     bool
	urlSet      bool
	usernameSet bool
	passwordSet bool
	timeoutSet  bool
	checkTlsSet bool
	publicSet   bool
	pretty      bool
)

var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an AI adapter",
	Long:  `This command updates an existing AI adapter`,
	Run: func(cmd *cobra.Command, args []string) {
		c, err := client.NewClient()
		utils.ExitIfError(err)

		existingAdapter, err := c.GetAIAdapterById(adapterId)
		utils.ExitIfError(err)

		updateRequest := client.AIAdapterRequest{
			Name:     existingAdapter.Name,
			Url:      existingAdapter.Url,
			Username: existingAdapter.Username,
			Password: existingAdapter.Password,
			Headers:  existingAdapter.Headers,
			Timeout:  existingAdapter.Timeout,
			CheckTls: existingAdapter.CheckTls,
			IsPublic: existingAdapter.IsPublic,
		}

		if nameSet {
			updateRequest.Name = name
		}
		if urlSet {
			updateRequest.Url = url
		}
		if usernameSet {
			updateRequest.Username = username
		}
		if passwordSet {
			updateRequest.Password = password
		}
		if timeoutSet {
			updateRequest.Timeout = timeout
		}
		if checkTlsSet {
			updateRequest.CheckTls = checkTls
		}
		if publicSet {
			updateRequest.IsPublic = isPublic
		}

		response, err := c.UpdateAIAdapter(adapterId, updateRequest)
		utils.ExitIfError(err)

		user.HandleUpdateAIAdapter(response, &pretty)
	},
}

func init() {
	UpdateCmd.Flags().StringVarP(&adapterId, "adapter-id", "a", "", "ID of the AI adapter (required)")
	UpdateCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the AI adapter")
	UpdateCmd.Flags().StringVarP(&url, "url", "u", "", "URL of the AI adapter")
	UpdateCmd.Flags().StringVarP(&username, "username", "s", "", "Username for authentication")
	UpdateCmd.Flags().StringVarP(&password, "password", "w", "", "Password for authentication")
	UpdateCmd.Flags().IntVarP(&timeout, "timeout", "t", 0, "Timeout in seconds")
	UpdateCmd.Flags().BoolVarP(&checkTls, "check-tls", "k", false, "Check TLS certificate")
	UpdateCmd.Flags().BoolVarP(&isPublic, "public", "p", false, "Make adapter public")
	UpdateCmd.Flags().BoolVarP(&pretty, "pretty", "", false, "Pretty print the output (optional)")

	UpdateCmd.MarkFlagRequired("adapter-id")

	originalRun := UpdateCmd.Run
	UpdateCmd.Run = func(cmd *cobra.Command, args []string) {
		nameSet = cmd.Flags().Changed("name")
		urlSet = cmd.Flags().Changed("url")
		usernameSet = cmd.Flags().Changed("username")
		passwordSet = cmd.Flags().Changed("password")
		timeoutSet = cmd.Flags().Changed("timeout")
		checkTlsSet = cmd.Flags().Changed("check-tls")
		publicSet = cmd.Flags().Changed("public")

		originalRun(cmd, args)
	}
}
