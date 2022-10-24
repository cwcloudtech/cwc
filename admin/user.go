package admin

import (
	"bytes"
	"encoding/json"
	"fmt"
)
func (c *Client) GetAllUsers() (*ResponseUsers, error) {
	body, err := c.httpRequest(fmt.Sprintf("/admin/user/all"), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	ResponseUsers :=ResponseUsers{}
	err = json.NewDecoder(body).Decode(&ResponseUsers)

	if err != nil {
		return nil, err
	}
	return &ResponseUsers, nil
}

func (c *Client) GetUser(id string) (*ResponseUser, error) {
	body, err := c.httpRequest(fmt.Sprintf("/admin/user/%s",id), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	ResponseUser :=ResponseUser{}
	err = json.NewDecoder(body).Decode(&ResponseUser)

	if err != nil {
		return nil, err
	}
	return &ResponseUser, nil
}



