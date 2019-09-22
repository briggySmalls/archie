package views

import (
	"github.com/briggysmalls/archie/internal/types"
)

// CreateSubmodel creates a sub-model from the full model
func CreateSubmodel(model *types.Model, elements []*types.Element) (types.Model, error) {
	// Copy the model
	new := *model
	// Overwrite elements with relevant ones
	element, err := getRelevantElements(&new, elements)
	if err != nil {
		return types.Model{}, err
	}
	new.Elements = element
	// Overwrite relationships with relevant ones
	new.Associations = getRelevantRelationships(&new, elements)
	// Fixup the composition relationships
	for child := range new.Composition {
		if !contains(new.Elements, child) {
			delete(new.Composition, child)
		}
	}
	return new, nil
}

// Select the elements that are relevant, including the implicit ones
func getRelevantElements(model *types.Model, elements []*types.Element) ([]*types.Element, error) {
	// Prepare an empty index of relevant elements
	// Note: We use a map just so elements are not duplicated
	relevant := make(map[*types.Element]bool)
	for _, el := range elements {
		// Add the relevant element, and all its ancenstors
		err := addAllAncestors(model, relevant, el)
		if err != nil {
			return nil, err
		}
	}
	// Create a slice from the map keys
	keys := make([]*types.Element, len(relevant))
	i := 0
	for k := range relevant {
		keys[i] = k
		i++
	}
	return keys, nil
}

// Select the relationships that are relevant, including implicit ones
func getRelevantRelationships(model *types.Model, elements []*types.Element) []types.Relationship {
	var relationships []types.Relationship
	for _, rel := range model.ImplicitAssociations() {
		// Add relationships that link relevant elements
		if contains(elements, rel.Source) && contains(elements, rel.Destination) {
			relationships = append(relationships, rel)
		}
	}
	return relationships
}

// Add the specified element to the map, and all its ancestors
func addAllAncestors(model *types.Model, elements map[*types.Element]bool, el *types.Element) error {
	for {
		// Add the current element
		elements[el] = true
		// Find the parent
		parent, err := model.Parent(el)
		if err != nil {
			return err
		}
		// Try add parent
		if parent == nil {
			// Parent is root, we're done here
			return nil
		} else if _, ok := elements[parent]; ok {
			// Parent is already in map, someone got there first
			// Return early
			return nil
		} else {
			// Now recurse up through all the parents
			err := addAllAncestors(model, elements, parent)
			if err != nil {
				return err
			}
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
