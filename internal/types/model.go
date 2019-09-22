package types

import (
	"fmt"
)

const (
	ROOT_INDEX = 0
)

type Relationship struct {
	Source      *Element
	Destination *Element
}

type Model struct {
	Associations []Relationship
	Composition  map[*Element]*Element
	Elements     []*Element
}

// NewModel creates an initialises new model
func NewModel() Model {
	// Create a model
	model := Model{
		Composition: make(map[*Element]*Element),
	}
	// Create a root elment and add it
	root := newModelRoot()
	model.Elements = append(model.Elements, &root)
	return model
}

func (m *Model) AddElement(new, parent *Element) {
	// Add the new element
	m.Elements = append(m.Elements, new)
	// Associate the element with its parent
	m.Composition[new] = parent
}

// Add an element to the root of the model
func (m *Model) AddRootElement(new *Element) {
	// Add element as a child of the root
	m.AddElement(new, m.root())
}

// Add an association between Elements
func (m *Model) AddAssociation(source, destination *Element) {
	// Append to relationships
	m.Associations = append(m.Associations, Relationship{source, destination})
}

// Get a slice of all relationships, including implicit parent relationships
func (m *Model) ImplicitAssociations() []Relationship {
	// Get all the relationships
	rels := m.Associations
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
			if parent := m.parent(dest); m.IsRoot(parent) {
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
		parent, err := m.Parent(el)
		if err != nil {
			// Failed to find parent
			return 0, err
		}
		if m.IsRoot(parent) {
			// We're done!
			return depth, nil
		}
		// Keep bubblin'
		depth++
		el = parent
	}
}

func (m *Model) parent(element *Element) *Element {
	// Look up the element's parent
	element, err := m.Parent(element)
	// We use this function internally, so panic if we fail to find it
	if err != nil {
		panic(err)
	}
	return element
}

func (m *Model) Parent(element *Element) (*Element, error) {
	// Lookup
	element, ok := m.Composition[element]
	if !ok {
		return nil, fmt.Errorf("Element '%s' not found in model", element.Name)
	}
	return element, nil
}

func (m *Model) Children(element *Element) []*Element {
	var children []*Element
	for child, parent := range m.Composition {
		if parent == element {
			children = append(children, child)
		}
	}
	return children
}

func (m *Model) IsRoot(el *Element) bool {
	// First, check if the element itself is a root
	if !el.isRoot() {
		return false
	}
	if el != m.root() {
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
		if parent := m.parent(source); m.IsRoot(parent) {
			// We've reached the root, we're done!
			break
		} else {
			// Iterate for the source's parent
			source = parent
		}

	}
}

func (m *Model) root() *Element {
	return m.Elements[ROOT_INDEX]
}
