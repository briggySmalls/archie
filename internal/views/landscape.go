
type Landscape struct {
	Model Model
}

func NewLandscape(model Model) Landscape {
	// Create a landscape using the model
	return Landscape{model}
}

func (l *Landscape) Elements() {
	// All items of depth 0
	return l.Model.Elements
}

func (l *Landscape) Relationships() {
	// Get all the relationships connected to one of the elements
	els := l.Elements()
	var relationships []*Relationship
	for rel := range l.Model.Relationships {
		if rel.Source ==
	}
}