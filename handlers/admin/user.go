package admin

import (
	"cwc/admin"
	"cwc/config"
	"cwc/utils"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
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
}

func HandleGetUsers(responseUsers *admin.ResponseUsers, pretty *bool) {
	users := responseUsers.Result

	if config.IsPrettyFormatExpected(pretty) {
		displayUsersAsTable(users)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(users)
	} else {
		utils.PrintMultiRow(admin.User{}, responseUsers.Result)
	}
}

func HandleGetUser(responseUser *admin.ResponseUser, pretty *bool) {
	user := responseUser.Result

	if config.IsPrettyFormatExpected(pretty) {
		utils.PrintPretty("User's informations", user)
	} else if config.GetDefaultFormat() == "json" {
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

func displayUsersAsTable(users []admin.User) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Email", "Is Admin", "Confirmed"})

	if len(users) == 0 {
		fmt.Println("No users found")
	} else {
		for _, user := range users {
			table.Append([]string{
				fmt.Sprintf("%d", user.Id),
				user.Email,
				fmt.Sprintf("%t", user.IsAdmin),
				fmt.Sprintf("%t", user.Confirmed),
			})
		}
		table.Render()
	}
}
