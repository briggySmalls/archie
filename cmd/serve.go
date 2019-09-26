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
	"github.com/briggysmalls/archie/internal/readers"
	"github.com/briggysmalls/archie/internal/server"
	"github.com/briggysmalls/archie/internal/types"
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
		var m *types.Model
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
		go open("http://localhost:8080")
		// Serve (blocking call)
		s.Serve(":8080")
	},
}

func init() {
	// Add as a subcommand
	rootCmd.AddCommand(serveCmd)
}

// open opens the specified URL in the default browser of the user.
func open(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}
