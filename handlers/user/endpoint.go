package user

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
	fmt.Printf("Default endpoint = %v\n", value)

}

func HandlerSetDefaultFormat(value string) {

	client.SetDefaultFormat(value)
	fmt.Printf("Default output format = %v\n", value)

}

func HandlerGetDefaultFormat() {

	format := client.GetDefaultFormat()
	fmt.Printf("Default output format = %v\n", format)

}
