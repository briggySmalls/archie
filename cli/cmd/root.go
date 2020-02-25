// Package cmd provides cobra command line configuration
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "archie",
	Short: "Tool for developing system architecture",
	Long: `
archie-cli is an application that utilises the archie framework
for developing lightweight system architecture modules`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	handleError(err)
}

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
