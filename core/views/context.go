package views

import (
	mdl "github.com/briggysmalls/archie/core/model"
)

// Create a context view
// TODO: This should probably return an error
func NewContextView(model *mdl.Model, scope *mdl.Element) mdl.Model {
	// Find relevant elements
	var elements []*mdl.Element

	if len(model.Children(scope)) > 0 {
		// The main elements of interest are the children of the scope
		elements = append(elements, model.Children(scope)...)
	} else {
		// The scope has no children, so add it
		elements = []*mdl.Element{scope}
	}

	// We also want to add elements:
	// - related to these children
	// - no deeper than scope
	// - only root element for independent systems
	scopeDepth, err := model.Depth(scope)
	if err != nil {
		panic(err)
	}
	for _, rel := range model.ImplicitAssociations() {
		if linked := getLinked(rel, scope); linked != nil {
			// Association links an element of interest...
			linkedDepth, err := model.Depth(linked)
			if err != nil {
				panic(err)
			}
			// If there is no common ancestor, we only want the root
			if !model.ShareAncestor(scope, linked) && linkedDepth != 0 {
				continue
			}
			// Ensure the element is not more specific than scope
			if linkedDepth <= scopeDepth {
				elements = append(elements, linked)
			}
		}
	}

	// Create a model from the model's root elements
	view, err := CreateSubmodel(model, elements)
	// We shouldn't error (we've pulled elements out sensibly)
	if err != nil {
		panic(err)
	}
	return view
}

// Get the linked element, if the specified element is in the relationship
func getLinked(relationship mdl.Relationship, element *mdl.Element) *mdl.Element {
	if relationship.Source == element {
		return relationship.Destination
	}
	if relationship.Destination == element {
		return relationship.Source
	}
	return nil
}
