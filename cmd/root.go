/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"cwc/cmd/bucket"
	"cwc/cmd/configure"
	"cwc/cmd/dnszones"
	"cwc/cmd/environment"
	"cwc/cmd/instance"
	"cwc/cmd/login"
	"cwc/cmd/project"
	"cwc/cmd/provider"
	"cwc/cmd/region"
	"cwc/cmd/registry"
	"cwc/handlers"
	"os"

	"github.com/spf13/cobra"
)

var (
	fversion bool
)
var Version = "dev"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cwc",
	Short: "A Command Line interface to manage your cloud resources in comwork cloud",
	Long:  `A Command Line interface to manage your cloud resources in comwork cloud`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if fversion {

			handlers.HandleVersion(Version)
		} else {
			cmd.Help()
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cwc.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	rootCmd.Flags().BoolVarP(&fversion, "version", "v", false, "The cli version")
	rootCmd.AddCommand(project.ProjectCmd)
	rootCmd.AddCommand(bucket.BucketCmd)
	rootCmd.AddCommand(instance.InstanceCmd)
	rootCmd.AddCommand(registry.RegistryCmd)
	rootCmd.AddCommand(login.LoginCmd)
	rootCmd.AddCommand(provider.ProviderCmd)
	rootCmd.AddCommand(environment.EnvironmentCmd)
	rootCmd.AddCommand(region.RegionCmd)
	rootCmd.AddCommand(dnszones.DnsZonesCmd)
	rootCmd.AddCommand(configure.ConfigureCmd)

}
