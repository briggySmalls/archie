package views

import (
	"github.com/briggysmalls/archie/internal/types"
)

// CreateSubmodel creates a sub-model from the full model
func CreateSubmodel(model *types.Model, elements []*types.Element) types.Model {
	// Copy the model
	new := model.Copy()
	// Get a list of relevant elements
	relevantEls := getRelevantElements(model, elements)
	// Overwrite relationships with relevant ones
	new.Relationships = getRelevantRelationships(model, elements)
	// Remove irrelevant elements
	checkChildren(&new, relevantEls, &new.Root)
	return new
}

func checkChildren(model *types.Model, relevant map[*types.Element]bool, element *types.Element) {
	// Make a copy of the children array (we want to modify!)
	newSlice := make([]*types.Element, len(element.Children))
	copy(newSlice, element.Children)
	// Iterate
	for i, child := range newSlice {
		// Remove if irrelevant
		if !relevant[child] {
			remove(element.Children, i)
		}
		// Otherwise recurse
		checkChildren(model, relevant, child)
	}
}

// Select the relationships that are relevant, including implicit ones
func getRelevantRelationships(model *types.Model, elements []*types.Element) []types.Relationship {
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
func getRelevantElements(model *types.Model, elements []*types.Element) map[*types.Element]bool {
	// Prepare an empty index of relevant elements
	relevant := make(map[*types.Element]bool)
	// Add elements
	for _, el := range elements {
		// First, add the core element
		addAllAncestors(model, relevant, el)
	}
	return relevant
}

// Add the specified element to the map, and all its ancestors
func addAllAncestors(model *types.Model, elements map[*types.Element]bool, el *types.Element) map[*types.Element]bool {
	for {
		// Add the current element
		elements[el] = true
		// Try add parent
		if parent := el.Parent; parent == nil {
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

func remove(s []*types.Element, i int) []*types.Element {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}

// Check if an element is in the slice
func contains(haystack []*types.Element, needle *types.Element) bool {
	for _, el := range haystack {
		if el == needle {
			return true
		}
	}
	return false
}
