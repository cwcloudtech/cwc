package client

import (
	"bytes"
	"encoding/json"
)

func (c *Client) CreateData(data Data) (*Data, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(data)
	if nil != err {
		return nil, err
	}
	body, err := c.httpRequest("/iot/data", "POST", buf)
	if nil != err {
		return nil, err
	}
	response := Data{}
	err = json.NewDecoder(body).Decode(&response)
	if nil != err {
		return nil, err
	}
	return &response, nil
}
