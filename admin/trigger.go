package admin

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (c *Client) GetAllTriggers() (*[]Trigger, error) {
	body, err := c.httpRequest(fmt.Sprintf("/admin/faas/triggers"), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	response := TriggersResponse{}
	err = json.NewDecoder(body).Decode(&response)

	if err != nil {
		return nil, err
	}
	return &response.Results, nil
}

func (c *Client) GetTriggerOwnerById(id string) (*TriggerOwner, error) {
	body, err := c.httpRequest(fmt.Sprintf("/admin/faas/trigger/%s/owner", id), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	response := TriggerOwner{}
	err = json.NewDecoder(body).Decode(&response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}