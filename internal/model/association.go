package model

type association struct {
	source      Element
	destination Element
	tags        []string
}

// Association describes an association
type Association interface {
	Source() Element
	Destination() Element
	Tags() []string
}

// NewAssociation creates an association between the specified elements, with the given tag
func NewAssociation(source, destination Element, tags []string) Association {
	return association{source: source, destination: destination, tags: tags}
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
func (r association) Tags() []string {
	return r.tags
}
