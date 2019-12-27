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

func NewAssociation(source, destination Element, tag string) Association {
	return association{source: source, destination: destination, tag: tag}
}

func (r association) Source() Element {
	return r.source
}

func (r association) Destination() Element {
	return r.destination
}

func (r association) Tag() string {
	return r.tag
}
