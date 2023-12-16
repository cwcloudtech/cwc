package create

import (
	"cwc/client"
	"cwc/handlers/user"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	trigger client.Trigger
	argumentsValues []string
	interactive bool = false
)

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a trigger in the cloud",
	Long: `This command lets you create a trigger in the cloud.`,
	Run: func(cmd *cobra.Command, args []string) {
		user.HandleAddTrigger(&trigger,&argumentsValues,&interactive)
	},
}

func init() {
	CreateCmd.Flags().StringVarP(&trigger.Content.Function_id, "function_id", "f", "", "The trigger function id")
	CreateCmd.Flags().BoolVarP(&interactive, "interactive", "i", false, "Interactive mode")
	CreateCmd.Flags().StringVarP(&trigger.Content.Name, "name", "n", "", "The trigger name")
	CreateCmd.Flags().StringVarP(&trigger.Kind, "kind", "k", "", "The trigger kind")
	CreateCmd.Flags().StringVarP(&trigger.Content.Cron_expr, "cron_expr", "c", "", "The trigger cron expression")
	CreateCmd.Flags().StringArrayVarP(&argumentsValues, "args", "a", []string{}, "The trigger arguments values")

	if err := CreateCmd.MarkFlagRequired("function_id"); err != nil {
		fmt.Println(err)
	}
	if err := CreateCmd.MarkFlagRequired("name"); err != nil {
		fmt.Println(err)
	}
	if err := CreateCmd.MarkFlagRequired("kind"); err != nil {
		fmt.Println(err)
	}
	if err := CreateCmd.MarkFlagRequired("cron_expr"); err != nil {
		fmt.Println(err)
	}
	if err := CreateCmd.MarkFlagRequired("args"); err != nil {
		fmt.Println(err)
	}
}