package handlers

import (
	"flag"
	"fmt"
	"os"
	"runtime"
)

func HandleVersion(versionCmd *flag.FlagSet, version string) {
	versionCmd.Parse(os.Args[2:])
	fmt.Printf("cwc-cli/%v %v %v\n", version, runtime.GOOS, runtime.GOARCH)
}
