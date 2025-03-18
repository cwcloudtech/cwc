package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (c *Client) SendPrompt(adapter string, message string, listId string) (*PromptResponse, error) {
	buf := bytes.Buffer{}
	prompt := Prompt{
		Adapter: adapter,
		Message: message,
		ListId: listId,
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

func (c *Client) GetPromptHistory(startIndex, maxResults int) (*PromptHistoryResponse, error) {
	//? Set default values if not provided
	if startIndex < 0 {
		startIndex = 0
	}
	if maxResults <= 0 {
		maxResults = 7
	}

	//? Construct URL with query parameters
	url := fmt.Sprintf("/ai/prompt/list/all?start_index=%d&max_results=%d", startIndex, maxResults)

	respBody, err := c.httpRequest(url, "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}

	response := &PromptHistoryResponse{}
	err = json.NewDecoder(respBody).Decode(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}


func (c *Client) GetPromptDetails(listId string) (*PromptDetailsResponse, error) {
	url := fmt.Sprintf("/ai/prompt/list/%s", listId)
	
	respBody, err := c.httpRequest(url, "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}

	response := &PromptDetailsResponse{}
	err = json.NewDecoder(respBody).Decode(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
