package update

import (
	"cwc/client"
	"cwc/handlers/user"
	"cwc/utils"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var (
	monitorId    string
	monitor      client.Monitor
	rawHeaders   string
	rawCallbacks string
)

var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a particular monitor",
	Long: `This command lets you update a particular monitor.
To use this command you have to provide the monitor ID`,
	Run: func(cmd *cobra.Command, args []string) {
		if rawHeaders != "" {
			headers, err := parseHeaders(rawHeaders)
			utils.ExitIfError(err)
			monitor.Headers = headers
		} else {
			monitor.Headers = []client.Header{}
		}

		if rawCallbacks != "" {
			callbacks, err := parseCallbacks(rawCallbacks)
			utils.ExitIfError(err)
			monitor.Callbacks = callbacks
		} else {
			monitor.Callbacks = []client.CallbacksContent{}
		}

		user.HandleUpdateMonitor(&monitorId, &monitor)
	},
}

func init() {
	UpdateCmd.Flags().StringVarP(&monitorId, "id", "i", "", "The monitor ID")
	UpdateCmd.Flags().StringVarP(&monitor.Type, "type", "y", "http", "Type of the monitor (http, tcp, icmp)")
	UpdateCmd.Flags().StringVarP(&monitor.Name, "name", "n", "", "Name of the monitor")
	UpdateCmd.Flags().StringVarP(&monitor.Family, "family", "f", "", "Family of the monitor")
	UpdateCmd.Flags().StringVarP(&monitor.Url, "url", "u", "", "Url of the monitor")
	UpdateCmd.Flags().StringVarP(&monitor.Method, "method", "m", "GET", "Method of the request in the monitor (GET, POST, PUT)")
	UpdateCmd.Flags().StringVarP(&monitor.Expected_http_code, "expected_http_code", "e", "20*", "Expected http code in the response of the request in the monitor (200, 201, 401...)")
	UpdateCmd.Flags().StringVarP(&monitor.Body, "body", "b", "", "Body of the request in the monitor")
	UpdateCmd.Flags().StringVarP(&monitor.Expected_contain, "expected_contain", "c", "", "Expected contain in the response of the request in the monitor")
	UpdateCmd.Flags().IntVarP(&monitor.Timeout, "timeout", "t", 30, "Timeout of the request in the monitor")
	UpdateCmd.Flags().StringVarP(&monitor.Username, "username", "s", "", "Username of the request in the monitor")
	UpdateCmd.Flags().StringVarP(&monitor.Password, "password", "p", "", "Password of the request in the monitor")
	UpdateCmd.Flags().BoolVarP(&monitor.CheckTls, "check_tls", "k", true, "Check tls of the request in the monitor")
	UpdateCmd.Flags().StringVarP(&monitor.Level, "level", "l", "info", "Log level of the monitor (INFO or DEBUG)")
	UpdateCmd.Flags().StringVarP(&rawHeaders, "headers", "H", "", "Headers of the request in the monitor (e.g., key1:value1,key2:value2)")
	UpdateCmd.Flags().StringVarP(&rawCallbacks, "callbacks", "C", "", "Callbacks for the monitor (format: type:http,endpoint:https://example.com,token:123;type:mqtt,endpoint:mqtt://broker.com,topic:test)")

	err := UpdateCmd.MarkFlagRequired("id")
	if nil != err {
		fmt.Println(err)
	}
}

// ? Helper function to parse headers string into []Header
func parseHeaders(raw string) ([]client.Header, error) {
	var headers []client.Header
	pairs := strings.Split(raw, ",")
	for _, pair := range pairs {
		kv := strings.SplitN(pair, ":", 2)
		if len(kv) != 2 {
			return nil, fmt.Errorf("header %q is not in key:value format", pair)
		}
		headers = append(headers, client.Header{
			Name:  strings.TrimSpace(kv[0]),
			Value: strings.TrimSpace(kv[1]),
		})
	}
	return headers, nil
}

// ? Helper function to parse callbacks string into []CallbacksContent
func parseCallbacks(raw string) ([]client.CallbacksContent, error) {
	if raw == "" {
		return []client.CallbacksContent{}, nil
	}

	var callbacks []client.CallbacksContent
	callbackStrings := strings.Split(raw, ";")

	for _, callback := range callbackStrings {
		props := strings.Split(callback, ",")
		var cb client.CallbacksContent

		for _, prop := range props {
			kv := strings.SplitN(prop, ":", 2)
			if len(kv) != 2 {
				return nil, fmt.Errorf("callback property %q is not in key:value format", prop)
			}
			key := strings.TrimSpace(kv[0])
			value := strings.TrimSpace(kv[1])

			switch key {
			case "type":
				cb.Type = value
			case "endpoint":
				cb.Endpoint = value
			case "token":
				cb.Token = value
			case "client_id":
				cb.Client_id = value
			case "user_data":
				cb.User_data = value
			case "username":
				cb.Username = value
			case "password":
				cb.Password = value
			case "port":
				cb.Port = value
			case "subscription":
				cb.Subscription = value
			case "qos":
				cb.Qos = value
			case "topic":
				cb.Topic = value
			default:
				return nil, fmt.Errorf("unknown callback property: %s", key)
			}
		}

		// Validate required fields
		if cb.Type == "" {
			return nil, fmt.Errorf("callback type is required")
		}
		if cb.Endpoint == "" {
			return nil, fmt.Errorf("callback endpoint is required")
		}

		callbacks = append(callbacks, cb)
	}

	return callbacks, nil
}
