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
	"github.com/briggysmalls/archie"
	"github.com/briggysmalls/archie/writers"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
)

var view string
var scope string
var tag string
var customFooter string
var generateModelFile string
var generateModelYaml string

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a diagram from an architecture model",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Read in the yaml file
		model, err := ioutil.ReadFile(generateModelFile)
		handleError(err)
		generateModelYaml = string(model)
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Create a writer from config
		writer := writers.PlantUmlStrategy{CustomFooter: viper.GetString("footer")}
		// Create an archie from the yaml
		arch, err := archie.New(writer, generateModelYaml)
		if err != nil {
			panic(err)
		}
		// Create a view from the model
		var diagram string
		switch view {
		case "landscape":
			diagram, err = arch.LandscapeView()
		case "context":
			// Construct the view
			diagram, err = arch.ContextView(scope)
		case "tag":
			diagram, err = arch.TagView(scope, tag)
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
