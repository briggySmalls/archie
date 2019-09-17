package types

const (
	ACTOR = iota
	ITEM
	MODEL
)

type Element struct {
	Name     string
	kind     uint
	Children []*Element
}

// Create a new item
func NewItem(name string) Element {
	return Element{
		Name: name,
		kind: ITEM,
	}
}

// Create a new actor
func NewActor(name string) Element {
	return Element{
		Name: name,
		kind: ACTOR,
	}
}

// Add a new element as a child of the specified element
func (e *Element) AddChild(new *Element) error {
	// Append the child
	e.Children = append(e.Children, new)
	return nil
}

// Helper function to determine if element is an actor
func (e *Element) IsActor() bool {
	return e.kind == ACTOR
}
