package types

const (
	ACTOR = iota
	ITEM
	MODEL_ROOT
)

type Element struct {
	Name string
	Kind uint
}

// Create a new root element
func newModelRoot() Element {
	el := newElement(MODEL_ROOT)
	return el
}

// Create a new item
func NewItem(name string) Element {
	el := newElement(ITEM)
	el.Name = name
	return el
}

// Create a new actor
func NewActor(name string) Element {
	el := newElement(ACTOR)
	el.Name = name
	return el
}

// Create an element
func newElement(kind uint) Element {
	return Element{
		Kind: kind,
	}
}

// Internal helper to signify an element is a dummy/root element
func (e *Element) isRoot() bool {
	return e.Kind == MODEL_ROOT
}
