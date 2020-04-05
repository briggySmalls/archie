package views

import (
	mdl "github.com/briggysmalls/archie/internal/model"
)

type srcAndDest struct {
	Source      mdl.Element
	Destination mdl.Element
}

// CreateSubmodel creates a sub-model from the full model
func CreateSubmodel(model *mdl.Model, primary, secondary []mdl.Element) (mdl.Model, error) {
	// Copy the model
	new := model.Copy()
	// Overwrite elements with relevant ones
	relevant, err := getRelevantElements(&new, append(primary, secondary...))
	if err != nil {
		return mdl.Model{}, err
	}
	new.Elements = relevant
	// Overwrite associations with relevant ones
	new.Associations = getRelevantAssociations(&new, primary, secondary)
	// Coalesce associations
	new.Associations = coalesceAssociations(new.Associations)
	// Fixup the composition relationships
	for child := range new.Composition {
		if !contains(new.Elements, child) {
			delete(new.Composition, child)
		}
	}
	return new, nil
}

// Select the elements that are relevant, including the implicit ones
func getRelevantElements(model *mdl.Model, elements []mdl.Element) ([]mdl.Element, error) {
	// Prepare an empty index of relevant elements
	// Note: We use a map just so elements are not duplicated
	relevant := make(map[mdl.Element]bool)
	for _, el := range elements {
		// Add the relevant element, and all its ancenstors
		err := addAllAncestors(model, relevant, el)
		if err != nil {
			return nil, err
		}
	}
	// Create a slice from the map keys
	keys := make([]mdl.Element, len(relevant))
	i := 0
	for k := range relevant {
		keys[i] = k
		i++
	}
	return keys, nil
}

// Select the relationships that are relevant, including implicit ones
func getRelevantAssociations(model *mdl.Model, primary, secondary []mdl.Element) []mdl.Association {
	// Union the primary and secondary elements
	allRelevant := append(primary, secondary...)
	var relationships []mdl.Association
	for _, rel := range model.ImplicitAssociations() {
		// Add relationships that link relevant elements
		sourcePrimary := contains(primary, rel.Source()) && contains(allRelevant, rel.Destination())
		destPrimary := contains(primary, rel.Destination()) && contains(allRelevant, rel.Source())
		if sourcePrimary || destPrimary {
			relationships = append(relationships, rel)
		}
	}
	return relationships
}

// Coalesce multiple associations between the same two items
func coalesceAssociations(associations []mdl.Association) []mdl.Association {
	// Group associations into bins of source/destination pair
	assMap := make(map[srcAndDest][]mdl.Association, len(associations))
	for _, ass := range associations {
		// Create a key
		key := srcAndDest{Source: ass.Source(), Destination: ass.Destination()}
		// Add this association to those
		assMap[key] = append(assMap[key], ass)
	}
	// Coalesce the associations for a given source/destination
	coalesced := make([]mdl.Association, 0, len(assMap))
	for pair, set := range assMap {
		// Add each distinct tag, for each tag in each association
		tagMap := make(map[string]struct{})
		for _, ass := range set {
			for _, t := range ass.Tags() {
				tagMap[t] = struct{}{}
			}
		}
		// Grab the tags
		var tags []string
		if len(tagMap) > 0 {
			tags = make([]string, 0, len(tagMap))
			for t := range tagMap {
				tags = append(tags, t)
			}
		}
		// Add the coalesced association
		coalesced = append(coalesced, mdl.NewAssociation(pair.Source, pair.Destination, tags))
	}
	return coalesced
}

// Add the specified element to the map, and all its ancestors
func addAllAncestors(model *mdl.Model, elements map[mdl.Element]bool, el mdl.Element) error {
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

// Check if an element is in the slice
func contains(haystack []mdl.Element, needle mdl.Element) bool {
	for _, el := range haystack {
		if el == needle {
			return true
		}
	}
	return false
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

// Get the linked element, if the specified element is in the relationship
func getLinked(relationship mdl.Association, element mdl.Element) mdl.Element {
	if relationship.Source() == element {
		return relationship.Destination()
	}
	if relationship.Destination() == element {
		return relationship.Source()
	}
	return nil
}
