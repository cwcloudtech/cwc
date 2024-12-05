package update

import (
	adminClient "cwc/admin"
	"cwc/handlers/admin"
	"cwc/utils"
	"strings"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	monitorId string
	monitor    adminClient.Monitor
	rawHeaders string
)

var UpdateCmd = &cobra.Command{
	Use: "update",
	Short: "Update a monitor in the cloud",
	Long: "This command lets you update a monitor in the cloud.",
	Run: func(cmd *cobra.Command, args []string) {
		if rawHeaders != "" {
			headers, err := parseHeaders(rawHeaders)
			utils.ExitIfError(err)
			monitor.Headers = headers
		} else {
			monitor.Headers = []adminClient.Header{}
		}
		admin.HandleUpdateMonitor(&monitorId, &monitor)
	},
}

func init() {
	UpdateCmd.Flags().StringVarP(&monitorId, "id", "m", "", "The monitor ID")
	UpdateCmd.Flags().StringVarP(&monitor.Type, "type", "y", "http", "Type of the monitor (http, tcp, icmp)")
	UpdateCmd.Flags().StringVarP(&monitor.Name, "name", "n", "", "Name of the monitor")
	UpdateCmd.Flags().StringVarP(&monitor.Family, "family", "f", "", "Family of the monitor")
	UpdateCmd.Flags().StringVarP(&monitor.Url, "url", "u", "", "Url of the monitor")
	UpdateCmd.Flags().StringVarP(&monitor.Method, "method", "M", "GET", "Method of the request in the monitor (GET, POST, PUT)")
	UpdateCmd.Flags().StringVarP(&monitor.Expected_http_code, "expected_http_code", "e", "20*", "Expected http code in the response of the request in the monitor (200, 201, 401...)")
	UpdateCmd.Flags().StringVarP(&monitor.Body, "body", "b", "", "Body of the request in the monitor")
	UpdateCmd.Flags().StringVarP(&monitor.Expected_contain, "expected_contain", "c", "", "Expected contain in the response of the request in the monitor")
	UpdateCmd.Flags().IntVarP(&monitor.Timeout, "timeout", "t", 30, "Timeout of the request in the monitor")
	UpdateCmd.Flags().StringVarP(&monitor.Username, "username", "s", "", "Username of the request in the monitor")
	UpdateCmd.Flags().StringVarP(&monitor.Password, "password", "p", "", "Password of the request in the monitor")
	UpdateCmd.Flags().StringVarP(&rawHeaders, "headers", "H", "", "Headers of the request in the monitor (e.g., key1:value1,key2:value2)")
	UpdateCmd.Flags().IntVarP(&monitor.User_id, "user_id", "I", 0, "User ID")

	err := UpdateCmd.MarkFlagRequired("id")
	if nil != err {
		fmt.Println(err)
	}
}


//? Helper function to parse headers string into []Header
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
