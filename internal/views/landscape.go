package views

import (
	"github.com/briggysmalls/archie/internal/types"
)

type Landscape struct {
	Model types.Model
}

func NewLandscape(model types.Model) Landscape {
	// Create a landscape using the model
	return Landscape{model}
}

func (l *Landscape) Elements() []*types.Element {
	// All items of depth 0
	return l.Model.Elements()
}

func (l *Landscape) Relationships() []types.Relationship {
	// Get the view elements
	viewElements := l.Elements()
	// Prepare return value
	var relationships []types.Relationship
	// Iterate through _all_ of the the model's relationships
	for _, rel := range l.Model.ImplicitRelationships() {
		// Add relationships that link relevant elements
		if contains(viewElements, rel.Source) && contains(viewElements, rel.Destination) {
			relationships = append(relationships, rel)
		}
	}
	return relationships
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
