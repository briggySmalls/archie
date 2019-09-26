/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	mdl "github.com/briggysmalls/archie/core/model"
	"github.com/briggysmalls/archie/core/views"
	"github.com/briggysmalls/archie/io/readers"
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
		m, err = readers.ParseYaml(string(dat))
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
			viewModel = views.NewItemContextView(m, scopeItem)
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
