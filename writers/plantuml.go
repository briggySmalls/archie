package writers

import (
	"fmt"
)

type PlantUmlStrategy struct {
}

func (p PlantUmlStrategy) Header(scribe Scribe) {
	scribe.WriteLine("@startuml")
}

func (p PlantUmlStrategy) Footer(scribe Scribe) {
	scribe.WriteLine("skinparam shadowing false")
	scribe.WriteLine("skinparam nodesep 10")
	scribe.WriteLine("skinparam ranksep 20")
	scribe.WriteLine("@enduml")
}

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

func (p PlantUmlStrategy) StartParentElement(scribe Scribe, element Element) {
	scribe.WriteString(true, "package \"%s\"", element.Name())
	addStereotypes(scribe, element.Tags())
	scribe.WriteString(false, " {\n")
	scribe.UpdateIndent(1)
}

func (p PlantUmlStrategy) EndParentElement(scribe Scribe, element Element) {
	scribe.UpdateIndent(-1)
	scribe.WriteLine("}")
}

func (p PlantUmlStrategy) Association(scribe Scribe, association Relationship) {
	linkStr := fmt.Sprintf("%s --> %s", association.Source().ID(), association.Destination().ID())
	if association.Tag() != "" {
		scribe.WriteLine("%s : \"%s\"", linkStr, association.Tag())
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
