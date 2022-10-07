package handlers

import (
	"cwc/client"
	"flag"
	"fmt"
	"os"
)

func HandleLogin(loginCmd *flag.FlagSet, access_key *string, secret_key *string) {

	loginCmd.Parse(os.Args[2:])
	if *access_key == "" || *secret_key == "" {
		fmt.Println("email and password are required to login")
		loginCmd.PrintDefaults()
		os.Exit(1)
	}
	client := client.NewClient()

	err := client.UserLogin(*access_key, *secret_key)
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("You are successfully logged in\n")
}
