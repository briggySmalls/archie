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
	"github.com/gobuffalo/packr"
	"github.com/spf13/cobra"
	"net/http"
	"text/template"
)

// browseCmd represents the browse command
var browseCmd = &cobra.Command{
	Use:   "browse",
	Short: "Browse an architecture model",
	Long: `Launches a browser application for navigating the
different diagrams available from the provided model`,
	Run: func(cmd *cobra.Command, args []string) {
		// Grab the static files
		box := packr.NewBox("../bin")
		// Handle requests
		http.HandleFunc("/", pageHandler)
		http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(box)))
		// Start the server
		err := http.ListenAndServe(":8080", nil)
		handleError(err)
	},
}

func init() {
	rootCmd.AddCommand(browseCmd)
}

func pageHandler(w http.ResponseWriter, r *http.Request) {
	// Compile the template into a page
	t, err := template.ParseFiles("./cmd/index.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, modelYaml)
}

func fileHandler(w http.ResponseWriter, r *http.Request) {

}
