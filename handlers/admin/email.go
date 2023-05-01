package admin

import (
	"cwc/admin"
	"cwc/utils"
	"fmt"
	"os"
)

func HandleSendEmail(from_email *string, to_email *string, subject *string, content *string, templated *bool) {
	client, err := admin.NewClient()
	email, err := client.AdminSendEmail(*from_email, *to_email, *subject, *content, *templated)
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	if admin.GetDefaultFormat() == "json" {
		utils.PrintJson(email)
	} else {
		utils.PrintRow(*email)
	}

}
