package client

import (
	"bytes"
	"encoding/json"
)

func (c *Client) SendPrompt(model string, message string) (*PromptResponse, error) {
	buf := bytes.Buffer{}
	prompt := Prompt{
		Model:   model,
		Message: message,
	}

	err := json.NewEncoder(&buf).Encode(prompt)
	if nil != err {
		return nil, err
	}
	respBody, err := c.httpRequest("/ai/prompt", "POST", buf)
	if nil != err {
		return nil, err
	}

	response := &PromptResponse{}
	err = json.NewDecoder(respBody).Decode(response)
	if nil != err {
		return nil, err
	}

	return response, nil
}
