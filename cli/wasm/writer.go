package main

import (
	"github.com/briggysmalls/archie/writers"
)

type strategy struct {
}

func (s strategy) Header(scribe writers.Scribe) {
	scribe.WriteLine("graph TD")
	scribe.UpdateIndent(1)
}

func (s strategy) Footer(scribe writers.Scribe) {
	// Do nothing
}

func (s strategy) Element(scribe writers.Scribe, element writers.Element) {
	scribe.WriteLine("id-%s(%s)", element.ID(), element.Name())
	// Also add a hyperlink
	scribe.WriteLine("click id-%s %s", element.ID(), "mermaidCallback")
}

func (s strategy) StartParentElement(scribe writers.Scribe, element writers.Element) {
	scribe.WriteLine("subgraph %s", element.Name())
}

func (s strategy) EndParentElement(scribe writers.Scribe, element writers.Element) {
	scribe.WriteLine("end")
}

func (s strategy) Association(scribe writers.Scribe, association writers.Relationship) {
	scribe.WriteLine("id-%s-->id-%s", association.Source().ID(), association.Destination().ID())
}
