package writers

import (
	"fmt"
	"strings"
)

// PlantUmlStrategy is the strategy for drawing a PlantUML diagram from a model/view.
type PlantUmlStrategy struct {
	CustomFooter string
}

// Header writes the header PlantUML syntax.
func (p PlantUmlStrategy) Header(scribe Scribe) {
	scribe.WriteLine("@startuml")
}

// Footer writes a footer in PlantUML syntax.
func (p PlantUmlStrategy) Footer(scribe Scribe) {
	scribe.WriteString(true, p.CustomFooter)
	scribe.WriteLine("@enduml")
}

// Element writes an element in PlantUML syntax.
func (p PlantUmlStrategy) Element(scribe Scribe, element Element) {
	// Format actors with special style
	if element.IsActor() {
		scribe.WriteLine("actor \"%s\" as %s", element.Name(), element.ID())
		return
	}
	// Format all other items
	scribe.WriteString(true, "rectangle \"%s\" as %s", element.Name(), element.ID())
	// Add a list of tags as stereotypes (if present)
	addStereotypes(scribe, element.Tags())
	// Terminate with a newline
	scribe.WriteString(false, "\n")
}

// StartParentElement writes the start of an enclosing/parent element in PlantUML syntax.
func (p PlantUmlStrategy) StartParentElement(scribe Scribe, element Element) {
	scribe.WriteString(true, "package \"%s\"", element.Name())
	addStereotypes(scribe, element.Tags())
	scribe.WriteString(false, " {\n")
	scribe.UpdateIndent(1)
}

// EndParentElement writes the end of an enclosing/parent element in PlantUML syntax.
func (p PlantUmlStrategy) EndParentElement(scribe Scribe, element Element) {
	scribe.UpdateIndent(-1)
	scribe.WriteLine("}")
}

// Association writes an association in PlantUML syntax
func (p PlantUmlStrategy) Association(scribe Scribe, association Association) {
	linkStr := fmt.Sprintf("%s --> %s", association.Source().ID(), association.Destination().ID())
	if len(association.Tags()) > 0 {
		scribe.WriteLine("%s : \"%s\"", linkStr, strings.Join(association.Tags(), ", "))
	} else {
		scribe.WriteLine(linkStr)
	}
}

// Add a list of tags as stereotypes (if present)
func addStereotypes(scribe Scribe, tags []string) {
	if len(tags) == 0 {
		return
	}
	// Add a preceeding space
	scribe.WriteString(false, " ")
	// Now append tags as stereotypes
	for _, tag := range tags {
		scribe.WriteString(false, "<<%s>>", tag)
	}
}
