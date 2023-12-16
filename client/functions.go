package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (c *Client) GetAllFunctions() (*[]Function, error) {
	body, err := c.httpRequest("/faas/functions", "GET", bytes.Buffer{})
	if nil != err {
		return nil, err
	}
	response := FunctionsResponse{}
	err = json.NewDecoder(body).Decode(&response)

	if nil != err {
		return nil, err
	}

	return &response.Results, nil
}

func (c *Client) GetFunctionById(function_id string) (*Function, error) {
	body, err := c.httpRequest(fmt.Sprintf("/faas/function/%s", function_id), "GET", bytes.Buffer{})
	if nil != err {
		return nil, err
	}
	function := &Function{}
	err = json.NewDecoder(body).Decode(function)
	if nil != err {
		return nil, err
	}
	return function, nil
}

func (c *Client) AddFunction(function Function) (*Function, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(function)
	if nil != err {
		return nil, err
	}
	respBody, err := c.httpRequest(fmt.Sprintf("/faas/function"), "POST", buf)
	if nil != err {
		return nil, err
	}
	created_function := &Function{}
	err = json.NewDecoder(respBody).Decode(created_function)
	if nil != err {
		return nil, err
	}
	return created_function, nil
}

func (c *Client) DeleteFunctionById(functionId string) error {
	_, err := c.httpRequest(fmt.Sprintf("/faas/function/%s", functionId), "DELETE", bytes.Buffer{})
	if nil != err {
		return err
	}
	return nil
}

func (c *Client) UpdateFunction(function Function) (*Function, error) {
	buf := bytes.Buffer{}
	updated_function := &UpdateFunctionBody{
		Id:        function.Id,
		Is_public: function.Is_public,
		Content:   function.Content,
	}
	err := json.NewEncoder(&buf).Encode(function)
	if nil != err {
		return nil, err
	}
	respBody, err := c.httpRequest(fmt.Sprintf("/faas/function/%s", updated_function.Id), "PUT", buf)
	if nil != err {
		return nil, err
	}
	returned_function := &Function{}
	err = json.NewDecoder(respBody).Decode(returned_function)
	if nil != err {
		return nil, err
	}
	return returned_function, nil
}

func (c *Client) GetFunctionCodeTemplate(args []string, language string) (*string, error) {
	buf := bytes.Buffer{}
	functionCodeTemplate := &FunctionCodeTemplate{
		Args:     args,
		Language: language,
	}
	err := json.NewEncoder(&buf).Encode(functionCodeTemplate)
	if nil != err {
		return nil, err
	}
	respBody, err := c.httpRequest(fmt.Sprintf("/faas/template"), "POST", buf)
	if nil != err {
		return nil, err
	}
	functionCodeTemplateResponse := &FunctionCodeTemplateResponse{}
	err = json.NewDecoder(respBody).Decode(functionCodeTemplateResponse)
	if nil != err {
		return nil, err
	}
	return &functionCodeTemplateResponse.Template, nil
}
