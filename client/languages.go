package client 

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func GetLanguages() (*LanguagesResponse, error) {
	c, _ := NewClient()
	body, err := c.httpRequest("/faas/languages", "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	languages := &LanguagesResponse{}
	err = json.NewDecoder(body).Decode(languages)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return languages, nil
}