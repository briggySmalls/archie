package writers

type MermaidStrategy struct {
}

func (p MermaidStrategy) Header(scribe Scribe) {
	scribe.WriteLine("graph TD")
	scribe.UpdateIndent(1)
}

func (p MermaidStrategy) Footer(scribe Scribe) {
	// Do nothing
}

func (p MermaidStrategy) Element(scribe Scribe, element Element) {
	scribe.WriteLine("%p(%s)", element, element.Name())
}

func (p MermaidStrategy) StartParentElement(scribe Scribe, element Element) {
	scribe.WriteLine("subgraph %s", element.Name())
}

func (p MermaidStrategy) EndParentElement(scribe Scribe, element Element) {
	scribe.WriteLine("end")
}

func (p MermaidStrategy) Association(scribe Scribe, association Relationship) {
	scribe.WriteLine("%s-->%s", association.Source().ID(), association.Destination().ID())
}
