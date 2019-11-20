package views

import (
	mdl "github.com/briggysmalls/archie/internal/model"
)

// Create a context view
// TODO: This should probably return an error
func NewTagView(model *mdl.Model, scope mdl.Element, tag string) mdl.Model {
	// Find elements with correct tag
	taggedElements := findElements(model, scope, tag)

	// We also want to add other elements that:
	// - Have a tag (untagged means a simple logical divider)
	// - Haven't got a tagged parent (too specific/deep)
	var relatedElements []mdl.Element
	for _, rel := range model.ImplicitAssociations() {
		for _, el := range taggedElements {
			if linked := getLinked(rel, el); linked != nil {
				// This relationship links one of the tagged elements
				if parent, err := model.Parent(linked); err != nil {
					panicOnError(err)
				} else if parent != nil && len(parent.Tags()) == 0 && len(linked.Tags()) > 0 {
					relatedElements = append(relatedElements, linked)
				}
			}
		}
	}

	// Create a model from the model's root elements
	view, err := CreateSubmodel(model, taggedElements, relatedElements)
	// We shouldn't error (we've pulled elements out sensibly)
	panicOnError(err)
	return view
}

func findElements(model *mdl.Model, scope mdl.Element, tag string) []mdl.Element {
	var elements []mdl.Element
	for _, child := range model.Children(scope) {
		// Check if the child has the tag
		if containsString(child.Tags(), tag) {
			// The element is tagged. Add and bail.
			elements = append(elements, child)
			continue
		}
		// Recursively search children for first sign of tag
		elements = append(elements, findElements(model, child, tag)...)
	}
	return elements
}

func containsString(haystack []string, needle string) bool {
	for _, query := range haystack {
		if query == needle {
			return true
		}
	}
	return false
}
