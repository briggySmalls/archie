package drawers

import (
	"fmt"
	"strings"

	"github.com/briggysmalls/archie/internal/types"
	"github.com/briggysmalls/archie/internal/views"
)

const (
	SPACES_IN_TAB = 4
)

type PlantUmlDrawer struct {
	indent  uint
	builder strings.Builder
}

func (p *PlantUmlDrawer) Draw(view views.View) string {
	// Reset the drawer
	p.indent = 0
	p.builder.Reset()
	// Add the header
	p.writeLine("@startuml")
	// Draw the elements, recursively
	for _, el := range view.Elements() {
		p.drawComponent(el)
	}
	// Now draw the relationships
	for _, rel := range view.Relationships() {
		p.writeLine("[%s] --> [%s]", rel.Source.Name, rel.Destination.Name)
	}
	// Write footer
	p.writeLine("@enduml")
	// Return result
	return p.builder.String()
}

func (p *PlantUmlDrawer) drawComponent(el *types.Element) {
	if len(el.Children) == 0 {
		// Write a simple component
		p.writeLine("[%s]", el.Name)
	} else {
		// Start a new package
		p.writeLine("package \"%s\" {", el.Name)
		p.indent++
		for _, child := range el.Children {
			// Recurse through children
			p.drawComponent(child)
		}
		p.indent--
		p.writeLine("}")
	}
}

func (p *PlantUmlDrawer) writeLine(format string, args ...interface{}) {
	p.builder.WriteString(fmt.Sprintf("%*s%s\n", p.indent*SPACES_IN_TAB, "", fmt.Sprintf(format, args...)))
}
