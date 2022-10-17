/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package configure

import (
	"cwc/client"
	"cwc/handlers"
	"cwc/utils"
	"fmt"

	"github.com/spf13/cobra"
)

// configureCmd represents the configure command
var ConfigureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configuring the cli default values like default endpoint, provider and region",
	Long: `This command lets you Configure the cli default values like default endpoint, provider and region
The configure command takes no arguments it will prompt you for each default value
	 `,
	Run: func(cmd *cobra.Command, args []string) {
		default_endpoint := client.GetDefaultEndpoint()
		fmt.Printf(fmt.Sprintf("Default endpoint [%s]: ", default_endpoint))
		new_endpoint := utils.PromptUserForValue()
		if new_endpoint != "" {
			handlers.HandlerSetDefaultEndpoint(new_endpoint)
		}

		default_provider := client.GetDefaultProvider()
		fmt.Printf(fmt.Sprintf("Default provider [%s]: ", default_provider))
		new_default_provider := utils.PromptUserForValue()
		if new_default_provider != "" {
			handlers.HandlerSetDefaultProvider(new_default_provider)
		}

		default_region := client.GetDefaultRegion()
		fmt.Printf(fmt.Sprintf("Default region [%s]: ", default_region))
		new_default_region := utils.PromptUserForValue()
		if new_default_region != "" {
			handlers.HandlerSetDefaultRegion(new_default_region)
		}

	},
}

func init() {
	ConfigureCmd.DisableFlagsInUseLine = true

}
