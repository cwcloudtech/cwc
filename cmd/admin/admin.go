/*
Copyright © 2022 comwork.io contact.comwork.io

*/
package admin

import (
	"cwc/cmd/admin/bucket"
	"cwc/cmd/admin/environment"
	"cwc/cmd/admin/user"

	"cwc/cmd/admin/instance"
	"cwc/cmd/admin/project"
	"cwc/cmd/admin/registry"

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


}
