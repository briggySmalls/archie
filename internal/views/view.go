package views

import (
	"github.com/briggysmalls/archie/internal/types"
)

// CreateSubmodel creates a sub-model from the full model
func CreateSubmodel(model *types.Model, elements []*types.ModelElement) types.Model {
	// Get the relevant elements (including implicit)
	relevantElements := getRelevantElements(model, elements)
	// Create a new model
	subModel := types.NewModel()
	// Copy over the elements
	for _, el := range model.Elements() {
		// Add the root elements, if relevant
		if _, ok := relevantElements[el]; ok {
			subModel.AddRootElement(el.Data)
		}
		// Recurse through children
		copyChildren(&subModel, relevantElements, el)
	}
	// Copy over relevant relationships
	for _, rel := range getRelevantRelationships(model, elements) {
		subModel.AddRelationship(rel.Source.Data, rel.Destination.Data)
	}
	return subModel
}

// Copy the children of the specified element, recursively
func copyChildren(dest *types.Model, relevant map[*types.ModelElement]bool, element *types.ModelElement) {
	for _, child := range element.Children {
		// Skip child (and all its children) if not relevant
		if !relevant[child] {
			continue
		}
		// Add the child
		dest.AddChild(element.Data, child.Data)
		// Add the child's children
		copyChildren(dest, relevant, child)
	}
}

// Select the relationships that are relevant, including implicit ones
func getRelevantRelationships(model *types.Model, elements []*types.ModelElement) []types.Relationship {
	var relationships []types.Relationship
	for _, rel := range model.ImplicitRelationships() {
		// Add relationships that link relevant elements
		if contains(elements, rel.Source) && contains(elements, rel.Destination) {
			relationships = append(relationships, rel)
		}
	}
	return relationships
}

// Select the elements that are relevant, including the implicit ones
func getRelevantElements(model *types.Model, elements []*types.ModelElement) map[*types.ModelElement]bool {
	// Prepare an empty index of relevant elements
	relevant := make(map[*types.ModelElement]bool)
	// Add elements
	for _, el := range elements {
		// First, add the core element
		addAllAncestors(model, relevant, el)
	}
	return relevant
}

// Add the specified element to the map, and all its ancestors
func addAllAncestors(model *types.Model, elements map[*types.ModelElement]bool, el *types.ModelElement) map[*types.ModelElement]bool {
	for {
		// Add the current element
		elements[el] = true
		// Try add parent
		if parent, err := model.Parent(el); err != nil {
			// Parent isn't in model
			panic(err)
		} else if parent == nil {
			// Parent is root, we're done here
			return elements
		} else if _, ok := elements[parent]; ok {
			// Parent is already in map, someone got there first
			// Return early
			return elements
		} else {
			// Now recurse up through all the parents
			addAllAncestors(model, elements, parent)
		}
	}
}

// Check if an element is in the slice
func contains(haystack []*types.ModelElement, needle *types.ModelElement) bool {
	for _, el := range haystack {
		if el == needle {
			return true
		}
	}
	return false
}
