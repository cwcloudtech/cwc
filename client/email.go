package client

import (
	"bytes"
	"encoding/json"
)

func (c *Client) SendEmail(from string, to string, bcc string, subject string, content string) (*EmailResponse, error) {
	buf := bytes.Buffer{}
	email := Email{
		From:    from,
		To:      to,
		Bcc:     bcc,
		Subject: subject,
		Content: content,
	}

	err := json.NewEncoder(&buf).Encode(email)
	if err != nil {
		return nil, err
	}
	respBody, err := c.httpRequest("/email", "POST", buf)
	if err != nil {
		return nil, err
	}

	response := &EmailResponse{}
	err = json.NewDecoder(respBody).Decode(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
