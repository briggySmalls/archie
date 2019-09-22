package drawers

import (
	"fmt"
	"strings"

	"github.com/briggysmalls/archie/internal/types"
)

const (
	SPACES_IN_TAB = 4
)

type PlantUmlDrawer struct {
	indent  uint
	builder strings.Builder
}

func (p *PlantUmlDrawer) Draw(model types.Model) string {
	// Reset the drawer
	p.indent = 0
	p.builder.Reset()
	// Add the header
	p.writeLine("@startuml")
	// Draw the elements, recursively
	for _, el := range model.Children(nil) {
		p.drawComponent(&model, el)
	}
	// Now draw the relationships
	for _, rel := range model.Associations {
		p.writeLine("[%s] --> [%s]", rel.Source.Name, rel.Destination.Name)
	}
	// Write footer
	p.writeLine("@enduml")
	// Return result
	return p.builder.String()
}

func (p *PlantUmlDrawer) drawComponent(model *types.Model, el *types.Element) {
	children := model.Children(el)
	if len(children) == 0 {
		// Write a simple component
		p.writeLine("[%s]", el.Name)
	} else {
		// Start a new package
		p.writeLine("package \"%s\" {", el.Name)
		p.indent++
		for _, child := range children {
			// Recurse through children
			p.drawComponent(model, child)
		}
		p.indent--
		p.writeLine("}")
	}
}

func (p *PlantUmlDrawer) writeLine(format string, args ...interface{}) {
	p.builder.WriteString(fmt.Sprintf("%*s%s\n", p.indent*SPACES_IN_TAB, "", fmt.Sprintf(format, args...)))
}
