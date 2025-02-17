package client

import (
	"bytes"
	"encoding/json"
)

func (c *Client) SendPrompt(adapter string, message string) (*PromptResponse, error) {
	buf := bytes.Buffer{}
	prompt := Prompt{
		Adapter: adapter,
		Message: message,
	}

	err := json.NewEncoder(&buf).Encode(prompt)
	if nil != err {
		return nil, err
	}
	resp_body, err := c.httpRequest("/ai/prompt", "POST", buf)
	if nil != err {
		return nil, err
	}

	response := &PromptResponse{}
	err = json.NewDecoder(resp_body).Decode(response)
	if nil != err {
		return nil, err
	}

	return response, nil
}
