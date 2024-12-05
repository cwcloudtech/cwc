package admin

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (c *Client) GetAllMonitors() (*[]Monitor, error) {
	body, err := c.httpRequest("/admin/monitor/all", "GET", bytes.Buffer{})
	if nil != err {
		return nil, err
	}
	response := []Monitor{}
	err = json.NewDecoder(body).Decode(&response)

	if nil != err {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetMonitorById(monitor_id string) (*Monitor, error) {
	body, err := c.httpRequest(fmt.Sprintf("/admin/monitor/%s", monitor_id), "GET", bytes.Buffer{})
	if nil != err {
		return nil, err
	}
	monitor := &Monitor{}
	err = json.NewDecoder(body).Decode(monitor)
	if nil != err {
		return nil, err
	}
	return monitor, nil
}

func (c *Client) AddMonitor(monitor Monitor) (*Monitor, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(monitor)
	if nil != err {
		return nil, err
	}
	resp_body, err := c.httpRequest("/admin/monitor", "POST", buf)
	if nil != err {
		return nil, err
	}
	created_monitor := &Monitor{}
	err = json.NewDecoder(resp_body).Decode(created_monitor)
	if nil != err {
		return nil, err
	}
	return created_monitor, nil
}

func (c *Client) UpdateMonitorById(monitorId string, monitor Monitor) (*Monitor, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(monitor)
	if nil != err {
		return nil, err
	}
	resp_body, err := c.httpRequest(fmt.Sprintf("/admin/monitor/%s", monitorId), "PUT", buf)
	if nil != err {
		return nil, err
	}
	updated_monitor := &Monitor{}
	err = json.NewDecoder(resp_body).Decode(updated_monitor)
	if nil != err {
		return nil, err
	}
	return updated_monitor, nil
}

func (c *Client) DeleteMonitorById(monitorId string) error {
	_, err := c.httpRequest(fmt.Sprintf("/admin/monitor/%s", monitorId), "DELETE", bytes.Buffer{})
	if nil != err {
		return err
	}
	return nil
}
