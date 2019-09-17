package types

type Relationship struct {
	Source      *Element
	Destination *Element
}

type Model struct {
	Elements      []*Element
	Relationships []Relationship
}

func (m *Model) AddElement(new *Element) error {
	// Ensure the element appears 'top'
	new.Parent = nil
	// Add to the model
	m.Elements = append(m.Elements, new)
	return nil
}

func (m *Model) AddRelationship(source *Element, destination *Element) error {
	// Append to relationships
	m.Relationships = append(m.Relationships, Relationship{source, destination})
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
			bubbleUpSource(relsMap, rel.Source, dest)
			// Iterate destination
			if dest.Parent == nil {
				break
			} else {
				dest = dest.Parent
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

func bubbleUpSource(relationships map[Relationship]bool, source *Element, dest *Element) {
	for {
		// Create the relationship
		relationships[Relationship{Source: source, Destination: dest}] = true
		if source.Parent == nil {
			break
		} else {
			// Update the pointer
			source = source.Parent
		}
	}
}
