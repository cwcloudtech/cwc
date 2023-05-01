/*
Copyright © 2022 comwork.io contact@comwork.io
*/
package email

import (
	"cwc/handlers/admin"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	from_email string
	to_email   string
	subject    string
	content    string
	templated  bool
)

// createCmd represents the create command
var EmailCmd = &cobra.Command{
	Use:   "email",
	Short: "Send an email",
	Long:  `This command allows you to send email using cwcloud api`,
	Run: func(cmd *cobra.Command, args []string) {
		admin.HandleSendEmail(&from_email, &to_email, &subject, &content, &templated)
	},
}

func init() {
	EmailCmd.Flags().StringVarP(&from_email, "from_email", "f", "", "The expeditor email address")
	EmailCmd.Flags().StringVarP(&to_email, "to_email", "t", "", "The recipient email address")
	EmailCmd.Flags().StringVarP(&content, "bcc_email", "b", "", "Bcc email address")
	EmailCmd.Flags().StringVarP(&subject, "subject", "s", "", "The subject")
	EmailCmd.Flags().StringVarP(&content, "content", "c", "", "The subject")
	EmailCmd.Flags().BoolVarP(&templated, "templated", "m", false, "The subject")

	if err := EmailCmd.MarkFlagRequired("to_email"); err != nil {
		fmt.Println(err)
	}
	if err := EmailCmd.MarkFlagRequired("subject"); err != nil {
		fmt.Println(err)
	}
	if err := EmailCmd.MarkFlagRequired("content"); err != nil {
		fmt.Println(err)
	}
}
