package ls

import (
	adminClient "cwc/admin"
	"cwc/handlers/admin"
	"cwc/utils"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	numericOnly bool = false
	stringOnly  bool = false
	pretty      bool = false
)

var LsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all available data",
	Long: `This command lets you list all available data in the cloud
This command takes no arguments`,
	Run: func(cmd *cobra.Command, args []string) {
		if (numericOnly && stringOnly) || (!numericOnly && !stringOnly) {
			fmt.Println("Error: You must provide exactly one flag: --numeric or --string")
			return
		}
		c, err := adminClient.NewClient()
		utils.ExitIfError(err)
		if numericOnly {
			numericData, err := c.GetAllNumericData()
			utils.ExitIfError(err)
			admin.HandleGetNumericData(numericData, &pretty)
		} else if stringOnly {
			stringData, err := c.GetAllStringData()
			utils.ExitIfError(err)
			admin.HandleGetStringData(stringData, &pretty)
		}
	},
}

func init() {
	LsCmd.Flags().BoolVarP(&numericOnly, "numeric", "n", false, "Show only numeric data")
	LsCmd.Flags().BoolVarP(&stringOnly, "string", "s", false, "Show only string data")
}
