package ls

import (
	adminClient "cwc/admin"
	"cwc/handlers/admin"
	"cwc/utils"

	"github.com/spf13/cobra"
)

var (
	formId string
	pretty bool = false
)

var LsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List available contact forms",
	Long: `This command lets you list your available contact forms with cwcloud
This command takes no arguments`,
	Run: func(cmd *cobra.Command, args []string) {
		c, err := adminClient.NewClient()
		utils.ExitIfError(err)
		if utils.IsBlank(formId) {
			forms, err := c.GetAllForms()
			utils.ExitIfError(err)
			admin.HandleGetForms(forms, &pretty)
		} else {
			form, err := c.GetFormById(formId)
			utils.ExitIfError(err)
			admin.HandleGetForm(form, &pretty)
		}
	},
}

func init() {
	LsCmd.Flags().StringVarP(&formId, "id", "f", "", "The contact form id")
	LsCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "Pretty print the output (optional)")
}
