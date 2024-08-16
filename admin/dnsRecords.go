package admin

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (c *Client) AdminAddDnsRecord(recordName string, dnsZone string, dnsType string, ttl int, data string) (*DnsRecord, error) {
	buf := bytes.Buffer{}
	dnsRecord := DnsRecordCreate {
		RecordName: recordName,
		DnsZone: dnsZone,
		Type: dnsType,
		Ttl: ttl,
		Data: data,
	}

	err := json.NewEncoder(&buf).Encode(dnsRecord)
	if nil != err {
		return nil, err
	}

	resp_body, err := c.httpRequest(fmt.Sprintf("/admin/dns/%s/create", c.provider), "POST", buf)
	if nil != err {
		return nil, err
	}

	created_record := &DnsRecord{}

	err = json.NewDecoder(resp_body).Decode(created_record)
	if nil != err {
		return nil, err
	}

	return created_record, nil
}

func (c *Client) GetAllDnsRecords() (*[]DnsRecord, error) {
	body, err := c.httpRequest(fmt.Sprintf("/admin/dns/%s/list", c.provider), "GET", bytes.Buffer{})
	if nil != err {
		return nil, err
	}

	dnsRecords := []DnsRecord{}

	err = json.NewDecoder(body).Decode(&dnsRecords)
	if nil != err {
		return nil, err
	}

	return &dnsRecords, nil
}

func (c *Client) DeleteDnsRecord(recordId string, recordName string, dnsZone string) error {
	buf := bytes.Buffer{}
	dnsRecord := DnsRecordDelete {
		Id: recordId,
		RecordName: recordName,
		DnsZone: dnsZone,
	}
	
	err := json.NewEncoder(&buf).Encode(dnsRecord)
	if nil != err {
		return err
	}

	_, err = c.httpRequest(fmt.Sprintf("/admin/dns/%s/delete", c.provider), "PATCH", buf)
	if nil != err {
		return err
	}

	return nil
}
