package admin

import (
	"bytes"
	"encoding/json"
)

func (c *Client) AdminSendEmail(from_email string, to_email string, subject string, content string, templated bool) (*EmailResponse, error) {
	buf := bytes.Buffer{}
	email := Email{
		From:      from_email,
		To:        to_email,
		Subject:   subject,
		Content:   content,
		Templated: templated,
	}

	err := json.NewEncoder(&buf).Encode(email)
	if err != nil {
		return nil, err
	}
	respBody, err := c.httpRequest("/admin/email", "POST", buf)
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
