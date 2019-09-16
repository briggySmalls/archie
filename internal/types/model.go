package types

type Relationship struct {
	Source      *Element
	Destination *Element
}

type Model struct {
	Elements      []*Element
	Relationships []*Relationship
}

func (m *Model) AddElement(new Element) error {
	// Ensure the element appears 'top'
	new.Parent = nil
	// Add to the model
	m.Elements = append(m.Elements, &new)
	return nil
}

func (m *Model) AddRelationship(source Element, destination Element) error {
	// Append to relationships
	m.Relationships = append(m.Relationships, &Relationship{&source, &destination})
	return nil
}
