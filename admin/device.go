package admin

import (
	"bytes"
	"encoding/json"
)

func (c *Client) GetAllDevices() (*[]Device, error) {
	body, err := c.httpRequest("/admin/iot/devices", "GET", bytes.Buffer{})
	if nil != err {
		return nil, err
	}
	response := []Device{}
	err = json.NewDecoder(body).Decode(&response)
	if nil != err {
		return nil, err
	}
	return &response, nil
}

func (c *Client) DeleteDeviceById(deviceId string) error {
	_, err := c.httpRequest("/admin/iot/device/"+deviceId, "DELETE", bytes.Buffer{})
	if nil != err {
		return err
	}
	return nil
}
