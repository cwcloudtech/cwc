package user

import (
	"cwc/client"
	"cwc/config"
	"cwc/utils"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

func HandleGetMonitors(monitors *[]client.Monitor, pretty *bool) {
	if config.IsPrettyFormatExpected(pretty) {
		displayMonitorsAsTable(*monitors)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(monitors)
	} else {
		var monitorsDisplay []client.Monitor
		for i, monitor := range *monitors {
			monitorsDisplay = append(monitorsDisplay, client.Monitor{
				Id:            monitor.Id,
				Name:          monitor.Name,
				Family:        monitor.Family,
				Url:           monitor.Url,
				Method:        monitor.Method,
				Timeout:       monitor.Timeout,
				Updated_at:    monitor.Updated_at,
				Status:        monitor.Status,
				Response_time: monitor.Response_time,
			})
			monitorsDisplay[i].Id = monitor.Id
		}
		utils.PrintMultiRow(client.Monitor{}, monitorsDisplay)
	}
}

func HandleGetMonitor(monitor *client.Monitor, pretty *bool) {
	var monitorDisplay client.Monitor
	monitorDisplay.Id = monitor.Id
	monitorDisplay.Type = monitor.Type
	monitorDisplay.Name = monitor.Name
	monitorDisplay.Family = monitor.Family
	monitorDisplay.Url = monitor.Url
	monitorDisplay.Method = monitor.Method
	monitorDisplay.Expected_http_code = monitor.Expected_http_code
	if monitor.Method == "POST" || monitor.Method == "PUT" {
		monitorDisplay.Body = monitor.Body
	}
	monitorDisplay.Timeout = monitor.Timeout
	monitorDisplay.Username = monitor.Username
	monitorDisplay.Password = monitor.Password
	monitorDisplay.Headers = monitor.Headers
	monitorDisplay.Status = monitor.Status
	monitorDisplay.CheckTls = monitor.CheckTls
	monitorDisplay.Level = monitor.Level
	monitorDisplay.Response_time = monitor.Response_time
	monitorDisplay.Updated_at = monitor.Updated_at

	if config.IsPrettyFormatExpected(pretty) {
		utils.PrintPretty("Found monitor", monitorDisplay)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(monitor)
	} else {
		utils.PrintRow(monitorDisplay)
	}
}

func checkMonitorConfig(monitor *client.Monitor) error {
	if monitor.Type != "http" && monitor.Type != "tcp" {
		return fmt.Errorf("invalid monitor type. Monitor type must be either http or tcp")
	}

	if monitor.Method != "GET" && monitor.Method != "POST" && monitor.Method != "PUT" {
		return fmt.Errorf("invalid method. Method must be either GET, POST or PUT")
	}

	if monitor.Level != "info" && monitor.Level != "debug" {
		return fmt.Errorf("invalid log level. Log level must be either info or debug")
	}

	return nil
}

func PrepareAddMonitor(monitor *client.Monitor) (client.Monitor, error) {
	c, err := client.NewClient()
	utils.ExitIfError(err)

	checkMonitorConfig(monitor)

	created_monitor, err := c.AddMonitor(*monitor)
	utils.ExitIfError(err)
	return *created_monitor, err
}

func HandleAddMonitor(createdMonitor *client.Monitor, pretty *bool) {
	if config.IsPrettyFormatExpected(pretty) {
		utils.PrintPretty("Monitor successfully created", *createdMonitor)
	} else if config.GetDefaultFormat() == "json" {
		utils.PrintJson(createdMonitor)
	} else {
		utils.PrintRow(*createdMonitor)
	}
}

func HandleUpdateMonitor(monitorId *string, updatedMonitor *client.Monitor) {
	c, err := client.NewClient()
	utils.ExitIfError(err)

	checkMonitorConfig(updatedMonitor)

	monitor, err := c.GetMonitorById(*monitorId)
	utils.ExitIfError(err)

	if utils.IsNotBlank(updatedMonitor.Type) {
		monitor.Type = updatedMonitor.Type
	}

	if utils.IsNotBlank(updatedMonitor.Name) {
		monitor.Name = updatedMonitor.Name
	} else {
		monitor.Name = utils.ShortName(monitor.Name, monitor.Hash)
	}

	if utils.IsNotBlank(updatedMonitor.Family) {
		monitor.Family = updatedMonitor.Family
	}

	if utils.IsNotBlank(updatedMonitor.Url) {
		monitor.Url = updatedMonitor.Url
	}

	if utils.IsNotBlank(updatedMonitor.Method) {
		monitor.Method = updatedMonitor.Method
	}

	if utils.IsNotBlank(updatedMonitor.Body) {
		monitor.Body = updatedMonitor.Body
	}

	if utils.IsNotBlank(updatedMonitor.Expected_contain) {
		monitor.Expected_contain = updatedMonitor.Expected_contain
	}

	if utils.IsNotBlank(updatedMonitor.Username) {
		monitor.Username = updatedMonitor.Username
	}

	if utils.IsNotBlank(updatedMonitor.Password) {
		monitor.Password = updatedMonitor.Password
	}

	monitor.CheckTls = updatedMonitor.CheckTls

	if utils.IsNotBlank(updatedMonitor.Level) {
		monitor.Level = updatedMonitor.Level
	}

	if len(updatedMonitor.Headers) > 0 {
		monitor.Headers = updatedMonitor.Headers
	}

	if len(updatedMonitor.Callbacks) > 0 {
		monitor.Callbacks = updatedMonitor.Callbacks
	}

	_, updateError := c.UpdateMonitorById(*monitorId, *monitor)
	utils.ExitIfError(updateError)

	fmt.Println("Monitor successfully updated")
}

func HandleDeleteMonitor(monitorId *string) {
	c, err := client.NewClient()
	utils.ExitIfError(err)

	err = c.DeleteMonitorById(*monitorId)
	utils.ExitIfError(err)

	fmt.Println("Monitor successfully deleted")
}

func displayMonitorsAsTable(monitors []client.Monitor) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Id", "Name", "Family", "Method", "Url", "Updated_at", "Status", "Response_time"})

	if len(monitors) == 0 {
		table.Append([]string{"No monitors available", "404", "404", "404", "404", "404", "404", "404"})
	} else {
		for _, monitor := range monitors {
			table.Append([]string{
				monitor.Id,
				monitor.Name,
				monitor.Family,
				monitor.Method,
				monitor.Url,
				monitor.Updated_at,
				monitor.Status,
				monitor.Response_time,
			})
		}
		table.Render()
	}
}
