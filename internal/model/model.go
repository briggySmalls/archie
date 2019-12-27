package model

import (
	"fmt"
	"strings"
)

type relationshipNoTag struct {
	Source      Element
	Destination Element
}

// Model holds a fully defined Archie model for processing
type Model struct {
	Associations []Relationship
	Composition  map[Element]Element
	Elements     []Element
}

// NewModel creates and initialises new model
func NewModel() Model {
	// Create a model
	model := Model{
		Composition: make(map[Element]Element),
	}
	return model
}

// AddElement adds a new element to the model, as a child of the specified parent
func (m *Model) AddElement(new, parent Element) {
	// Add the new element
	m.Elements = append(m.Elements, new)
	// Associate the element with its parent
	m.Composition[new] = parent
}

// AddRootElement adds an element to the root of the model
func (m *Model) AddRootElement(new Element) {
	// Add element as a child of the root
	m.AddElement(new, nil)
}

// AddAssociation directionally associates the two specified elements
func (m *Model) AddAssociation(source, destination Element, tag string) {
	// Append to relationships
	m.Associations = append(m.Associations, NewRelationship(source, destination, tag))
}

// RootElements returns a slice of the root elements in a model
func (m *Model) RootElements() []Element {
	// Root is 'nil'
	return m.Children(nil)
}

// Copy safely duplicates a model
func (m *Model) Copy() Model {
	// Initially copy the struct
	new := *m
	// Deep copy the reference fields
	// Elements
	new.Elements = make([]Element, len(m.Elements))
	copy(new.Elements, m.Elements)
	// Associations
	new.Associations = make([]Relationship, len(m.Associations))
	copy(new.Associations, m.Associations)
	// Composition
	new.Composition = make(map[Element]Element)
	for k, v := range m.Composition {
		new.Composition[k] = v
	}
	return new
}

func (m *Model) ImplicitAssociations() []Relationship {
	// Get all the relationships
// ImplicitAssociations gets a slice of all associations, including implicit ones.
// Implicit associations are those that link element parents of explicit associations.
	rels := m.Associations
	// Prepare a list of implicit relationships (we map to ensure no duplicates)
	relsMap := make(map[relationshipNoTag]Relationship)
	// Now add implicit relationships
	for _, rel := range rels {
		dest := rel.Destination()
		// Now link each of source's ancestors to destination
		for {
			// Link all source's ancestors to destination
			m.bubbleUpSource(relsMap, rel.Source(), dest, rel.Tag())
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
	for _, v := range relsMap {
		keys[i] = v
		i++
	}
	// Return the relationships
	return keys
}

// Get the depth of an element
func (m *Model) Depth(el Element) (uint, error) {
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

func (m *Model) parent(element Element) Element {
	// Look up the element's parent
	element, err := m.Parent(element)
	// We use this function internally, so panic if we fail to find it
	if err != nil {
		panic(err)
	}
	return element
}

// Parent gets the parent of a specified element.
func (m *Model) Parent(element Element) (Element, error) {
	// Lookup
	element, ok := m.Composition[element]
	if !ok {
		return nil, fmt.Errorf("Element '%s' not found in model", element.Name())
	}
	return element, nil
}

// Children gets the children of a specified element.
func (m *Model) Children(element Element) []Element {
	var children []Element
	for child, parent := range m.Composition {
		if parent == element {
			children = append(children, child)
		}
	}
	return children
}

// IsAncestor indicates whether the queired element is a descendent of a specified element.
func (m *Model) IsAncestor(descendant, ancestor Element) bool {
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

// Name gets the fully namespaced name of an element.
// The namespaced name is of the form 'parent_name/child_name',
// i.e. the name of all anscestors with a '/' separator.
func (m *Model) Name(element Element) (string, error) {
	// Build full name for element
	parts := []string{}
	for {
		// Check if we're at the root
		if element == nil {
			break
		}
		// Prepend the name
		parts = append([]string{element.Name()}, parts...)
		// Iterate
		var err error
		element, err = m.Parent(element)
		if err != nil {
			return "", err
		}
	}
	return strings.Join(parts, "/"), nil
}

// ShareAncestor checks whether two elements have a common ancestor
func (m *Model) ShareAncestor(a, b Element) bool {
	// Find the respective root elements
	return m.getRoot(a) == m.getRoot(b)
}

func (m *Model) getRoot(element Element) Element {
	for {
		parent := m.parent(element)
		// Check if we've found the root
		if parent == nil {
			// We've found the root
			return element
		}
		// Iterate
		element = parent
	}
}

// LookupName returns a model element corresponding to the fully-namespaced name.
func (m *Model) LookupName(name string) (Element, error) {
	// Split the string by slashes
	parts := strings.Split(name, "/")
	// Search down the tree
	var parent Element
	parent = nil
NameLoop:
	for i, name := range parts {
		// Look for a child with the given name
		for _, el := range m.Children(parent) {
			if el.Name() == name {
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
			return nil, fmt.Errorf("Couldn't find child with name '%s' in '%s'", name, parent.Name())
		}
		return nil, fmt.Errorf("Couldn't find child with name '%s' in root", name)
	}
	panic(fmt.Errorf("It should be impossible to reach this code"))
}

func (m *Model) bubbleUpSource(relationships map[relationshipNoTag]Relationship, source Element, dest Element, tag string) {
	for {
		if m.IsAncestor(dest, source) || m.IsAncestor(source, dest) {
			// We never link sub-items to their parents
			return
		}
		// Create the relationship
		key := relationshipNoTag{Source: source, Destination: dest}
		if val, ok := relationships[key]; ok && tag != val.Tag() {
			// We have:
			// a) Already got an association with this source/dest pair
			// b) but it has a different tag to this new one
			// Indicate the tags are complex
			relationships[key] = NewRelationship(source, dest, "...")
		} else {
			// No previous source/dest implicit association
			relationships[key] = NewRelationship(source, dest, tag)
		}
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
