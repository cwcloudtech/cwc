package admin

import (
	"bytes"
	"cwc/utils"
	"encoding/json"
	"fmt"
)

func (c *Client) AdminAddBucket(user_email string, name string, reg_type string) (*Bucket, error) {
	buf := bytes.Buffer{}
	bucket := Bucket{
		Name:  name,
		Type:  reg_type,
		Email: user_email,
	}

	err := json.NewEncoder(&buf).Encode(bucket)
	if nil != err {
		return nil, err
	}

	respBody, err := c.httpRequest(fmt.Sprintf("/admin/bucket/%s/%s/provision", c.provider, c.region), "POST", buf)
	if nil != err {
		return nil, err
	}

	created_bucket := &Bucket{}

	err = json.NewDecoder(respBody).Decode(created_bucket)
	if nil != err {
		return nil, err
	}

	return created_bucket, nil
}

func (c *Client) GetAllBuckets() (*[]Bucket, error) {
	body, err := c.httpRequest(fmt.Sprintf("/admin/bucket/%s/%s/all", c.provider, c.region), "GET", bytes.Buffer{})
	if nil != err {
		return nil, err
	}

	buckets := []Bucket{}

	err = json.NewDecoder(body).Decode(&buckets)
	if nil != err {
		return nil, err
	}

	return &buckets, nil
}

func (c *Client) GetBucket(bucket_id string) (*Bucket, error) {
	body, err := c.httpRequest(fmt.Sprintf("/admin/bucket/%s", bucket_id), "GET", bytes.Buffer{})
	if nil != err {
		return nil, err
	}

	bucket := &Bucket{}
	err = json.NewDecoder(body).Decode(bucket)
	if nil != err {
		return nil, err
	}

	return bucket, nil
}

func (c *Client) UpdateBucket(id string, email string) error {
	buf := bytes.Buffer{}
	var updateCreds bool = true

	renew := RenewCredentials{
		Email:       email,
		UpdateCreds: updateCreds,
	}

	if utils.IsValidEmail(email) {
		updateCreds = false
	}

	encode_err := json.NewEncoder(&buf).Encode(renew)
	if nil != encode_err {
		return encode_err
	}

	_, err := c.httpRequest(fmt.Sprintf("/admin/bucket/%s", id), "PATCH", buf)
	if nil != err {
		return err
	}

	return nil
}

func (c *Client) DeleteBucket(bucketId string) error {
	_, err := c.httpRequest(fmt.Sprintf("/admin/bucket/%s", bucketId), "DELETE", bytes.Buffer{})
	if nil != err {
		return err
	}

	return nil
}
