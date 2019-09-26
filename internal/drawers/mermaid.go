package drawers

import (
	"github.com/briggysmalls/archie/internal/types"
)

func NewMermaidDrawer() Drawer {
	// Create a new drawer with correct config
	d := newDrawer(MermaidConfig{})
	return &d
}

type MermaidConfig struct {
}

func (p MermaidConfig) Header(writer Writer) {
	writer.Write("graph TD")
	writer.UpdateIndent(1)
}

func (p MermaidConfig) Footer(writer Writer) {
	// Do nothing
}

func (p MermaidConfig) Element(writer Writer, element *types.Element) {
	writer.Write("%p(%s)", element, element.Name)
}

func (p MermaidConfig) StartParentElement(writer Writer, element *types.Element) {
	writer.Write("subgraph %s", element.Name)
}

func (p MermaidConfig) EndParentElement(writer Writer, element *types.Element) {
	writer.Write("end")
}

func (p MermaidConfig) Association(writer Writer, association types.Relationship) {
	writer.Write("%p-->%p", association.Source, association.Destination)
}
