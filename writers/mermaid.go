package writers

// MermaidStrategy is the strategy for drawing a mermaid diagram from a model/view.
type MermaidStrategy struct {
}

// Header writes the header mermaid syntax.
func (p MermaidStrategy) Header(scribe Scribe) {
	scribe.WriteLine("graph TD")
	scribe.UpdateIndent(1)
}

// Footer writes a footer in mermaid syntax.
func (p MermaidStrategy) Footer(scribe Scribe) {
	// Do nothing
}

// Element writes an element in mermaid syntax.
func (p MermaidStrategy) Element(scribe Scribe, element Element) {
	scribe.WriteLine("%p(%s)", element, element.Name())
}

// StartParentElement writes the start of an enclosing/parent element in mermaid syntax.
func (p MermaidStrategy) StartParentElement(scribe Scribe, element Element) {
	scribe.WriteLine("subgraph %s", element.Name())
}

// EndParentElement writes the end of an enclosing/parent element in mermaid syntax.
func (p MermaidStrategy) EndParentElement(scribe Scribe, element Element) {
	scribe.WriteLine("end")
}

// Association writes an association in mermaid syntax
func (p MermaidStrategy) Association(scribe Scribe, association Association) {
	scribe.WriteLine("%s-->%s", association.Source().ID(), association.Destination().ID())
}
