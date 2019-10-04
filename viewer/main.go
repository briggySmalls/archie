package main

import (
	"github.com/briggysmalls/archie"
	"github.com/briggysmalls/archie/internal/writers"
	"syscall/js"
)

var arch archie.Archie

// Initialise a new instance of archie
func Init(this js.Value, inputs []js.Value) interface{} {
	// Assume yaml is passed in as first argument
	yaml := inputs[0].String()
	// Create a new archie
	var err error
	arch, err = archie.New(writers.MermaidStrategy{}, yaml)
	if err != nil {
		return nil
	}
	return true
}

func Elements(this js.Value, inputs []js.Value) interface{} {
	els := make(map[string]interface{})
	for id, name := range arch.Elements() {
		els[id] = name
	}
	return els
}

// Create a landscape view
func LandscapeView(this js.Value, inputs []js.Value) interface{} {
	// Construct the view
	result, err := arch.LandscapeView()
	if err != nil {
		return nil
	}
	return result
}

// Create a context view
func ContextView(this js.Value, inputs []js.Value) interface{} {
	// Assume scope is passed in as first argument
	scope := inputs[0].String()

	// Construct the view
	result, err := arch.ContextView(scope)
	if err != nil {
		return nil
	}
	return result
}

func main() {
	funcs := map[string]func(js.Value, []js.Value) interface{}{
		"init":          Init,
		"landscapeView": LandscapeView,
		"contextView":   ContextView,
		"elements":      Elements,
	}
	// Register the functions
	for name, handle := range funcs {
		js.Global().Set(name, js.FuncOf(handle))
	}
	// Wait indefinitely
	c := make(chan bool)
	<-c
}
