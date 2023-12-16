package admin

import (
	"cwc/admin"
	"cwc/utils"
	"fmt"
	"os"
)

type User struct {
	Id                 int    `json:"id"`
	Email              string `json:"email"`
	RegistrationNumber string `json:"registration_number"`
	Address            string `json:"address"`
	CompanyName        string `json:"company_name"`
	ContactInfo        string `json:"contact_info"`
	IsAdmin            bool   `json:"is_admin"`
	Confirmed          bool   `json:"confirmed"`
	Billable           bool   `json:"billable"`
}

func HandleGetUsers() {

	client, err := admin.NewClient()

	responseUsers, err := client.GetAllUsers()

	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	users := responseUsers.Result
	if admin.GetDefaultFormat() == "json" {
		utils.PrintJson(users)
	} else {
		utils.PrintMultiRow(admin.User{}, responseUsers.Result)
	}
	return

}

func HandleGetUser(id *string) {

	client, err := admin.NewClient()

	responseUser, err := client.GetUser(*id)

	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	user := responseUser.Result

	if admin.GetDefaultFormat() == "json" {
		utils.PrintJson(user)
	} else {
		utils.PrintRow(user)
	}

}

func HandleDeleteUser(id *string) {
	client, err := admin.NewClient()

	err = client.DeleteUser(*id)
	if nil != err {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("User %v successfully deleted\n", *id)
}
