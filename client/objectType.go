package client

import (
	"bytes"
	"encoding/json"
)

func (c *Client) CreateObjectType(objectType ObjectType) (*ObjectType, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(objectType)
	if nil != err {
		return nil, err
	}
	resp_body, err := c.httpRequest("/iot/object-type", "POST", buf)
	if nil != err {
		return nil, err
	}
	created_objectType := &ObjectType{}
	err = json.NewDecoder(resp_body).Decode(created_objectType)
	if nil != err {
		return nil, err
	}
	return created_objectType, nil
}

func (c *Client) DeleteObjectTypeById(objectTypeId string) error {
	_, err := c.httpRequest("/iot/object-type/"+objectTypeId, "DELETE", bytes.Buffer{})
	if nil != err {
		return err
	}
	return nil
}
