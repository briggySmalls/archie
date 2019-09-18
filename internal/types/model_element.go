package types

// Wrapper for an element
type ModelElement struct {
	Data *Element // The wrapped element
	Children []*ModelElement // The children in the model
}

// Create a new ModelElement
func NewModelElement(e *Element) ModelElement {
	// Create a new model element
	return ModelElement{Data: e}
}

// Indicates if this element represents the root of the model (dummy)
func (e *ModelElement) IsRoot() bool {
	return e.Data == nil
}
