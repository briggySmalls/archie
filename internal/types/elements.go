package types

const (
	ACTOR = iota
	ITEM
)

type Element struct {
	name   string
	kind   uint
	Parent *Element
}

// Create a new item
func NewItem(name string) Element {
	return Element{
		name: name,
		kind: ITEM,
	}
}

// Create a new actor
func NewActor(name string) Element {
	return Element{
		name: name,
		kind: ACTOR,
	}
}

// Add a new element as a child of the specified element
func (e *Element) AddChild(new *Element) error {
	// Set the element to point to its parent
	new.Parent = e
	return nil
}

// Helper function to determine if element is an actor
func (e *Element) IsActor() bool {
	return e.kind == ACTOR
}

// Get the depth of the element
func (e *Element) Depth() uint {
	if e.Parent == nil {
		// We are at the top level
		return 0
	}
	// We are 1 deeper than the parent
	return e.Parent.Depth() + 1
}
