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
	"github.com/briggysmalls/archie/cli/server"

	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "REST API for archie system architecture tool",
	Long: `Exposes endpoints for obtaining plantuml views from a supplied model

All endpoints respond to a POST request that supplies the model (YAML) in the body`,
	Run: func(cmd *cobra.Command, args []string) {
		// Get the desired port from configuration
	    port := viper.GetInt("port")
	    // Run the server
	    fmt.Printf("Starting server on port: %d\n", port)
	    if err := server.Serve(fmt.Sprintf(":%d", port), viper.GetString("footer")); err != nil {
	      panic(err)
	    }
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)

	// Server port (bind to config too)
	apiCmd.Flags().Int("port", 8080, "Port to run server on")
	viper.BindPFlag("port", apiCmd.Flags().Lookup("port"))
}
