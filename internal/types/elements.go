package types

const (
	ACTOR = iota
	ITEM
)

type Element struct {
	Name     string
	kind     uint
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

// Helper function to determine if element is an actor
func (e *Element) IsActor() bool {
	return e.kind == ACTOR
}
