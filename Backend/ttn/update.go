package main

import (
	"path"
	"strings"

	ttn "github.com/TheThingsNetwork/ttn/cmd"

	"github.com/spf13/cobra/doc"
)

func main() {
	filePrepender := func(filename string) string {
		return ""
	}

	linkHandler := func(name string) string {
		base := strings.TrimSuffix(name, path.Ext(name))
		return base
	}

	doc.GenMarkdownTreeCustom(ttn.RootCmd, "./", filePrepender, linkHandler)
}
