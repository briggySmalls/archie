package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/briggysmalls/archie"
	"github.com/briggysmalls/archie/cli/utils"
	"github.com/spf13/cobra"
	"log"
)

var generateModelFile string
var arch archie.Archie
var diagram string
var err error

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a diagram from an architecture model",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Read in the yaml file
		modelAndConfig, err := ioutil.ReadFile(generateModelFile)
		handleError(err)
		arch, err = utils.ReadModel(modelAndConfig)
		handleError(err)
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		// Check the global error
		if err != nil {
			log.Fatal(err)
		}
		// Draw the view (print json for now)
		fmt.Print(diagram)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	// Add flags
	pfs := generateCmd.PersistentFlags()
	pfs.StringVarP(&generateModelFile, "model", "m", "", "Model to generate diagrams from")
	cobra.MarkFlagRequired(pfs, "model")
}
