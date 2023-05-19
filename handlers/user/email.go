package user

import (
	"cwc/client"
	"cwc/utils"
	"fmt"
	"os"
)

func HandleSendEmail(from *string, to *string, bcc *string, subject *string, content *string) {
	c, err := client.NewClient()
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	email, err := c.SendEmail(*from, *to, *bcc, *subject, *content)
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	if client.GetDefaultFormat() == "json" {
		utils.PrintJson(email)
	} else {
		utils.PrintRow(*email)
	}

}
