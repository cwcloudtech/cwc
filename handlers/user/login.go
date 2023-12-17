package user

import (
	"cwc/client"
	"cwc/utils"
	"fmt"
)

func HandleLogin(access_key *string, secret_key *string) {
	client, _ := client.NewClient()
	err := client.UserLogin(*access_key, *secret_key)
	utils.ExitIfError(err)

	fmt.Println("You are successfully logged in")
}
