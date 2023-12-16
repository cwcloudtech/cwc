package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func GetTriggerKinds() (*TriggerKindsResponse, error) {
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
