package admin

import (
	"cwc/admin"
	"fmt"
	"os"
)

type User struct {
	Id          int    `json:"id"`
	Email        string `json:"email"`
	RegistrationNumber       string `json:"registration_number"`
	Address    string `json:"address"`
	CompanyName   string   `json:"company_name"`
	ContactInfo string `json:"contact_info"`
	IsAdmin bool `json:"is_admin"`
	Confirmed bool `json:"confirmed"`
	Billable bool `json:"billable"`

}

func HandleGetUsers() {

	client, err := admin.NewClient()

	responseUsers, err := client.GetAllUsers()

	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("ID\temail\tregistration\taddress\tcompany\tcontact\tadmin\t confirmed\tbillable\n")
	users := responseUsers.Result
	for _, user := range users {
		fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t\n", user.Id, user.Email, user.RegistrationNumber, user.Address,user.CompanyName,user.ContactInfo,user.IsAdmin,user.Confirmed,user.Billable)
	}
	return

}