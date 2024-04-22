package client 

import (
	"bytes"
	"encoding/json"
)

func (c *Client) GetAllDevices() (*[]Device, error) {
	body, err := c.httpRequest("/iot/devices", "GET", bytes.Buffer{})
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

func (c *Client) GetDeviceById(deviceId string) (*Device, error) {
	body, err := c.httpRequest("/iot/devices/"+deviceId, "GET", bytes.Buffer{})
	if nil != err {
		return nil, err
	}
	response := Device{}
	err = json.NewDecoder(body).Decode(&response)
	if nil != err {
		return nil, err
	}
	return &response, nil
}
