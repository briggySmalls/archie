package types

const (
	ACTOR = iota
	ITEM
	MODEL_ROOT
)

type Element struct {
	Name     string
	Kind     uint
	Children []*Element
	Parent   *Element
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

// Update the model to track an element as a child of another
func (e *Element) AddChild(child *Element) {
	// Record that new element has a parent
	child.Parent = e
	// Copy new element into the model (make it a child)
	e.Children = append(e.Children, child)
}

// Internal helper to signify an element is a dummy/root element
func (e *Element) isRoot() bool {
	return e.Kind == MODEL_ROOT
}
