package contactform

import (
	"cwc/cmd/admin/contactform/create"
	"cwc/cmd/admin/contactform/delete"
	"cwc/cmd/admin/contactform/ls"
	"cwc/cmd/admin/contactform/update"

	"github.com/spf13/cobra"
)

var ContactFormCmd = &cobra.Command{
	Use:   "cf",
	Short: "Manage your monitors with cwcloud",
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
