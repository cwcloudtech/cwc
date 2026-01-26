package create

import (
	adminClient "cwc/admin"
	"cwc/handlers/admin"
	"cwc/utils"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	form   adminClient.ContactForm
	pretty bool = false
)

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a contact form with cwcloud",
	Long:  "This command lets you create a contact form with cwcloud.",
	Run: func(cmd *cobra.Command, args []string) {
		created_form, err := admin.PrepareAddForm(&form)
		utils.ExitIfError(err)
		admin.HandleAddForm(&created_form, &pretty)
	},
}

func init() {
	CreateCmd.Flags().StringVarP(&form.Name, "name", "n", "", "Name of the contact form")
	CreateCmd.Flags().StringVarP(&form.MailFrom, "from", "f", "", "Expeditor of the contact form")
	CreateCmd.Flags().StringVarP(&form.MailTo, "to", "t", "", "Recipient of the contact form")
	CreateCmd.Flags().StringVarP(&form.CopyrightName, "copyright", "c", "", "Copyright name of the contact form")
	CreateCmd.Flags().StringVarP(&form.LogoUrl, "logo_url", "l", "", "Logo URL of the contact form")
	CreateCmd.Flags().StringVarP(&form.TrustedIps, "trusted_ips", "a", "", "Trusted ips for the contact form")
	CreateCmd.Flags().IntVarP(&form.UserId, "user_id", "i", 0, "User ID")

	err := CreateCmd.MarkFlagRequired("name")
	if nil != err {
		fmt.Println(err)
	}

	err = CreateCmd.MarkFlagRequired("from")
	if nil != err {
		fmt.Println(err)
	}

	err = CreateCmd.MarkFlagRequired("to")
	if nil != err {
		fmt.Println(err)
	}

	err = CreateCmd.MarkFlagRequired("user_id")
	if nil != err {
		fmt.Println(err)
	}
}
