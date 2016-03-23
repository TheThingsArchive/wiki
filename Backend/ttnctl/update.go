package main

import (
	"path"
	"strings"

	ttnctl "github.com/TheThingsNetwork/ttn/ttnctl/cmd"

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

	doc.GenMarkdownTreeCustom(ttnctl.RootCmd, "./", filePrepender, linkHandler)
}
