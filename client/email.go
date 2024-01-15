package client

import (
	"bytes"
	"cwc/utils"
	"encoding/json"
)

func (c *Client) SendEmail(from string, to string, bcc string, subject string, content string) (*EmailResponse, error) {
	buf := bytes.Buffer{}

	email := Email{
		To:      to,
		Subject: subject,
		Content: "<p>"+content+"</p>",
		Templated: false,
	}

	if !utils.IsBlank(from) {
		email.From = &from
	}

	if !utils.IsBlank(bcc) {
		email.Bcc = &bcc
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
