/*
Copyright © 2022 comwork.io contact.comwork.io

*/
package region

import (
	"cwc/cmd/region/ls"

	"github.com/spf13/cobra"
)

// RegionCmd represents the provider command
var RegionCmd = &cobra.Command{
	Use:   "region",
	Short: "Get informations about available regions that you can associate to virtual machines, s3 buckets and registries",
	Long:  `Get informations about available regions that you can associate to virtual machines, s3 buckets and registries`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	RegionCmd.DisableFlagsInUseLine = true
	RegionCmd.AddCommand(ls.LsCmd)
}
