package admin

import (
	"bytes"
	"encoding/json"
)

func (c *Client) GetAllClusters() (*[]Cluster, error) {
	body, err := c.httpRequest("/admin/kubernetes/cluster", "GET", bytes.Buffer{})
	if nil != err {
		return nil, err
	}

	clusters := []Cluster{}
	err = json.NewDecoder(body).Decode(&clusters)
	if nil != err {
		return nil, err
	}

	return &clusters, nil
}

func (c *Client) DeleteCluster(clusterId string) error {
	_, err := c.httpRequest("/admin/kubernetes/cluster/"+clusterId, "DELETE", bytes.Buffer{})
	if nil != err {
		return err
	}

	return nil
}
