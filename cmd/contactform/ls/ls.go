package ls

import (
	"cwc/client"
	"cwc/handlers/user"
	"cwc/utils"

	"github.com/spf13/cobra"
)

var (
	formId string
	pretty bool = false
)

// lsCmd represents the ls command
var LsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List available contact forms",
	Long: `This command lets you list the available contact forms with cwcloud
This command takes no arguments`,
	Run: func(cmd *cobra.Command, args []string) {
		c, err := client.NewClient()
		utils.ExitIfError(err)
		if utils.IsBlank(formId) {
			forms, err := c.GetAllForms()
			utils.ExitIfError(err)
			user.HandleGetForms(forms, &pretty)
		} else {
			form, err := c.GetFormById(*&formId)
			utils.ExitIfError(err)
			user.HandleGetForm(form, &pretty)
		}
	},
}

func init() {
	LsCmd.Flags().StringVarP(&formId, "id", "f", "", "The form id")
	LsCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "Pretty print the output (optional)")
}
