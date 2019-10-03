package writers

import (
	mdl "github.com/briggysmalls/archie/internal/model"
)

type MermaidStrategy struct {
	linkAddress string
}

func (p MermaidStrategy) Header(scribe Scribe) {
	scribe.WriteLine("graph TD")
	scribe.UpdateIndent(1)
}

func (p MermaidStrategy) Footer(scribe Scribe) {
	// Do nothing
}

func (p MermaidStrategy) Element(scribe Scribe, element mdl.Element) {
	scribe.WriteLine("%p(%s)", element, element.Name())
}

func (p MermaidStrategy) StartParentElement(scribe Scribe, element mdl.Element) {
	scribe.WriteLine("subgraph %s", element.Name())
}

func (p MermaidStrategy) EndParentElement(scribe Scribe, element mdl.Element) {
	scribe.WriteLine("end")
}

func (p MermaidStrategy) Association(scribe Scribe, association mdl.Relationship) {
	scribe.WriteLine("%s-->%s", association.Source.ID(), association.Destination.ID())
}
