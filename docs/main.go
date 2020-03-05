package main

import (
	"log"

	"github.com/briggysmalls/archie/cli/cmd"
	"github.com/spf13/cobra/doc"
)

func main() {
	err := doc.GenMarkdownTree(cmd.RootCmd, "/tmp")
	if err != nil {
		log.Fatal(err)
	}
}
