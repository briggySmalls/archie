package views

import (
	mdl "github.com/briggysmalls/archie/internal/model"
)

// NewContextView creates a view that shows the context of the specified element.
// The view contains: a) all child elements of the scope, b) relevant associated elements.
// A relevant associated element is one that is associated to one of the child elements of the scope, where either:
// the parent is an ancestor of scope, or it is a root element.
func NewContextView(model *mdl.Model, scope mdl.Element) mdl.Model {
	// Find relevant elements
	var primary []mdl.Element
	if len(model.Children(scope)) > 0 {
		// The main elements of interest are the children of the scope
		primary = append(primary, model.Children(scope)...)
	} else {
		// The scope has no children, so add it
		primary = []mdl.Element{scope}
	}

	// We also want to add elements related to scope
	// where one of the following is true:
	// - parent is an ancestor of scope
	// - it is a root element
	var secondary []mdl.Element
	for _, rel := range model.ImplicitAssociations() {
		if linked := getLinked(rel, scope); linked != nil {
			// Fetch parent of linked element
			parent, err := model.Parent(linked)
			panicOnError(err)
			if parent == nil {
				// Add any linked root elements
				secondary = append(secondary, linked)
			} else if model.IsAncestor(scope, parent) {
				// Add elements who's parents are ancestors of scope
				secondary = append(secondary, linked)
			}
		}
	}

	// Create a model from the model's root elements
	view, err := CreateSubmodel(model, primary, secondary)
	// We shouldn't error (we've pulled elements out sensibly)
	panicOnError(err)
	return view
}
