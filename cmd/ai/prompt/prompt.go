package prompt

import (
	"cwc/handlers/user"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	adapter string
	message string
	listId  string
)

// createCmd represents the create command
var PromptCmd = &cobra.Command{
	Use:   "prompt",
	Short: "Send a prompt",
	Long:  `This command allows you to send prompt using cwai api`,
	Run: func(cmd *cobra.Command, args []string) {
		user.HandleSendPrompt(&adapter, &message, &listId)
	},
}

func init() {
	PromptCmd.Flags().StringVarP(&adapter, "adapter", "a", "", "The chosen adapter")
	PromptCmd.Flags().StringVarP(&message, "message", "m", "", "The message input")
	PromptCmd.Flags().StringVarP(&listId, "list", "l", "", "Optional list ID")

	err := PromptCmd.MarkFlagRequired("adapter")
	if nil != err {
		fmt.Println(err)
	}

	err = PromptCmd.MarkFlagRequired("message")
	if nil != err {
		fmt.Println(err)
	}
}
