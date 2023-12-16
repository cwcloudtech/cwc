package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func GetLanguages() (*LanguagesResponse, error) {
	c, _ := NewClient()
	body, err := c.httpRequest("/faas/languages", "GET", bytes.Buffer{})
	if nil != err {
		return nil, err
	}
	languages := &LanguagesResponse{}
	err = json.NewDecoder(body).Decode(languages)
	if nil != err {
		fmt.Println(err.Error())
		return nil, err
	}
	return languages, nil
}
