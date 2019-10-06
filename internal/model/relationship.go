package model

type relationship struct {
	source      Element
	destination Element
}

type Relationship interface {
	Source() Element
	Destination() Element
}

func (r relationship) Source() Element {
	return r.source
}

func (r relationship) Destination() Element {
	return r.destination
}
