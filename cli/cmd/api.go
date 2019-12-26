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
		if err := server.Serve(fmt.Sprintf(":%d", port)); err != nil {
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
