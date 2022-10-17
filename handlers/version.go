package handlers

import (
	"fmt"
	"runtime"
)

func HandleVersion(version string) {

	fmt.Printf("cwc-cli/%v %v %v\n", version, runtime.GOOS, runtime.GOARCH)
}
