package handlers

import (
	"cwc/client"
	"fmt"
)

func HandlerGetDefaultEndpoint() {

	endpoint := client.GetDefaultEndpoint()
	fmt.Printf("Default endpoint = %v\n", endpoint)

}

func HandlerSetDefaultEndpoint(value string) {

	client.SetDefaultEndpoint(value)

}
