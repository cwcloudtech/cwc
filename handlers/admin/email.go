package admin

import (
	"cwc/admin"
	"cwc/utils"
)

func HandleSendEmail(from *string, to *string, bcc *string, subject *string, content *string, templated *bool) {
	c, err := admin.NewClient()
	utils.ExitIfError(err)

	email, err := c.AdminSendEmail(*from, *to, *bcc, *subject, *content, *templated)
	utils.ExitIfError(err)

	if admin.GetDefaultFormat() == "json" {
		utils.PrintJson(email)
	} else {
		utils.PrintRow(*email)
	}
}
