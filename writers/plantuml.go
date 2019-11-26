package writers

import (
	"fmt"
	"strings"
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
	for _, tag := range element.Tags() {
		scribe.WriteString(false, "<<%s>>", tag)
	}
	// Terminate with a newline
	scribe.WriteString(false, "\n")
}

func (p PlantUmlStrategy) StartParentElement(scribe Scribe, element Element) {
	scribe.WriteLine("package \"%s\" {", element.Name())
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

func buildStereotypes(tags []string) string {
	// Build up a list of tags (if present)
	var tagsBuilder strings.Builder
	for _, tag := range tags {
		_, err := tagsBuilder.WriteString(fmt.Sprintf("<<%s>>", tag))
		if err != nil {
			panic(err)
		}
	}
	return tagsBuilder.String()
}
