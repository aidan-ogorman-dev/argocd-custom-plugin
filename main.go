package main

import (
	"os"

	"github.com/aidan-ogorman-dev/cmd"
)

func main() {
	if err := cmd.NewRootCommand().Execute(); err != nil {
		os.Exit(1)
	}
}
