package admin

import (
	"cwc/cmd/admin/bucket"
	"cwc/cmd/admin/email"
	"cwc/cmd/admin/environment"
	"cwc/cmd/admin/faas"
	"cwc/cmd/admin/instance"
	"cwc/cmd/admin/iot"
	"cwc/cmd/admin/project"
	"cwc/cmd/admin/registry"
	"cwc/cmd/admin/user"

	"github.com/spf13/cobra"
)

// bucketCmd represents the bucket command
var AdminCmd = &cobra.Command{
	Use:   "admin",
	Short: "Administrative command",
	Long:  `Administrative command`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	AdminCmd.DisableFlagsInUseLine = true
	AdminCmd.AddCommand(project.ProjectCmd)
	AdminCmd.AddCommand(registry.RegistryCmd)
	AdminCmd.AddCommand(environment.EnvironmentCmd)
	AdminCmd.AddCommand(bucket.BucketCmd)
	AdminCmd.AddCommand(instance.InstanceCmd)
	AdminCmd.AddCommand(user.UserCmd)
	AdminCmd.AddCommand(email.EmailCmd)
	AdminCmd.AddCommand(faas.FaasCmd)
	AdminCmd.AddCommand(iot.IotCmd)
}
