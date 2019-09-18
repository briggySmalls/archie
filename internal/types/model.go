package types

import (
	"fmt"
)

type Relationship struct {
	Source      *ModelElement
	Destination *ModelElement
}

type Model struct {
	root            ModelElement
	Relationships   []Relationship
	modelElementMap map[*Element]*ModelElement
	parentMap       map[*ModelElement]*ModelElement
}

// NewModel creates an initialises new model
func NewModel() Model {
	// Create a model
	return Model{
		modelElementMap: make(map[*Element]*ModelElement),
	}
}

// Get the root elements of the model
func (m *Model) Elements() []*ModelElement {
	return m.root.Children
}

// Add an element to the root of the model
func (m *Model) AddRootElement(new *Element) {
	// Add to the model
	me := m.addElement(new)
	// Make new ModelElement a child of the root
	m.root.Children = append(m.root.Children, me)
}

// Update the model to track an element as a child of another
func (m *Model) AddChild(parent, child *Element) error {
	// Get the parent by lookup
	pme, ok := m.modelElementMap[parent]
	if !ok {
		return fmt.Errorf("Parent '%s' not found in model", parent.Name)
	}
	// Add the child to the model
	cme := m.addElement(child)
	// Record that new element is a child
	pme.Children = append(pme.Children, cme)
	// Indicate everything was successful
	return nil
}

// Add a link between ModelElements representing two Elements
func (m *Model) AddRelationship(source *Element, destination *Element) error {
	// Find the corresponding model elements
	sourceModelElement, ok := m.modelElementMap[source]
	if !ok {
		return fmt.Errorf("Source element '%s' not found in model", source.Name)
	}
	destModelElement, ok := m.modelElementMap[destination]
	if !ok {
		return fmt.Errorf("Destination element '%s' not found in model", source.Name)
	}

	// Append to relationships
	m.Relationships = append(m.Relationships, Relationship{sourceModelElement, destModelElement})
	return nil
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

func (m *Model) bubbleUpSource(relationships map[Relationship]bool, source *ModelElement, dest *ModelElement) {
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

// Get the parent of a ModelElement
func (m *Model) Parent(el *ModelElement) (*ModelElement, error) {
	// Index if necessary
	if !m.isIndexed() {
		m.populateIndex()
	}
	// Fetch the parent from the index
	if parent, ok := m.parentMap[el]; ok {
		if parent.IsRoot() {
			return nil, nil
		}
		return parent, nil
	}
	return nil, fmt.Errorf("Element not found")
}

// Internal function for adding an element to the model
func (m *Model) addElement(new *Element) *ModelElement {
	// Wrap the element in a ModelElement
	me := NewModelElement(new)
	// Cache the Element/ModelElement mapping
	m.modelElementMap[new] = &me
	// Return the ModelElement
	return &me
}

// Helper function to check if the parents are indexed
func (m *Model) isIndexed() bool {
	return len(m.parentMap) != 0
}

// Populate the index
func (m *Model) populateIndex() {
	// Create an empty map
	m.parentMap = make(map[*ModelElement]*ModelElement)
	// Index the tree
	m.indexChildren(&m.root)
}

// Depth-first indexing of parents
func (m *Model) indexChildren(el *ModelElement) {
	// Look at each child
	for _, child := range el.Children {
		// Add to map
		if el.IsRoot() {
			m.parentMap[child] = &m.root
		} else {
			m.parentMap[child] = el
		}
		// Recurse
		m.indexChildren(child)
	}
}

// Helper function for looking up ModelElements from Elements
func (m *Model) lookup(el *Element) *ModelElement {
	if me, ok := m.modelElementMap[el]; ok {
		return me
	}
	panic(fmt.Errorf("Element '%s' not found in model", el.Name))
}
