package views

import (
	mdl "github.com/briggysmalls/archie/internal/model"
)

// Create a context view
// TODO: This should probably return an error
func NewContextView(model *mdl.Model, scope mdl.Element) mdl.Model {
	// Find relevant elements
	var elements []mdl.Element
	if len(model.Children(scope)) > 0 {
		// The main elements of interest are the children of the scope
		elements = append(elements, model.Children(scope)...)
	} else {
		// The scope has no children, so add it
		elements = []mdl.Element{scope}
	}

	// We also want to add elements related to scope
	// where one of the following is true:
	// - parent is an ancestor of scope
	// - it is a root element
	for _, rel := range model.ImplicitAssociations() {
		if linked := getLinked(rel, scope); linked != nil {
			// Fetch parent of linked element
			parent, err := model.Parent(linked)
			panicOnError(err)
			if parent == nil {
				// Add any linked root elements
				elements = append(elements, linked)
			} else if model.IsAncestor(scope, parent) {
				// Add elements who's parents are ancestors of scope
				elements = append(elements, linked)
			}
		}
	}

	// Create a model from the model's root elements
	view, err := CreateSubmodel(model, elements)
	// We shouldn't error (we've pulled elements out sensibly)
	panicOnError(err)
	return view
}

// Get the linked element, if the specified element is in the relationship
func getLinked(relationship mdl.Relationship, element mdl.Element) mdl.Element {
	if relationship.Source() == element {
		return relationship.Destination()
	}
	if relationship.Destination() == element {
		return relationship.Source()
	}
	return nil
}
