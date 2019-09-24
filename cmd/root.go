package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var model string

var rootCmd = &cobra.Command{
	Use:   "archie",
	Short: "Archie is a simple system architecture model manager",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// Ensure a model is provided
		if model == "" {
			return fmt.Errorf("Model must be provided")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func init() {
	// Add some flags
	rootCmd.PersistentFlags().StringVar(&model, "model", "", "model file (yaml)")
	rootCmd.MarkFlagRequired("model")
	rootCmd.MarkFlagFilename("model")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
	    fmt.Println(err)
	    os.Exit(1)
	}
}
