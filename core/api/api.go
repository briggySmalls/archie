package api

import (
	"github.com/briggysmalls/archie/core/io"
	mdl "github.com/briggysmalls/archie/core/model"
	"github.com/briggysmalls/archie/core/views"
	"github.com/briggysmalls/archie/core/writers"
)

type Archie interface {
	LandscapeView() (string, error)
	ContextView(element string) (string, error)
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
