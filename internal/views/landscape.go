package views

import (
	mdl "github.com/briggysmalls/archie/internal/model"
)

// NewLandscapeView creates a top-level view of the system.
// The landscape comprises of root actors and elements, and the associations between them.
func NewLandscapeView(model *mdl.Model) mdl.Model {
	// Create a model from the model's root elements
	view, err := CreateSubmodel(model, model.RootElements(), []mdl.Element{})
	// We shouldn't error (we've pulled elements out sensibly)
	panicOnError(err)
	return view
}
