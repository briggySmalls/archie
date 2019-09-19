package types

import (
	"fmt"
)

type Relationship struct {
	Source      *Element
	Destination *Element
}

type Model struct {
	Root          Element
	Relationships []Relationship
}

// NewModel creates an initialises new model
func NewModel() Model {
	// Create a model
	return Model{
		Root: newModelRoot(),
	}
}

// Get the root elements of the model
func (m *Model) Elements() []*Element {
	return m.Root.Children
}

// Add an element to the root of the model
func (m *Model) AddRootElement(new *Element) {
	// Record the parent
	new.Parent = &m.Root
	// Make new Element a child of the root
	m.Root.Children = append(m.Root.Children, new)
}

// Add a link between Elements representing two Elements
func (m *Model) AddRelationship(source, destination *Element) {
	// Append to relationships
	m.Relationships = append(m.Relationships, Relationship{source, destination})
}

// Get a slice of all relationships, including implicit parent relationships
func (m *Model) ImplicitRelationships() []Relationship {
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
			if parent := dest.Parent; m.IsRoot(parent) {
				// This is a root element, so bail
				break
			} else {
				// Now do the same for parent of destination
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

// Get the depth of an element
func (m *Model) Depth(el *Element) (uint, error) {
	// Bubble up, while counting
	depth := uint(0)
	for {
		// Get the parent of the element
		parent := el.Parent
		if m.IsRoot(parent) {
			// We're done!
			return depth, nil
		}
		// Keep bubblin'
		depth++
		el = parent
	}
}

func (m *Model) Copy() (Model, map[*Element]*Element) {
	// Create a new model
	new := NewModel()
	// Copy the original model into the new
	elMap := make(map[*Element]*Element)
	new.Root = *m.copyChildren(&m.Root, elMap)
	// Copy the relationships
	for _, rel := range m.Relationships {
		new.AddRelationship(elMap[rel.Source], elMap[rel.Destination])
	}
	return new, elMap
}

func (m *Model) copyChildren(el *Element, elMap map[*Element]*Element) *Element {
	// Create a new element copy
	new := *el
	// Record the mapping
	elMap[el] = &new
	// Reset the slice
	new.Children = make([]*Element, len(new.Children))
	// Recursively copy children
	for _, child := range el.Children {
		new.Children = append(new.Children, m.copyChildren(child, elMap))
	}
	return &new
}

func (m *Model) IsRoot(el *Element) bool {
	// First, check if the element itself is a root
	if !el.isRoot() {
		return false
	}
	if el != &m.Root {
		// Element is the root of a different model!?
		panic(fmt.Errorf("Unexpected root found"))
	}
	return true
}

func (m *Model) bubbleUpSource(relationships map[Relationship]bool, source *Element, dest *Element) {
	for {
		// Create the relationship
		relationships[Relationship{Source: source, Destination: dest}] = true
		// Iterate
		if parent := source.Parent; m.IsRoot(parent) {
			// We've reached the root, we're done!
			break
		} else {
			// Iterate for the source's parent
			source = parent
		}
	}
}
