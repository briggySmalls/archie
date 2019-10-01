package main

import (
	"github.com/briggysmalls/archie/core/api"
	"github.com/briggysmalls/archie/core/views"
	"syscall/js"
)

func newLandscapeView(this js.Value, inputs []js.Value) interface{} {
	// Assume yaml is passed in as first argument
	model := inputs[0].String()

	// Read the model in
	m, err := api.ParseYaml(model)
	if err != nil {
		return err
	}

	// Convert to a view
	v := views.NewLandscapeView(m)
	if err != nil {
		return err
	}

	// Return the view as json
	json, err := api.ToJson(&v)
	if err != nil {
		return err
	}
	return json
}

func main() {
	// Create a channel to keep us open on
	c := make(chan bool)
	// Register functions for the application
	js.Global().Set("newLandscapeView", js.FuncOf(newLandscapeView))
	// Keep the application running
	<-c
}
