package types

import (
	"fmt"
	"strings"
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
	m.AddElement(new, nil)
}

// Add an association between Elements
func (m *Model) AddAssociation(source, destination *Element) {
	// Append to relationships
	m.Associations = append(m.Associations, Relationship{source, destination})
}

func (m *Model) RootElements() []*Element {
	// Root is 'nil'
	return m.Children(nil)
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
		// Now link each of source's ancestors to destination
		for {
			// Link all source's ancestors to destination
			m.bubbleUpSource(relsMap, rel.Source, dest)
			// Iterate destination
			if parent := m.parent(dest); parent == nil {
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
		if parent == nil {
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

func (m *Model) IsAncestor(descendant, ancestor *Element) bool {
	for {
		// Check for a match
		if descendant == ancestor {
			return true
		}
		// Check if we're at the root
		if descendant == nil {
			return false
		}
		// Otherwise iterate
		descendant = m.parent(descendant)
	}
}

func (m *Model) LookupName(name string) (*Element, error) {
	// Split the string by slashes
	parts := strings.Split(name, "/")
	// Search down the tree
	var parent *Element
	parent = nil
NameLoop:
	for i, name := range parts {
		// Look for a child with the given name
		for _, el := range m.Children(parent) {
			if el.Name == name {
				// We've found the right child
				if i == len(parts)-1 {
					// We've found our element
					return el, nil
				}
				// Move on to the next name
				parent = el
				continue NameLoop
			}
		}
		// We didn't find a child matching that name
		if parent != nil {
			return nil, fmt.Errorf("Couldn't find child with name '%s' in '%s'", name, parent.Name)
		}
		return nil, fmt.Errorf("Couldn't find child with name '%s' in root", name)
	}
	panic(fmt.Errorf("It should be impossible to reach this code..."))
}

func (m *Model) bubbleUpSource(relationships map[Relationship]bool, source *Element, dest *Element) {
	for {
		if m.IsAncestor(dest, source) || m.IsAncestor(source, dest) {
			// We never link sub-items to their parents
			return
		}
		// Create the relationship
		relationships[Relationship{Source: source, Destination: dest}] = true
		// Iterate
		if parent := m.parent(source); parent == nil {
			// We've reached the root, we're done!
			return
		} else {
			// Iterate for the source's parent
			source = parent
		}
	}
}
