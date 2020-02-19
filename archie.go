// Package archie provides a tool for generating diagrams from a system architecture model
package archie

import (
	"github.com/briggysmalls/archie/internal/io"
	mdl "github.com/briggysmalls/archie/internal/model"
	"github.com/briggysmalls/archie/internal/views"
	"github.com/briggysmalls/archie/writers"
)

// Archie diagram tool
type Archie interface {
	LandscapeView() (string, error)
	ContextView(element string) (string, error)
	TagView(element, tag string) (string, error)
	Elements() map[string]string
}

type archie struct {
	model  *mdl.Model
	writer writers.Writer
}

// New creates a new Archie instance from the provided YAML model.
// The provided writer strategy determines how to render a view.
func New(strategy writers.Strategy, yaml string) (Archie, error) {
	// Convert the yaml to a model
	model, err := io.ParseYaml(yaml)
	if err != nil {
		return nil, err
	}
	// Return a new archie
	return &archie{
		model:  model,
		writer: writers.New(strategy),
	}, nil
}

// LandscapeView creates a top-level diagram of the system.
// The landscape comprises of root actors and elements, and the associations between them.
func (a *archie) LandscapeView() (diagram string, err error) {
	// Create the view
	view := views.NewLandscapeView(a.model)
	// Convert to diagram
	diagram, err = a.writer.Write(view)
	return
}

// ContextView creates a diagram that shows the context of the specified element.
// The view contains: a) main elements of interest, b) relevant associated elements.
// The main elements of interest are those that are children of the scoping element.
// A relevant associated element is one that is associated to one of the child elements of the scope, where either:
// the parent is an ancestor of scope, or it is a root element.
func (a *archie) ContextView(scope string) (diagram string, err error) {
	// Lookup the element
	element, err := a.model.LookupName(scope, nil)
	if err != nil {
		return
	}
	// Create the view
	view := views.NewContextView(a.model, element)
	// Convert to diagram
	diagram, err = a.writer.Write(view)
	return
}

// TagView creates a diagram that shows the context of elements with a specified tag.
// The view contains: a) main elements of interest, b) relevant associated elements.
// The view contains: a) the 'oldest' element with the specified tag, b) relevant associated elements.
// The main elements of interest are the 'oldest' elements that have the specified tag.
// A relevant associated element is one that is associated to one of the child elements of the scope, where either:
// the parent is an ancestor of scope, or it is a root element.
func (a *archie) TagView(scope, tag string) (diagram string, err error) {
	// Lookup the element
	element, err := a.model.LookupName(scope, nil)
	if err != nil {
		return
	}
	// Create the view
	view := views.NewTagView(a.model, element, tag)
	// Convert to diagram
	diagram, err = a.writer.Write(view)
	return
}

// Elements returns a map of element ID to element name
func (a *archie) Elements() map[string]string {
	// Prepare a slice
	elementLookup := make(map[string]string)
	// Copy element names in
	for _, el := range a.model.Elements {
		// Get the full name of the element
		name, err := a.model.Name(el)
		if err != nil {
			// We are iterating through the model elements, so we should definitely find their name
			panic(err)
		}
		// Add to the slice
		elementLookup[el.ID()] = name
	}
	// Return the names
	return elementLookup
}
