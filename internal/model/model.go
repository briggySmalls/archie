package model

import (
	"fmt"
	"sort"
	"strings"
)

// Model holds a fully defined Archie model for processing
type Model struct {
	Associations []Association
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
func (m *Model) AddAssociation(source, destination Element, tags []string) {
	// Append to associations
	m.Associations = append(m.Associations, NewAssociation(source, destination, tags))
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
	new.Associations = make([]Association, len(m.Associations))
	copy(new.Associations, m.Associations)
	// Composition
	new.Composition = make(map[Element]Element)
	for k, v := range m.Composition {
		new.Composition[k] = v
	}
	return new
}

// ImplicitAssociations gets a slice of all associations, including implicit ones.
// Implicit associations are those that link element parents of explicit associations.
func (m *Model) ImplicitAssociations() []Association {
	// Get all the associations
	rels := m.Associations
	// Prepare a list of implicit associations (we map to ensure no duplicates)
	relsMap := make(map[*Association]struct{})
	// Now add implicit associations
	for _, rel := range rels {
		dest := rel.Destination()
		// Now link each of source's ancestors to destination
		for {
			// Link all source's ancestors to destination
			m.bubbleUpSource(relsMap, rel.Source(), dest, rel.Tags())
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
	// Extract the associations from the map
	keys := make([]Association, len(relsMap))
	i := 0
	for ass := range relsMap {
		keys[i] = *ass
		i++
	}
	// Return the associations
	return keys
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
	// Get the elements
	for child, parent := range m.Composition {
		if parent == element {
			children = append(children, child)
		}
	}
	// Sort elements by name
	sort.Slice(children, func(i, j int) bool {
		return children[i].Name() < children[j].Name()
	})
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
func (m *Model) LookupName(name string, parent Element) (Element, error) {
	// Split the string by slashes
	parts := strings.Split(name, "/")
	// Search down the tree
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

func (m *Model) bubbleUpSource(associations map[*Association]struct{}, source Element, dest Element, tags []string) {
	for {
		if m.IsAncestor(dest, source) || m.IsAncestor(source, dest) {
			// We never link sub-items to their parents
			return
		}
		// Register that this association should be present
		newAss := NewAssociation(source, dest, tags)
		associations[&newAss] = struct{}{}
		// Iterate
		parent := m.parent(source)
		if parent == nil {
			// We've reached the root, we're done!
			return
		}

		// Iterate for the source's parent
		source = parent
	}
}
