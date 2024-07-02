package configure

import (
	"cwc/cmd/configure/get"
	importConfig "cwc/cmd/configure/import"
	"cwc/cmd/configure/ls"
	"cwc/cmd/configure/set"
	switchConfig "cwc/cmd/configure/switch"
	"cwc/config"
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
The configure command takes no arguments it will prompt you for each default value`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			default_endpoint := config.GetDefaultEndpoint()
			fmt.Printf("Default endpoint [%s]: ", default_endpoint)
			new_endpoint := utils.PromptUserForValue()
			if utils.IsNotBlank(new_endpoint) {
				user.HandlerSetDefaultEndpoint(new_endpoint)
			}

			default_provider := config.GetDefaultProvider()
			fmt.Printf("Default provider [%s]: ", default_provider)
			new_default_provider := utils.PromptUserForValue()
			if utils.IsNotBlank(new_default_provider) {
				user.HandlerSetDefaultProvider(new_default_provider)
			}

			default_region := config.GetDefaultRegion()
			fmt.Printf("Default region [%s]: ", default_region)
			new_default_region := utils.PromptUserForValue()

			if utils.IsNotBlank(new_default_region) {
				user.HandlerSetDefaultRegion(new_default_region)
			} else {
				user.HandlerSetDefaultRegion(default_region)
			}

			default_kube_config_path := config.GetDefaultKubeConfigPath()
			fmt.Printf("Default kube config path [%s]: ", default_kube_config_path)
			new_default_kube_config_path := utils.PromptUserForValue()

			if utils.IsNotBlank(new_default_kube_config_path) {
				user.HandlerSetDefaultKubeConfigPath(new_default_kube_config_path)
			} else {
				user.HandlerSetDefaultKubeConfigPath(default_kube_config_path)
			}

			default_format := config.GetDefaultFormat()
			fmt.Printf("Default output format [%s]: ", default_format)
			new_default_format := utils.PromptUserForValue()
			if utils.IsNotBlank(new_default_format) {
				user.HandlerSetDefaultFormat(new_default_format)
			}
		}
	},
}

func init() {
	ConfigureCmd.DisableFlagsInUseLine = true
	ConfigureCmd.AddCommand(set.SetCmd)
	ConfigureCmd.AddCommand(get.GetCmd)
	ConfigureCmd.AddCommand(ls.LsCmd)
	ConfigureCmd.AddCommand(switchConfig.SwitchConfigCmd)
	ConfigureCmd.AddCommand(importConfig.ImportConfigCmd)
}
