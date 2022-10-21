package admin

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (c *Client) GetAllBuckets() (*[]Bucket, error) {
	body, err := c.httpRequest(fmt.Sprintf("/admin/bucket/%s/%s/all", c.provider, c.region), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	buckets := []Bucket{}
	err = json.NewDecoder(body).Decode(&buckets)

	if err != nil {
		return nil, err
	}
	return &buckets, nil
}

func (c *Client) UpdateBucket(id string) error {
	buf := bytes.Buffer{}

	_, err := c.httpRequest(fmt.Sprintf("/admin/bucket/%s", id), "PATCH", buf)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) DeleteBucket(bucketId string) error {
	_, err := c.httpRequest(fmt.Sprintf("/admin/bucket/%s", bucketId), "DELETE", bytes.Buffer{})
	if err != nil {
		return err
	}
	return nil
}
