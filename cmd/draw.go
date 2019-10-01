package cmd

import (
	"fmt"
	"github.com/briggysmalls/archie/core/api"
	mdl "github.com/briggysmalls/archie/core/model"
	"github.com/briggysmalls/archie/core/views"
	"github.com/briggysmalls/archie/io/writers"
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
		var dat []byte
		dat, err = ioutil.ReadFile(model)
		if err != nil {
			panic(err)
		}

		// Parse the yaml into a model
		var m *mdl.Model
		m, err = api.ParseYaml(string(dat))
		if err != nil {
			panic(err)
		}

		// Create a view from the model
		var viewModel mdl.Model
		switch view {
		case "landscape":
			viewModel = views.NewLandscapeView(m)
		case "context":
			// First get the scope
			var scopeItem *mdl.Element
			scopeItem, err = m.LookupName(scope)
			if err != nil {
				panic(err)
			}
			viewModel = views.NewContextView(m, scopeItem)
		}

		// Draw the view
		d := drawers.NewPlantUmlDrawer()
		fmt.Print(d.Draw(viewModel))
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
