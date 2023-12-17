package admin

import (
	"bytes"
	"encoding/json"
)

func (c *Client) AdminSendEmail(from string, to string, bcc string, subject string, content string, templated bool) (*EmailResponse, error) {
	buf := bytes.Buffer{}
	email := Email{
		From:      from,
		To:        to,
		Bcc:       bcc,
		Subject:   subject,
		Content:   content,
		Templated: templated,
	}

	err := json.NewEncoder(&buf).Encode(email)
	if nil != err {
		return nil, err
	}

	resp_body, err := c.httpRequest("/admin/email", "POST", buf)
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
