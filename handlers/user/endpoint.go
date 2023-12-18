package user

import (
	"cwc/config"
	"fmt"
)

func HandlerGetDefaultEndpoint() {
	endpoint := config.GetDefaultEndpoint()
	fmt.Printf("Default endpoint = %v\n", endpoint)
}

func HandlerSetDefaultEndpoint(value string) {
	config.SetDefaultEndpoint(value)
	fmt.Printf("Default endpoint = %v\n", value)
}

func HandlerSetDefaultFormat(value string) {
	config.SetDefaultFormat(value)
	fmt.Printf("Default output format = %v\n", value)
}

func HandlerGetDefaultFormat() {
	format := config.GetDefaultFormat()
	fmt.Printf("Default output format = %v\n", format)
}
