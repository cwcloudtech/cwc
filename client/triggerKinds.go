package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func GetTriggerKinds() (*TriggerKindsResponse, error) {
	c, _ := NewClient()
	body, err := c.httpRequest("/faas/trigger_kinds", "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	triggerKinds := &TriggerKindsResponse{}
	err = json.NewDecoder(body).Decode(triggerKinds)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return triggerKinds, nil
	
}