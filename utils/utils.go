package utils

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/mail"
	"os"
	"reflect"
	"strings"
)

func If[T any](cond bool, vtrue, vfalse T) T {
	if cond {
		return vtrue
	}
	return vfalse
}

func IsNotBlank(str string) bool {
	return len(str) > 0 && strings.TrimSpace(str) != ""
}

func IsBlank(str string) bool {
	return !IsNotBlank(str)
}

func IsValidEmail(email string) bool {
	if IsBlank(email) {
		return false
	}

	_, err := mail.ParseAddress(email)
	return err == nil
}

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
	if nil != err {
		fmt.Println("An error occured while reading input. Please try again", err)
		return ""
	}

	// remove the delimeter from the string
	value = strings.TrimSuffix(value, "\n")
	return value
}

func PrintJson(class interface{}) {
	marchal_class, err := json.Marshal(class)
	ExitIfError(err)

	fmt.Printf("%s\n", JsonPrettyPrint(string(marchal_class)))
}

func PrintHeader(class interface{}) {
	values := reflect.ValueOf(class)
	typesOf := values.Type()
	headerMsg := ""
	for i := 0; i < values.NumField(); i++ {
		headerMsg = fmt.Sprintf("%s%s\t", headerMsg, typesOf.Field(i).Name)
	}

	fmt.Println(headerMsg)
}

func PrintPretty(firstLine string, class interface{}) {
	fmt.Printf("%s:\n", firstLine)

	values := reflect.ValueOf(class)
	typesOf := values.Type()
	for i := 0; i < values.NumField(); i++ {
		v := values.Field(i).Interface()
		if nil == v || v == "" || v == 0 {
			continue
		}

		k := strings.Replace(typesOf.Field(i).Name, "_", " ", -1)
		fmt.Printf("  ➤ %s: %s\n", k, strings.TrimSpace(fmt.Sprintf("%v", v)))
	}
}

func PrintPrettyArray(firstLine string, lst []string) {
	fmt.Printf("%s:\n", firstLine)

	for _, elem := range lst {
		fmt.Printf("  ➤ %v\n", elem)
	}
}

func PrintArray(lst []string) {
	fmt.Printf("%s\n", strings.Join(lst, "\n"))
}

func PrintRow(class interface{}) {
	PrintHeader(class)
	values := reflect.ValueOf(class)
	valuesMsg := ""
	for i := 0; i < values.NumField(); i++ {
		v := values.Field(i).Interface()
		valuesMsg = fmt.Sprintf("%s%v\t", valuesMsg, strings.TrimSpace(fmt.Sprintf("%v", v)))
	}

	fmt.Println(valuesMsg)
}

func PrintMultiRow(type_class interface{}, class interface{}) {
	PrintHeader(type_class)
	s := reflect.ValueOf(class)
	for i := 0; i < s.Len(); i++ {
		v := reflect.Indirect(s.Index(i))
		valuesMsg := ""
		for i := 0; i < v.NumField(); i++ {
			valuesMsg = fmt.Sprintf("%s%v\t", valuesMsg, v.Field(i).Interface())
		}

		fmt.Println(valuesMsg)
	}
}

func JsonPrettyPrint(in string) string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(in), "", "\t")
	if nil != err {
		return in
	}

	return out.String()
}

func ExitIfError(err error) {
	ExitIfErrorWithMsg("Error", err)
}

func ExitIfErrorWithouMsg(err error) {
	ExitIfNeeded("", nil != err)
}

func ExitIfErrorWithMsg(msg string, err error) {
	ExitIfNeeded(fmt.Sprintf("%s: %s", msg, err), nil != err)
}

func ExitIfNeeded(msg string, exit bool) {
	if exit {
		if IsNotBlank(msg) {
			fmt.Println(msg)
		}
		os.Exit(1)
	}
}

func GetSystemEditor() string {
	editorCommand := os.Getenv("EDITOR")
	if IsBlank(editorCommand) {
		editorCommand = "vi"
	}

	return editorCommand
}

func ShortName(name string, hash string) string {
	if name == "" {
		return ""
	}
	if hash == "" {
		lastDashIndex := strings.LastIndex(name, "-")
		if lastDashIndex != -1 {
			return name[:lastDashIndex]
		}
		return name
	}
	hashWithDash := "-" + hash
	if strings.HasSuffix(name, hashWithDash) {
		return name[:len(name)-len(hashWithDash)]
	}
	return name
}
