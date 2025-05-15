package admin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
)

func (c *Client) GetAllStorageKVs(search string, startIndex int, maxResults int, userId string) (*StorageKVListResponse, error) {
	query := url.Values{}
	if search != "" {
		query.Add("search", search)
	}
	query.Add("start_index", fmt.Sprintf("%d", startIndex))
	query.Add("max_results", fmt.Sprintf("%d", maxResults))

	path := fmt.Sprintf("/admin/storage/kv/all?%s", query.Encode())
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

func (c *Client) GetUserStorageKVs(userId string, search string, startIndex int, maxResults int) (*StorageKVListResponse, error) {
	query := url.Values{}
	if search != "" {
		query.Add("search", search)
	}
	query.Add("start_index", fmt.Sprintf("%d", startIndex))
	query.Add("max_results", fmt.Sprintf("%d", maxResults))

	path := fmt.Sprintf("/admin/storage/kv/user/%s?%s", userId, query.Encode())
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

func (c *Client) DeleteUserStorageKV(userId string, key string) (*StorageKVResponse, error) {
	path := fmt.Sprintf("/admin/storage/kv/user/%s/storage/%s", userId, key)
	resp_body, err := c.httpRequest(path, "DELETE", bytes.Buffer{})
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
