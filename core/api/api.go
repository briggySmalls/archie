package api

import (
	"github.com/briggysmalls/archie/core/io"
	mdl "github.com/briggysmalls/archie/core/model"
	"github.com/briggysmalls/archie/core/views"
)

type Archie interface {
	LandscapeView() (string, error)
	ContextView(element string) (string, error)
}

type archie struct {
	model *mdl.Model
}

func NewArchieFromJson(json string) (Archie, error) {
	// Convert the JSON to a model
	model, err := io.ParseJson(json)
	if err != nil {
		return nil, err
	}
	// Return a new archie
	return &archie{model: model}, nil
}

func NewArchieFromYaml(yaml string) (Archie, error) {
	// Convert the JSON to a model
	model, err := io.ParseYaml(yaml)
	if err != nil {
		return nil, err
	}
	// Return a new archie
	return &archie{model: model}, nil
}

func (a *archie) LandscapeView() (json string, err error) {
	// Create the view
	view := views.NewLandscapeView(a.model)
	// Convert to json
	json, err = io.ToJson(&view)
	return
}

func (a *archie) ContextView(scope string) (json string, err error) {
	// Lookup the element
	element, err := a.model.LookupName(scope)
	if err != nil {
		return
	}
	// Create the view
	view := views.NewContextView(a.model, element)
	// Convert to json
	json, err = io.ToJson(&view)
	return
}
