package main

import (
	"os"

	"github.com/aidan-ogorman-dev/argocd-custom-plugin/cmd"
)

func main() {
	if err := cmd.NewRootCommand().Execute(); err != nil {
		os.Exit(1)
	}
}
