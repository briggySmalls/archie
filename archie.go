package archie

import (
	"github.com/briggysmalls/archie/internal/io"
	mdl "github.com/briggysmalls/archie/internal/model"
	"github.com/briggysmalls/archie/internal/views"
	"github.com/briggysmalls/archie/writers"
)

type Archie interface {
	LandscapeView() (string, error)
	ContextView(element string) (string, error)
	Elements() map[string]string
}

type archie struct {
	model  *mdl.Model
	writer writers.Writer
}

func New(strategy writers.Strategy, yaml string) (Archie, error) {
	// Convert the yaml to a model
	model, err := io.ParseYaml(yaml)
	if err != nil {
		return nil, err
	}
	// Create a new writer
	w := writers.New(strategy)
	// Return a new archie
	return &archie{
		model:  model,
		writer: &w,
	}, nil
}

func (a *archie) LandscapeView() (diagram string, err error) {
	// Create the view
	view := views.NewLandscapeView(a.model)
	// Convert to diagram
	diagram, err = a.writer.Write(view)
	return
}

func (a *archie) ContextView(scope string) (diagram string, err error) {
	// Lookup the element
	element, err := a.model.LookupName(scope)
	if err != nil {
		return
	}
	// Create the view
	view := views.NewContextView(a.model, element)
	// Convert to diagram
	diagram, err = a.writer.Write(view)
	return
}

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
