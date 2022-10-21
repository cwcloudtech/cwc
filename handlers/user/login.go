package user

import (
	"cwc/client"
	"fmt"
	"os"
)

func HandleLogin(access_key *string, secret_key *string) {
	client, _ := client.NewClient()
	err := client.UserLogin(*access_key, *secret_key)
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("You are successfully logged in\n")
}
