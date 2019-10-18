package main

import (
	"fmt"
	"github.com/briggysmalls/archie"
	"strings"
	"syscall/js"
)

var arch archie.Archie
var diagramID string
var diagramTitleDom js.Value
var idNameMap map[string]string

// Initialise a new instance of archie
func Init(this js.Value, inputs []js.Value) interface{} {
	// Assume yaml is passed in as first argument
	yaml := inputs[0].String()
	// Assume diagram element ID is passed in as second argument
	diagramID = inputs[1].String()
	// Assume diagram title ID is passed in as third argument
	diagramTitleDom = js.Global().Get("document").Call("getElementById", inputs[2].String())
	// Create a new archie
	var err error
	arch, err = archie.New(strategy{}, yaml)
	if err != nil {
		return nil
	}
	// Record the id/name map
	idNameMap = arch.Elements()

	// Set the diagram to be the landscape
	diagram, err := arch.LandscapeView()
	if err != nil {
		panic(err)
	}
	updateDiagram(diagram, "System landscape")

	return true
}

func MermaidCallback(this js.Value, inputs []js.Value) interface{} {
	// The nodeId is the paramter
	id, err := toId(inputs[0].String())
	if err != nil {
		panic(err)
	}
	// Get the name from the node
	name, ok := idNameMap[id]
	if !ok {
		panic(fmt.Errorf("Failed to find element with ID '%s'", id))
	}
	// Update the diagram with the new context
	setContext(name)
	return nil
}

func Elements(this js.Value, inputs []js.Value) interface{} {
	returnVal := make(map[string]interface{})
	for key, val := range arch.Elements() {
		returnVal[key] = val
	}
	return returnVal
}

func setContext(scope string) {
	// Construct the view
	diagram, err := arch.ContextView(scope)
	if err != nil {
		panic(err)
	}
	updateDiagram(diagram, fmt.Sprintf("Context view: %s", scope))
}

func updateDiagram(diagram, title string) {
	// Update the archie element
	el := js.Global().Get("document").Call("getElementById", diagramID)
	el.Set("innerHTML", diagram)
	el.Call("removeAttribute", "data-processed")
	// Reinitialise mermaid
	js.Global().Get("mermaid").Call("init", js.Undefined(), fmt.Sprintf("#%s", diagramID))
	// Set title
	diagramTitleDom.Set("innerHTML", title)
}

func toId(mermaidId string) (string, error) {
	// Split at "id"
	parts := strings.Split(mermaidId, "id-")
	if len(parts) != 2 {
		return "", fmt.Errorf("ID '%s' expected to begin with 'id-'", mermaidId)
	}
	return parts[1], nil
}

func main() {
	funcs := map[string]func(js.Value, []js.Value) interface{}{
		"init":            Init,
		"mermaidCallback": MermaidCallback,
		"elements":        Elements,
	}
	// Register the functions
	for name, handle := range funcs {
		js.Global().Set(name, js.FuncOf(handle))
	}
	// Wait indefinitely
	c := make(chan bool)
	<-c
}
