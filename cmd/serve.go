package cmd

import (
	"fmt"
	mdl "github.com/briggysmalls/archie/core/model"
	"github.com/briggysmalls/archie/io/readers"
	"github.com/briggysmalls/archie/server"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os/exec"
	"runtime"
)

// serveCmd represents the draw command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serves an architecture viewer",
	Long:  ``,
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

		// Create a server
		s, err := server.NewServer(m)
		if err != nil {
			panic(err)
		}
		// Open browser in goroutine
		go openBrowser("http://localhost:8080")
		// Serve (blocking call)
		s.Serve(":8080")
	},
}

func init() {
	// Add as a subcommand
	rootCmd.AddCommand(serveCmd)
}

func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		panic(err)
	}
}
