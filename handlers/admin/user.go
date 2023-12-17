package admin

import (
	"cwc/admin"
	"cwc/utils"
	"fmt"
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
	c, err := admin.NewClient()
	utils.ExitIfError(err)

	responseUsers, err := c.GetAllUsers()
	utils.ExitIfError(err)

	users := responseUsers.Result
	if admin.GetDefaultFormat() == "json" {
		utils.PrintJson(users)
	} else {
		utils.PrintMultiRow(admin.User{}, responseUsers.Result)
	}
}

func HandleGetUser(id *string) {
	c, err := admin.NewClient()
	utils.ExitIfError(err)

	responseUser, err := c.GetUser(*id)
	utils.ExitIfError(err)

	user := responseUser.Result
	if admin.GetDefaultFormat() == "json" {
		utils.PrintJson(user)
	} else {
		utils.PrintRow(user)
	}
}

func HandleDeleteUser(id *string) {
	c, err := admin.NewClient()
	utils.ExitIfError(err)

	err = c.DeleteUser(*id)
	utils.ExitIfError(err)

	fmt.Printf("User %v successfully deleted\n", *id)
}
