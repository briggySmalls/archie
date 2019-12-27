package model

type association struct {
	source      Element
	destination Element
	tag         string
}

// Association describes an association
type Association interface {
	Source() Element
	Destination() Element
	Tag() string
}

// NewAssociation creates an association between the specified elements, with the given tag
func NewAssociation(source, destination Element, tag string) Association {
	return association{source: source, destination: destination, tag: tag}
}

// Source gets the source element of the association
func (r association) Source() Element {
	return r.source
}

// Destination gets the destination element of the association
func (r association) Destination() Element {
	return r.destination
}

// Tag gets the tag given to the association
func (r association) Tag() string {
	return r.tag
}
