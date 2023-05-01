/*
Copyright © 2022 comwork.io contact@comwork.io
*/
package configure

import (
	"cwc/client"
	"cwc/cmd/configure/get"
	"cwc/cmd/configure/set"
	"cwc/handlers/user"

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
		if len(args) == 0 {
			default_endpoint := client.GetDefaultEndpoint()
			fmt.Printf(fmt.Sprintf("Default endpoint [%s]: ", default_endpoint))
			new_endpoint := utils.PromptUserForValue()
			if new_endpoint != "" {
				user.HandlerSetDefaultEndpoint(new_endpoint)
			}

			default_provider := client.GetDefaultProvider()
			fmt.Printf(fmt.Sprintf("Default provider [%s]: ", default_provider))
			new_default_provider := utils.PromptUserForValue()
			if new_default_provider != "" {
				user.HandlerSetDefaultProvider(new_default_provider)
			}

			default_region := client.GetDefaultRegion()
			fmt.Printf(fmt.Sprintf("Default region [%s]: ", default_region))
			new_default_region := utils.PromptUserForValue()
			if new_default_region != "" {
				user.HandlerSetDefaultRegion(new_default_region)
			} else {
				user.HandlerSetDefaultRegion(default_region)
			}
			default_format := client.GetDefaultFormat()
			fmt.Printf(fmt.Sprintf("Default output format [%s]: ", default_format))
			new_default_format := utils.PromptUserForValue()
			if new_default_format != "" {
				user.HandlerSetDefaultFormat(new_default_format)
			}
		}

	},
}

func init() {
	ConfigureCmd.DisableFlagsInUseLine = true
	ConfigureCmd.AddCommand(set.SetCmd)
	ConfigureCmd.AddCommand(get.GetCmd)
}
