package types

import (
	"fmt"
)

type Relationship struct {
	Source      *Element
	Destination *Element
}

type Model struct {
	dummyElement  Element
	Relationships []Relationship
	parentMap     map[*Element]*Element
}

// NewModel creates an initialises new model
func NewModel() Model {
	m := Model{}
	m.dummyElement.kind = MODEL
	return m
}

func (m *Model) AddElement(new *Element) error {
	// Add to the model
	m.dummyElement.Children = append(m.dummyElement.Children, new)
	return nil
}

func (m *Model) AddRelationship(source *Element, destination *Element) error {
	// Append to relationships
	m.Relationships = append(m.Relationships, Relationship{source, destination})
	return nil
}

func (m *Model) Elements() []*Element {
	return m.dummyElement.Children
}

// Get a slice of all relationships, including implicit parent relationships
func (m *Model) ImplicitRelationships() []Relationship {
	// Build a map of parents
	// Get all the relationships
	rels := m.Relationships
	// Prepare a list of implicit relationships (we map to ensure no duplicates)
	relsMap := make(map[Relationship]bool)
	// Now add implicit relationships
	for _, rel := range rels {
		dest := rel.Destination
		// Now link each of source's anscestors to destination
		for {
			// Link all source's anscestors to destination
			m.bubbleUpSource(relsMap, rel.Source, dest)
			// Iterate destination
			if parent, err := m.Parent(dest); err != nil {
				panic(err)
			} else if parent == nil {
				break
			} else {
				dest = parent
			}
		}
	}
	// Extract the keys of the map
	keys := make([]Relationship, len(relsMap))
	i := 0
	for k := range relsMap {
		keys[i] = k
		i++
	}
	// Return the relationships
	return keys
}

func (m *Model) bubbleUpSource(relationships map[Relationship]bool, source *Element, dest *Element) {
	for {
		// Create the relationship
		relationships[Relationship{Source: source, Destination: dest}] = true
		// Iterate
		if parent, err := m.Parent(source); err != nil {
			panic(err)
		} else if parent == nil {
			break
		} else {
			// Update the pointer
			source = parent
		}
	}
}

func (m *Model) Parent(el *Element) (*Element, error) {
	// Index if necessary
	if len(m.parentMap) == 0 {
		// Create an empty map
		m.parentMap = make(map[*Element]*Element)
		// Index the tree
		m.indexChildren(&m.dummyElement)
	}
	// Fetch the parent from the index
	if parent, ok := m.parentMap[el]; ok {
		return parent, nil
	}
	return nil, fmt.Errorf("Element not found")
}

// Depth-first indexing of parents
func (m *Model) indexChildren(el *Element) {
	// Look at each child
	for _, child := range el.Children {
		// Add to map
		if el.kind == MODEL {
			m.parentMap[child] = nil
		} else {
			m.parentMap[child] = el
		}
		// Recurse
		m.indexChildren(child)
	}
}
