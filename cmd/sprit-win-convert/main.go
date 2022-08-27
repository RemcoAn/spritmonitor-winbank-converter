package main

import (
	"os"
	"win-sprit-converter/cmd/sprit-win-convert/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
