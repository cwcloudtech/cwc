package admin

import (
	"cwc/admin"
	"cwc/utils"
	"fmt"
	"os"
)

func HandleSendEmail(from *string, to *string, bcc *string, subject *string, content *string, templated *bool) {
	client, err := admin.NewClient()
	email, err := client.AdminSendEmail(*from, *to, *bcc, *subject, *content, *templated)
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	if admin.GetDefaultFormat() == "json" {
		utils.PrintJson(email)
	} else {
		utils.PrintRow(*email)
	}

}
