/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package login

import (
	"cwc/handlers"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	access_key string
	secret_key string
)

// loginCmd represents the login command
var LoginCmd = &cobra.Command{
	Use:   "login",
	Short: "Authentificate your self to use the CLI using your access key and your secret key",
	Long: `To use the CLI you have to be authentificated. In order to login it you simply need to provide your access key
and your secret key to the login command`,
	Run: func(cmd *cobra.Command, args []string) {
		handlers.HandleLogin(&access_key, &secret_key)
	},
}

func init() {

	LoginCmd.Flags().StringVarP(&access_key, "access_key", "a", "", "API access key")
	LoginCmd.Flags().StringVarP(&secret_key, "secret_key", "s", "", "API secret key")

	if err := LoginCmd.MarkFlagRequired("access_key"); err != nil {
		fmt.Println(err)
	}
	if err := LoginCmd.MarkFlagRequired("secret_key"); err != nil {
		fmt.Println(err)
	}
}
