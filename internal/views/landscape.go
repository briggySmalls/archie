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
	return l.Model.Elements
}

func (l *Landscape) Relationships() []*types.Relationship {
	// Prepare return value
	var relationships []*types.Relationship
	// Iterate through the model's relationships
	for _, rel := range l.Model.Relationships {
		// Iterate through the view's elements
		for _, el := range l.Elements() {
			// Add relationships that include a relevant element
			if rel.Source == el || rel.Destination == el {
				relationships = append(relationships, rel)
			}
		}
	}
	return relationships
}
