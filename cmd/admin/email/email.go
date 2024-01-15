package email

import (
	"cwc/handlers/admin"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	from      string
	to        string
	bcc       string
	subject   string
	content   string
	templated bool
)

// createCmd represents the create command
var EmailCmd = &cobra.Command{
	Use:   "email",
	Short: "Send an email",
	Long:  `This command allows you to send email using cwcloud api`,
	Run: func(cmd *cobra.Command, args []string) {
		admin.HandleSendEmail(&from, &to, &bcc, &subject, &content, &templated)
	},
}

func init() {
	EmailCmd.Flags().StringVarP(&from, "from", "f", "", "The expeditor email address")
	EmailCmd.Flags().StringVarP(&to, "to", "t", "", "The recipient email address")
	EmailCmd.Flags().StringVarP(&bcc, "bcc", "b", "", "Bcc email address")
	EmailCmd.Flags().StringVarP(&subject, "subject", "s", "", "The subject")
	EmailCmd.Flags().StringVarP(&content, "content", "c", "", "The content")
	EmailCmd.Flags().BoolVarP(&templated, "templated", "m", false, "Use the cwcloud's template")

	err := EmailCmd.MarkFlagRequired("to")
	if nil != err {
		fmt.Println(err)
	}

	err = EmailCmd.MarkFlagRequired("subject")
	if nil != err {
		fmt.Println(err)
	}

	err = EmailCmd.MarkFlagRequired("content")
	if nil != err {
		fmt.Println(err)
	}
}
