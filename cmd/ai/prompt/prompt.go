/*
Copyright © 2022 comwork.io contact@comwork.io
*/
package prompt

import (
	"cwc/handlers/user"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	model   string
	message string
)

// createCmd represents the create command
var PromptCmd = &cobra.Command{
	Use:   "prompt",
	Short: "Send a prompt",
	Long:  `This command allows you to send prompt using cwai api`,
	Run: func(cmd *cobra.Command, args []string) {
		user.HandleSendPrompt(&model, &message)
	},
}

func init() {
	PromptCmd.Flags().StringVarP(&model, "model", "t", "", "The chosen model")
	PromptCmd.Flags().StringVarP(&message, "message", "m", "", "The message input")

	err := PromptCmd.MarkFlagRequired("model")
	if nil != err {
		fmt.Println(err)
	}

	err = PromptCmd.MarkFlagRequired("message")
	if nil != err {
		fmt.Println(err)
	}
}
