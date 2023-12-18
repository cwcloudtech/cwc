package client

import (
	"bytes"
	"cwc/config"
	"cwc/utils"
	"encoding/json"
	"errors"
	"fmt"
)

func GetProviderRegions() (*ProviderRegions, error) {
	provider := config.GetDefaultProvider()
	if utils.IsBlank(provider) {
		return nil, errors.New("provider is not set")
	}

	c, _ := NewClient()
	body, err := c.httpRequest(fmt.Sprintf("/provider/%s/region", provider), "GET", bytes.Buffer{})
	if nil != err {
		return nil, err
	}

	providerRegions := &ProviderRegions{}
	err = json.NewDecoder(body).Decode(providerRegions)
	if nil != err {
		return nil, err
	}

	return providerRegions, nil
}

func GetProviders() (*AvailableProviders, error) {
	c, _ := NewClient()
	body, err := c.httpRequest("/provider", "GET", bytes.Buffer{})
	if nil != err {
		return nil, err
	}

	available_providers := &AvailableProviders{}
	err = json.NewDecoder(body).Decode(available_providers)
	if nil != err {
		fmt.Println(err.Error())
		return nil, err
	}

	return available_providers, nil
}
