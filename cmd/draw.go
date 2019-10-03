package cmd

import (
	"fmt"
	"github.com/briggysmalls/archie/core"
	"github.com/briggysmalls/archie/core/writers"
	"io/ioutil"

	"github.com/spf13/cobra"
)

var view string
var scope string

// drawCmd represents the draw command
var drawCmd = &cobra.Command{
	Use:   "draw",
	Short: "Draw a diagram from the model",
	Long:  ``,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		// Ensure a view is provided
		if view == "" {
			return fmt.Errorf("View must be provided")
		}
		// Ensure a scope is provided if required
		if view != "landscape" && scope == "" {
			return fmt.Errorf("Scope must be provided for view: %s", view)
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		// Read in the yaml file
		yaml, err := ioutil.ReadFile(model)
		if err != nil {
			panic(err)
		}

		// Parse the yaml into a model
		m, err := core.New(writers.PlantUmlStrategy{}, string(yaml))
		if err != nil {
			panic(err)
		}

		// Create a view from the model
		var viewModel string
		switch view {
		case "landscape":
			viewModel, err = m.LandscapeView()
		case "context":
			viewModel, err = m.ContextView(scope)
		}
		if err != nil {
			panic(err)
		}

		// Draw the view (print json for now)
		fmt.Print(viewModel)
	},
}

func init() {
	// Add as a subcommand
	rootCmd.AddCommand(drawCmd)
	// Add flags
	drawCmd.PersistentFlags().StringVar(&view, "view", "", "view to create")
	drawCmd.PersistentFlags().StringVar(&scope, "scope", "", "scope for the view")
	// Mark some as required
	drawCmd.MarkFlagRequired("view")
}
