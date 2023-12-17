package user

import (
	"cwc/client"
	"cwc/utils"
)

func HandleSendEmail(from *string, to *string, bcc *string, subject *string, content *string) {
	c, err := client.NewClient()
	utils.ExitIfError(err)

	email, err := c.SendEmail(*from, *to, *bcc, *subject, *content)
	utils.ExitIfError(err)

	if client.GetDefaultFormat() == "json" {
		utils.PrintJson(email)
	} else {
		utils.PrintRow(*email)
	}
}
