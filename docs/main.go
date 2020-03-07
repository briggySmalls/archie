package main

import (
	"github.com/briggysmalls/archie/cli/archie/cmd"
	"github.com/spf13/cobra/doc"
	"log"
	"os"
)

func main() {
	// Get the args provided on the command line
	argsWithoutProg := os.Args[1:]
	// The first argument is the output path
	outputPath := argsWithoutProg[0]
	// Generate the cobra documentation
	err := doc.GenMarkdownTree(cmd.RootCmd, outputPath)
	if err != nil {
		log.Fatal(err)
	}
}
