/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"github.com/spf13/cobra"
)

var contextScope string

// contextCmd represents the context command
var contextCmd = &cobra.Command{
	Use:   "context",
	Short: "Generates a context diagram",
	Long: `Generate a diagram that shows the context of the specified element.

[1] Main elements of interest
Children of the scoping element.

[2] Relevant associated elements
Those that are associated to one of the child elements of the scope, where either:
- The parent is an ancestor of scope
- It is a root element`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Generate the diagram
		diagram, err = arch.ContextView(contextScope)
	},
}

func init() {
	generateCmd.AddCommand(contextCmd)

	fs := contextCmd.Flags()
	fs.StringVarP(&contextScope, "scope", "s", "", "scope for the context")
	cobra.MarkFlagRequired(fs, "scope")
}
