package create

import (
	adminClient "cwc/admin"
	"cwc/handlers/admin"
	"cwc/utils"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var (
	monitor    adminClient.Monitor
	pretty     bool = false
	rawHeaders string
)

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a monitor in the cloud",
	Long:  "This command lets you create a monitor in the cloud.",
	Run: func(cmd *cobra.Command, args []string) {
		if rawHeaders != "" {
			headers, err := parseHeaders(rawHeaders)
			utils.ExitIfError(err)
			monitor.Headers = headers
		} else {
			monitor.Headers = []adminClient.Header{}
		}

		created_monitor, err := admin.PrepareAddMonitor(&monitor)
		utils.ExitIfError(err)
		admin.HandleAddMonitor(&created_monitor, &pretty)
	},
}

func init() {
	CreateCmd.Flags().StringVarP(&monitor.Type, "type", "y", "http", "Type of the monitor (http, tcp, icmp)")
	CreateCmd.Flags().StringVarP(&monitor.Name, "name", "n", "", "Name of the monitor")
	CreateCmd.Flags().StringVarP(&monitor.Family, "family", "f", "", "Family of the monitor")
	CreateCmd.Flags().StringVarP(&monitor.Url, "url", "u", "", "Url of the monitor")
	CreateCmd.Flags().StringVarP(&monitor.Method, "method", "m", "GET", "Method of the request in the monitor (GET, POST, PUT)")
	CreateCmd.Flags().StringVarP(&monitor.Expected_http_code, "expected_http_code", "e", "20*", "Expected http code in the response of the request in the monitor (200, 201, 401...)")
	CreateCmd.Flags().StringVarP(&monitor.Body, "body", "b", "hello", "Body of the request in the monitor")
	CreateCmd.Flags().StringVarP(&monitor.Expected_contain, "expected_contain", "c", "", "Expected contain in the response of the request in the monitor")
	CreateCmd.Flags().IntVarP(&monitor.Timeout, "timeout", "t", 30, "Timeout of the request in the monitor")
	CreateCmd.Flags().StringVarP(&monitor.Username, "username", "s", "", "Username of the request in the monitor")
	CreateCmd.Flags().StringVarP(&monitor.Password, "password", "p", "", "Password of the request in the monitor")
	CreateCmd.Flags().BoolVarP(&monitor.CheckTls, "check_tls", "k", true, "Check tls of the request in the monitor")
	CreateCmd.Flags().StringVarP(&monitor.Level, "level", "l", "info", "Log level of the monitor (INFO or DEBUG)")
	CreateCmd.Flags().StringVarP(&rawHeaders, "headers", "H", "", "Headers of the request in the monitor (e.g., key1:value1,key2:value2)")
	CreateCmd.Flags().IntVarP(&monitor.User_id, "user_id", "i", 0, "User ID")

	err := CreateCmd.MarkFlagRequired("name")
	if nil != err {
		fmt.Println(err)
	}

	err = CreateCmd.MarkFlagRequired("url")
	if nil != err {
		fmt.Println(err)
	}

	err = CreateCmd.MarkFlagRequired("user_id")
	if nil != err {
		fmt.Println(err)
	}
}

// ? Helper function to parse headers string into []Header
func parseHeaders(raw string) ([]adminClient.Header, error) {
	var headers []adminClient.Header
	pairs := strings.Split(raw, ",")
	for _, pair := range pairs {
		kv := strings.SplitN(pair, ":", 2)
		if len(kv) != 2 {
			return nil, fmt.Errorf("header %q is not in key:value format", pair)
		}
		headers = append(headers, adminClient.Header{
			Name:  strings.TrimSpace(kv[0]),
			Value: strings.TrimSpace(kv[1]),
		})
	}
	return headers, nil
}
