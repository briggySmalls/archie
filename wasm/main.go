package main

import (
	"github.com/briggysmalls/archie/core/io/yaml"
	"github.com/briggysmalls/archie/core/views"
	"syscall/js"
)

func newLandscapeView(this js.Value, inputs []js.Value) interface{} {
	// Assume yaml is passed in as first argument
	model := inputs[0].String()

	// Read the model in
	m, err := yaml.ParseYaml(model)
	if err != nil {
		return err
	}

	// Convert to a view
	v := views.NewLandscapeView(m)
	if err != nil {
		return err
	}

	// Return the view as yaml
	// TODO: This is ridiculous in WASM. We should return as a Javascript Object
	view, err := yaml.ToYaml(&v)
	if err != nil {
		return err
	}
	return view
}

func main() {
	// Create a channel to keep us open on
	c := make(chan bool)
	// Register functions for the application
	js.Global().Set("newLandscapeView", js.FuncOf(newLandscapeView))
	// Keep the application running
	<-c
}
