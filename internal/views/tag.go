package views

import (
	mdl "github.com/briggysmalls/archie/internal/model"
)

// NewTagView creates a view that shows the context of elements with a specified tag.
// The view contains: a) main elements of interest, b) relevant associated elements.
// The main elements of interest are the 'oldest' elements that have the specified tag.
// A relevant associated element is one that is associated to one of the child elements of the scope, where either:
// the parent is an ancestor of scope, or it is a root element.
func NewTagView(model *mdl.Model, scope mdl.Element, tag string) mdl.Model {
	// Find elements with correct tag
	taggedElements := findTaggedElements(model, scope, tag)
	// Find related elements
	relatedElements, err := findRelatedElements(model, taggedElements)
	panicOnError(err)
	// Create a model from the model's root elements
	view, err := CreateSubmodel(model, taggedElements, relatedElements)
	// We shouldn't error (we've pulled elements out sensibly)
	panicOnError(err)
	return view
}

func findTaggedElements(model *mdl.Model, scope mdl.Element, tag string) []mdl.Element {
	// Short-circuit if scope has no children
	if len(model.Children(scope)) == 0 {
		return []mdl.Element{scope}
	}
	// Otherwise search for children with the scope
	var elements []mdl.Element
	for _, child := range model.Children(scope) {
		// Check if the child has the tag
		if containsString(child.Tags(), tag) {
			// The element is tagged. Add and bail.
			elements = append(elements, child)
			continue
		}
		// Recursively search children for first sign of tag
		elements = append(elements, findTaggedElements(model, child, tag)...)
	}
	return elements
}

func findRelatedElements(model *mdl.Model, taggedElements []mdl.Element) ([]mdl.Element, error) {
	// We also want to add other, related, elements that either:
	// - Have any tag and haven't got a tagged ancestor
	// - Have no children, and haven't got a tagged ancestor
	var relatedElements []mdl.Element
	for _, rel := range model.ImplicitAssociations() {
		for _, el := range taggedElements {
			if linked := getLinked(rel, el); linked != nil {
				// This relationship links one of the tagged elements
				if parent, err := model.Parent(linked); err != nil {
					return nil, err
				} else if hasTaggedParent, err := checkForTaggedAnscestor(model, parent); err != nil {
					return nil, err
				} else if hasTaggedParent == false {
					// The linked element has no tagged parent
					if len(linked.Tags()) > 0 || len(model.Children(linked)) == 0 {
						relatedElements = append(relatedElements, linked)
					}
				}
			}
		}
	}
	return relatedElements, nil
}

func checkForTaggedAnscestor(model *mdl.Model, el mdl.Element) (bool, error) {
	if el == nil {
		// We've reached the top
		return false, nil
	}
	// Check if it has a tag
	if len(el.Tags()) > 0 {
		return true, nil
	}
	// Recurse up the parent
	parent, err := model.Parent(el)
	if err != nil {
		return false, err
	}
	return checkForTaggedAnscestor(model, parent)
}

func containsString(haystack []string, needle string) bool {
	for _, query := range haystack {
		if query == needle {
			return true
		}
	}
	return false
}
