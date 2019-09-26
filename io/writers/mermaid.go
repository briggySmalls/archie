package drawers

import (
	"fmt"
	mdl "github.com/briggysmalls/archie/core/model"
	"net/url"
	"strings"
)

func NewMermaidDrawer(linkAddress string) Drawer {
	// Create a new drawer with correct config
	d := newDrawer(MermaidConfig{linkAddress: linkAddress})
	return &d
}

type MermaidConfig struct {
	linkAddress string
}

func (p MermaidConfig) Header(writer Writer) {
	writer.Write("graph TD")
	writer.UpdateIndent(1)
}

func (p MermaidConfig) Footer(writer Writer) {
	// Do nothing
}

func (p MermaidConfig) Element(writer Writer, element *mdl.Element) {
	writer.Write("%p(%s)", element, element.Name)
	// Add a link if necessary
	if p.linkAddress != "" {
		fullName, err := writer.FullName(element)
		if err != nil {
			panic(err)
		}
		escapedName := url.PathEscape(fullName)
		escapedName = strings.Replace(escapedName, "%2F", "/", -1)
		url := fmt.Sprintf("%s%s", p.linkAddress, escapedName)
		writer.Write("click %s \"%s\"", element.ID(), url)
	}
}

func (p MermaidConfig) StartParentElement(writer Writer, element *mdl.Element) {
	writer.Write("subgraph %s", element.Name)
}

func (p MermaidConfig) EndParentElement(writer Writer, element *mdl.Element) {
	writer.Write("end")
}

func (p MermaidConfig) Association(writer Writer, association mdl.Relationship) {
	writer.Write("%s-->%s", association.Source.ID(), association.Destination.ID())
}
