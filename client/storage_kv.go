package client

import (
	"bytes"
	"cwc/utils"
	"encoding/json"
	"fmt"
	"net/url"
)

func (c *Client) CreateStorageKV(storageKV StorageKVCreateRequest) (*StorageKVResponse, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(storageKV)
	if nil != err {
		return nil, err
	}

	resp_body, err := c.httpRequest("/storage/kv", "POST", buf)
	if nil != err {
		return nil, err
	}

	response := &StorageKVResponse{}
	err = json.NewDecoder(resp_body).Decode(response)
	if nil != err {
		return nil, err
	}

	return response, nil
}

func (c *Client) GetStorageKV(key string) (*StorageKVResponse, error) {
	resp_body, err := c.httpRequest(fmt.Sprintf("/storage/kv/%s", key), "GET", bytes.Buffer{})
	if nil != err {
		return nil, err
	}

	response := &StorageKVResponse{}
	err = json.NewDecoder(resp_body).Decode(response)
	if nil != err {
		return nil, err
	}

	return response, nil
}

func (c *Client) ListStorageKVs(search string, startIndex int, maxResults int) (*StorageKVListResponse, error) {
	query := url.Values{}
	if utils.IsNotBlank(search) {
		query.Add("search", search)
	}

	query.Add("start_index", fmt.Sprintf("%d", startIndex))
	query.Add("max_results", fmt.Sprintf("%d", maxResults))

	path := fmt.Sprintf("/storage/kv?%s", query.Encode())
	resp_body, err := c.httpRequest(path, "GET", bytes.Buffer{})
	if nil != err {
		return nil, err
	}

	response := &StorageKVListResponse{}
	err = json.NewDecoder(resp_body).Decode(response)
	if nil != err {
		return nil, err
	}

	return response, nil
}

func (c *Client) UpdateStorageKV(key string, updateRequest StorageKVUpdateRequest) (*StorageKVResponse, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(updateRequest)
	if nil != err {
		return nil, err
	}

	resp_body, err := c.httpRequest(fmt.Sprintf("/storage/kv/%s", key), "PUT", buf)
	if nil != err {
		return nil, err
	}

	response := &StorageKVResponse{}
	err = json.NewDecoder(resp_body).Decode(response)
	if nil != err {
		return nil, err
	}

	return response, nil
}

func (c *Client) DeleteStorageKV(key string) (*StorageKVResponse, error) {
	resp_body, err := c.httpRequest(fmt.Sprintf("/storage/kv/%s", key), "DELETE", bytes.Buffer{})
	if nil != err {
		return nil, err
	}

	response := &StorageKVResponse{}
	err = json.NewDecoder(resp_body).Decode(response)
	if nil != err {
		return nil, err
	}

	return response, nil
}
