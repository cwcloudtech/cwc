package update

import (
	adminClient "cwc/admin"
	"cwc/handlers/admin"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	formId string
	form   adminClient.ContactForm
)

var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a contact form with cwcloud",
	Long:  "This command lets you update a contact form with cwcloud.",
	Run: func(cmd *cobra.Command, args []string) {
		admin.HandleUpdateForm(&formId, &form)
	},
}

func init() {
	UpdateCmd.Flags().StringVarP(&formId, "id", "i", "", "The contact form ID")
	UpdateCmd.Flags().StringVarP(&form.Name, "name", "n", "", "Name of the contact form")
	UpdateCmd.Flags().StringVarP(&form.MailFrom, "from", "f", "", "Expeditor of the contact form")
	UpdateCmd.Flags().StringVarP(&form.MailTo, "to", "t", "", "Recipient of the contact form")
	UpdateCmd.Flags().StringVarP(&form.CopyrightName, "copyright", "c", "", "Copyright name of the contact form")
	UpdateCmd.Flags().StringVarP(&form.LogoUrl, "logo_url", "l", "", "Logo URL of the contact form")
	UpdateCmd.Flags().IntVarP(&form.UserId, "user_id", "I", 0, "User ID")

	err := UpdateCmd.MarkFlagRequired("id")
	if nil != err {
		fmt.Println(err)
	}
}
