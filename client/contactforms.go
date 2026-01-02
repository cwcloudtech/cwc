package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (c *Client) GetAllForms() (*[]ContactForm, error) {
	body, err := c.httpRequest("/contactform/all", "GET", bytes.Buffer{})
	if nil != err {
		return nil, err
	}
	response := []ContactForm{}
	err = json.NewDecoder(body).Decode(&response)

	if nil != err {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetFormById(form_id string) (*ContactForm, error) {
	body, err := c.httpRequest(fmt.Sprintf("/contactform/%s", form_id), "GET", bytes.Buffer{})
	if nil != err {
		return nil, err
	}

	form := &ContactForm{}
	err = json.NewDecoder(body).Decode(form)
	if nil != err {
		return nil, err
	}

	return form, nil
}

func (c *Client) AddForm(form ContactForm) (*ContactForm, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(form)
	if nil != err {
		return nil, err
	}

	resp_body, err := c.httpRequest("/contactform", "POST", buf)
	if nil != err {
		return nil, err
	}

	created_form := &ContactForm{}
	err = json.NewDecoder(resp_body).Decode(created_form)
	if nil != err {
		return nil, err
	}

	return created_form, nil
}

func (c *Client) UpdateFormById(formId string, form ContactForm) (*ContactForm, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(form)
	if nil != err {
		return nil, err
	}

	resp_body, err := c.httpRequest(fmt.Sprintf("/contactform/%s", formId), "PUT", buf)
	if nil != err {
		return nil, err
	}

	updated_form := &ContactForm{}
	err = json.NewDecoder(resp_body).Decode(updated_form)
	if nil != err {
		return nil, err
	}

	return updated_form, nil
}

func (c *Client) DeleteFormById(formId string) error {
	_, err := c.httpRequest(fmt.Sprintf("/contactform/%s", formId), "DELETE", bytes.Buffer{})
	if nil != err {
		return err
	}
	return nil
}
