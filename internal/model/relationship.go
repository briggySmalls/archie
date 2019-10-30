package model

type relationship struct {
	source      Element
	destination Element
	tag string
}

type Relationship interface {
	Source() Element
	Destination() Element
	Tag() string
}

func NewRelationship(source, destination Element, tag string) Relationship {
	return relationship{source: source, destination: destination, tag: tag}
}

func (r relationship) Source() Element {
	return r.source
}

func (r relationship) Destination() Element {
	return r.destination
}

func (r relationship) Tag() string {
	return r.tag
}
