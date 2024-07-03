package configure

import (
	"cwc/config"
	"cwc/handlers/user"

	"cwc/utils"
	"fmt"

	"github.com/spf13/cobra"
)

var ConfigureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configuring the cli kubernetes cluster",
	Long:  `This command lets you Configure the cli kubernetes cluster `,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			default_kube_config_path := config.GetDefaultKubeConfigPath()
			fmt.Printf("Kube config path [%s]: ", default_kube_config_path)
			new_default_kube_config_path := utils.PromptUserForValue()

			if utils.IsNotBlank(new_default_kube_config_path) {
				user.HandlerSetDefaultKubeConfigPath(new_default_kube_config_path)
			} else {
				user.HandlerSetDefaultKubeConfigPath(default_kube_config_path)
			}
		}
	},
}

func init() {
	ConfigureCmd.DisableFlagsInUseLine = true

}
