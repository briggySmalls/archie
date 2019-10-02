package writers

import (
	"fmt"
	mdl "github.com/briggysmalls/archie/core/model"
	"net/url"
	"strings"
)

func NewMermaidDrawer(linkAddress string) Writer {
	// Create a new drawer with correct config
	d := New(MermaidStrategy{linkAddress: linkAddress})
	return &d
}

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

func (p MermaidStrategy) Element(scribe Scribe, element *mdl.Element) {
	scribe.WriteLine("%p(%s)", element, element.Name)
	// Add a link if necessary
	if p.linkAddress != "" {
		fullName, err := scribe.FullName(element)
		if err != nil {
			panic(err)
		}
		escapedName := url.PathEscape(fullName)
		escapedName = strings.Replace(escapedName, "%2F", "/", -1)
		url := fmt.Sprintf("%s%s", p.linkAddress, escapedName)
		scribe.WriteLine("click %s \"%s\"", element.ID(), url)
	}
}

func (p MermaidStrategy) StartParentElement(scribe Scribe, element *mdl.Element) {
	scribe.WriteLine("subgraph %s", element.Name)
}

func (p MermaidStrategy) EndParentElement(scribe Scribe, element *mdl.Element) {
	scribe.WriteLine("end")
}

func (p MermaidStrategy) Association(scribe Scribe, association mdl.Relationship) {
	scribe.WriteLine("%s-->%s", association.Source.ID(), association.Destination.ID())
}
