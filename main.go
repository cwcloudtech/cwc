// Package main provides an entry point for cwc application.
package main

import "cwc/cmd"

// Version represents the application version.
var Version = "dev"

func main() {
	cmd.Execute(Version)
}
