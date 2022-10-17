package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func PromptUserForValue() string {
	reader := bufio.NewReader(os.Stdin)
	// ReadString will block until the delimiter is entered
	value, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return ""
	}

	// remove the delimeter from the string
	value = strings.TrimSuffix(value, "\n")
	return value
}
