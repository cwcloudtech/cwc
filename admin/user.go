package admin

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (c *Client) GetAllUsers() (*ResponseUsers, error) {
	body, err := c.httpRequest(fmt.Sprintf("/admin/user/all"), "GET", bytes.Buffer{})
	if nil != err {
		return nil, err
	}
	ResponseUsers := ResponseUsers{}
	err = json.NewDecoder(body).Decode(&ResponseUsers)

	if nil != err {
		return nil, err
	}
	return &ResponseUsers, nil
}

func (c *Client) GetUser(id string) (*ResponseUser, error) {
	body, err := c.httpRequest(fmt.Sprintf("/admin/user/%s", id), "GET", bytes.Buffer{})
	if nil != err {
		return nil, err
	}
	ResponseUser := ResponseUser{}
	err = json.NewDecoder(body).Decode(&ResponseUser)

	if nil != err {
		return nil, err
	}
	return &ResponseUser, nil
}

func (c *Client) DeleteUser(userId string) error {
	_, err := c.httpRequest(fmt.Sprintf("/admin/user/%s", userId), "DELETE", bytes.Buffer{})
	if nil != err {
		return err
	}
	return nil
}
