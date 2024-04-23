package admin

import (
	"bytes"
	"encoding/json"
)

func (c *Client) GetAllNumericData() (*[]NumericData, error) {
	body, err := c.httpRequest("/admin/iot/data/numeric", "GET", bytes.Buffer{})
	if nil != err {
		return nil, err
	}
	response := []NumericData{}
	err = json.NewDecoder(body).Decode(&response)
	if nil != err {
		return nil, err
	}
	return &response, nil
}

func (c *Client) GetAllStringData() (*[]StringData, error) {
	body, err := c.httpRequest("/admin/iot/data/string", "GET", bytes.Buffer{})
	if nil != err {
		return nil, err
	}
	response := []StringData{}
	err = json.NewDecoder(body).Decode(&response)
	if nil != err {
		return nil, err
	}
	return &response, nil
}

