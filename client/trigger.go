package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func GetKinds() (*TriggerKindsResponse, error) {
	c, _ := NewClient()
	body, err := c.httpRequest("/faas/trigger_kinds", "GET", bytes.Buffer{})
	if nil != err {
		return nil, err
	}
	triggerKinds := &TriggerKindsResponse{}
	err = json.NewDecoder(body).Decode(triggerKinds)
	if nil != err {
		fmt.Println(err.Error())
		return nil, err
	}
	return triggerKinds, nil
}

func GetFunctionByIdArgs(function_id string) ([]string, error) {
	c, _ := NewClient()
	body, err := c.httpRequest(fmt.Sprintf("/faas/function/%s", function_id), "GET", bytes.Buffer{})
	if nil != err {
		return nil, err
	}
	function := &Function{}
	err = json.NewDecoder(body).Decode(function)
	if nil != err {
		return nil, err
	}
	return function.Content.Args, nil
}

func (c *Client) GetAllTriggers() (*[]Trigger, error) {
	body, err := c.httpRequest(fmt.Sprintf("/faas/triggers"), "GET", bytes.Buffer{})
	if nil != err {
		return nil, err
	}
	response := TriggersResponse{}
	err = json.NewDecoder(body).Decode(&response)

	if nil != err {
		return nil, err
	}
	return &response.Results, nil
}

func (c *Client) GetTriggerById(trigger_id string) (*Trigger, error) {
	body, err := c.httpRequest(fmt.Sprintf("/faas/trigger/%s", trigger_id), "GET", bytes.Buffer{})
	if nil != err {
		return nil, err
	}
	trigger := &Trigger{}
	err = json.NewDecoder(body).Decode(trigger)
	if nil != err {
		return nil, err
	}
	return trigger, nil
}

func (c *Client) AddTrigger(trigger Trigger) (*Trigger, error) {
	buf := bytes.Buffer{}
	if len(trigger.Content.Args) == 0 {
		trigger.Content.Args = []Argument{}
	}
	added_trigger := &AddTriggerBody{
		Kind:    trigger.Kind,
		Content: trigger.Content,
	}
	err := json.NewEncoder(&buf).Encode(added_trigger)
	if nil != err {
		return nil, err
	}
	respBody, err := c.httpRequest(fmt.Sprintf("/faas/trigger"), "POST", buf)
	if nil != err {
		return nil, err
	}
	created_trigger := &Trigger{}
	err = json.NewDecoder(respBody).Decode(created_trigger)
	if nil != err {
		return nil, err
	}
	return created_trigger, nil
}

func (c *Client) DeleteTriggerById(triggerId string) error {
	_, err := c.httpRequest(fmt.Sprintf("/faas/trigger/%s", triggerId), "DELETE", bytes.Buffer{})
	if nil != err {
		return err
	}
	return nil
}

func (c *Client) TruncateTriggers() error {
	_, err := c.httpRequest(fmt.Sprintf("/faas/triggers"), "DELETE", bytes.Buffer{})
	if nil != err {
		return err
	}
	return nil
}
