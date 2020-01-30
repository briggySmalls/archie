package cmd

import (
	"fmt"
	"github.com/briggysmalls/archie"
	"github.com/briggysmalls/archie/cli/utils"
	"github.com/spf13/cobra"
	"io/ioutil"
)

var view string
var scope string
var tag string
var customFooter string
var generateModelFile string
var arch archie.Archie

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
	Run: func(cmd *cobra.Command, args []string) {
		// Create a view from the model
		var diagram string
		var err error
		switch view {
		case "landscape":
			diagram, err = arch.LandscapeView()
		case "context":
			// Construct the view
			diagram, err = arch.ContextView(scope)
		case "tag":
			diagram, err = arch.TagView(scope, tag)
		default:
			panic(fmt.Errorf("Unrecognised view: %s", view))
		}
		if err != nil {
			panic(err)
		}

		// Draw the view (print json for now)
		fmt.Print(diagram)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	// Add flags
	generateCmd.PersistentFlags().StringVarP(&generateModelFile, "model", "m", "", "Model to generate diagrams from")
	generateCmd.PersistentFlags().StringVarP(&view, "view", "v", "", "view to create")
	generateCmd.PersistentFlags().StringVarP(&scope, "scope", "s", "", "scope for the view")
	generateCmd.PersistentFlags().StringVarP(&tag, "tag", "t", "", "tag to filter by (tag diagram only)")
}
