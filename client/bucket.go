package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (c *Client) GetAllBuckets() (*[]Bucket, error) {
	body, err := c.httpRequest(fmt.Sprintf("/bucket/%s/%s", c.provider, c.region), "GET", bytes.Buffer{})
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

func (c *Client) GetBucket(bucket_id string) (*Bucket, error) {
	body, err := c.httpRequest(fmt.Sprintf("/bucket/%s/%s/%s", c.provider, c.region, bucket_id), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	bucket := &Bucket{}
	err = json.NewDecoder(body).Decode(bucket)
	if err != nil {
		return nil, err
	}
	return bucket, nil
}


func (c *Client) UpdateBucket(id string) error {
	buf := bytes.Buffer{}

	_, err := c.httpRequest(fmt.Sprintf("/bucket/%s/%s/%s", c.provider, c.region, id), "PATCH", buf)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) DeleteBucket(bucketId string) error {
	_, err := c.httpRequest(fmt.Sprintf("/bucket/%s/%s/%s", c.provider, c.region, bucketId), "DELETE", bytes.Buffer{})
	if err != nil {
		return err
	}
	return nil
}
