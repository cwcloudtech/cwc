/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package create

import (
	"cwc/handlers"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	name         string
	projectId    int
	projectName  string
	environment  string
	instanceType string
	zone         string
	dnsZone      string
)

// createCmd represents the create command
var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a virtual machine in the cloud",
	Long: `This command lets you create a virtual machine in the cloud.
You have to provide the project ID or the project name in which the instance will be associeted.
You also have to provide the environment that will be installed in the virtuals machines.
Other arguments are optional.`,
	Run: func(cmd *cobra.Command, args []string) {
		handlers.HandleAddInstance(&name, &projectId, &projectName, &environment, &instanceType, &zone, &dnsZone)
	},
}

func init() {
	CreateCmd.Flags().StringVarP(&name, "name", "n", "", "The instance name")
	CreateCmd.Flags().IntVarP(&projectId, "project_id", "i", 0, "The project id that you want to associete with the instance")
	CreateCmd.Flags().StringVarP(&projectName, "project_name", "p", "", "The project name that you want to associete with the instance")
	CreateCmd.Flags().StringVarP(&environment, "environment", "e", "", "The environment of the instance (code, wpaas)")
	CreateCmd.Flags().StringVarP(&instanceType, "instance_type", "t", "", "The instance size (DEV1-S, DEV1-M, DEV1-L, DEV1-XL)")
	CreateCmd.Flags().StringVarP(&zone, "zone", "z", "", "instance zone")
	CreateCmd.Flags().StringVarP(&dnsZone, "dns_zone", "d", "", "The root dns zones")

	if err := CreateCmd.MarkFlagRequired("name"); err != nil {
		fmt.Println(err)
	}
	if err := CreateCmd.MarkFlagRequired("environment"); err != nil {
		fmt.Println(err)
	}
	if err := CreateCmd.MarkFlagRequired("zone"); err != nil {
		fmt.Println(err)
	}
}
