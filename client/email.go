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
	if nil != err {
		return nil, err
	}
	resp_body, err := c.httpRequest("/email", "POST", buf)
	if nil != err {
		return nil, err
	}

	response := &EmailResponse{}
	err = json.NewDecoder(resp_body).Decode(response)
	if nil != err {
		return nil, err
	}

	return response, nil
}
