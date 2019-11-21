package views

import (
	mdl "github.com/briggysmalls/archie/internal/model"
)

// Create a system landscape view
func NewLandscapeView(model *mdl.Model) mdl.Model {
	// Create a model from the model's root elements
	view, err := CreateSubmodel(model, model.RootElements(), []mdl.Element{})
	// We shouldn't error (we've pulled elements out sensibly)
	panicOnError(err)
	return view
}
