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

func (c *Client) CreateDevice(device Device) (*Device, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(device)
	if nil != err {
		return nil, err
	}
	body, err := c.httpRequest("/iot/device", "POST", buf)
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

func (c *Client) DeleteDeviceById(deviceId string) error {
	_, err := c.httpRequest("/iot/device/"+deviceId, "DELETE", bytes.Buffer{})
	if nil != err {
		return err
	}
	return nil
}
