type Element interface {
	Name()
	IsActor() bool
}

type Model interface {
	AddElement(new Element, parent modelElement) (error)
	AddRelationship(source modelElement, destination modelElement) (error)
}

type modelElement struct {
	Id      uint
	Depth   uint
	Element Element
	Children []*modelElement
}

type relationship struct {
	Source *modelElement
	Destination *modelElement
}

type model struct {
	Elements      []*modelElement
	Relationships []*Relationship
}

// Add a new element as a child of the specified element
func (m *model) AddElement(new Element, parent modelElement) (error) {
	// Package the element up
	modelEl := modelElement{
		Depth: parent.Depth + 1,
		Element: new,
	}

	// Add to the parent
	parent.Children := append(parent.Children, &modelEl)

	return nil
}

func (m *model) AddRelationship(source modelElement, destination modelElement) (error) {
	// Append to relationships
	m.Relationships := append(m.Relationships, &Relationship{source, destination})

	return nil
}