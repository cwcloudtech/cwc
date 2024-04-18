package update

import (
	"cwc/client"
	"cwc/handlers/user"

	"github.com/spf13/cobra"
)

var (
	functionId  string
	interactive bool = false
	function    client.Function
)

var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a particular function",
	Long: `This command lets you update a particular function.
To use this command you have to provide the function ID`,
	Run: func(cmd *cobra.Command, args []string) {
		user.HandleUpdateFunction(&functionId, &function, &interactive)
	},
}

func init() {
	UpdateCmd.Flags().StringVarP(&functionId, "id", "f", "", "The function ID")
	UpdateCmd.Flags().BoolVarP(&interactive, "interactive", "i", false, "Interactive mode (optional)")
	UpdateCmd.Flags().StringVarP(&function.Content.Language, "language", "l", "", "The function language")
	UpdateCmd.Flags().StringVarP(&function.Content.Regexp, "regexp", "r", "", "The function regexp")
	UpdateCmd.Flags().StringVarP(&function.Content.Name, "name", "n", "", "The function name")
	UpdateCmd.Flags().BoolVarP(&function.Is_public, "is_public", "p", false, "The function is public")
	UpdateCmd.Flags().StringSliceVarP(&function.Content.Args, "args", "g", []string{}, "Arguments of the function")
	UpdateCmd.Flags().StringVarP(&function.Content.Code, "code", "c", "", "The function code")

	err := UpdateCmd.MarkFlagRequired("id")
	if nil != err {
		panic(err)
	}
}
