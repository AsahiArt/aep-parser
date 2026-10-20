package main

import (
	"os"

	"github.com/AsahiArt/aep-parser/internal/toolkitcli"
)

func main() {
	os.Exit(toolkitcli.Run(os.Args[1:], os.Stdout, os.Stderr))
}
