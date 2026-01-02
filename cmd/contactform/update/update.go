package update

import (
	"cwc/client"
	"cwc/handlers/user"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	formId string
	form   client.ContactForm
)

var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a particular contact form",
	Long: `This command lets you update a particular contact form.
To use this command you have to provide the contact form ID`,
	Run: func(cmd *cobra.Command, args []string) {
		user.HandleUpdateForm(&formId, &form)
	},
}

func init() {
	UpdateCmd.Flags().StringVarP(&formId, "id", "m", "", "The contact form ID")
	UpdateCmd.Flags().StringVarP(&form.Name, "name", "n", "", "Name of the contact form")
	UpdateCmd.Flags().StringVarP(&form.MailFrom, "from", "f", "", "Expeditor of the contact form")
	UpdateCmd.Flags().StringVarP(&form.MailTo, "to", "t", "", "Recipient of the contact form")
	UpdateCmd.Flags().StringVarP(&form.CopyrightName, "copyright", "c", "", "Copyright name of the contact form")
	UpdateCmd.Flags().StringVarP(&form.LogoUrl, "logo_url", "l", "", "Logo URL of the contact form")

	err := UpdateCmd.MarkFlagRequired("id")
	if nil != err {
		fmt.Println(err)
	}
}
