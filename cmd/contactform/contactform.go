package contactform

import (
	"cwc/cmd/contactform/create"
	"cwc/cmd/contactform/delete"
	"cwc/cmd/contactform/ls"
	"cwc/cmd/contactform/update"

	"github.com/spf13/cobra"
)

// providerCmd represents the provider command
var ContactFormCmd = &cobra.Command{
	Use:   "cf",
	Short: "Manage your contact forms with cwcloud",
	Long: `This command lets you manage your contact forms with cwcloud.
Several actions are associated with this command such listing your available monitors`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	ContactFormCmd.DisableFlagsInUseLine = true
	ContactFormCmd.AddCommand(ls.LsCmd)
	ContactFormCmd.AddCommand(create.CreateCmd)
	ContactFormCmd.AddCommand(update.UpdateCmd)
	ContactFormCmd.AddCommand(delete.DeleteCmd)
}
