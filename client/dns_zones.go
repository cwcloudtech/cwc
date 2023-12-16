package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func GetDnsZones() (*Dns_zones, error) {
	c, _ := NewClient()
	body, err := c.httpRequest(fmt.Sprintf("/dns_zones"), "GET", bytes.Buffer{})
	if nil != err {
		return nil, err
	}
	dns_zones := &Dns_zones{}
	err = json.NewDecoder(body).Decode(dns_zones)
	if nil != err {
		fmt.Println(err.Error())
		return nil, err
	}
	return dns_zones, nil
}
