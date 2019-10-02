package writers

import (
	mdl "github.com/briggysmalls/archie/core/model"
)

type PlantUmlStrategy struct {
}

func (p PlantUmlStrategy) Header(scribe Scribe) {
	scribe.WriteLine("@startuml")
}

func (p PlantUmlStrategy) Footer(scribe Scribe) {
	scribe.WriteLine("@enduml")
}

func (p PlantUmlStrategy) Element(scribe Scribe, element *mdl.Element) {
	scribe.WriteLine("[%s]", element.Name)
}

func (p PlantUmlStrategy) StartParentElement(scribe Scribe, element *mdl.Element) {
	scribe.WriteLine("package \"%s\" {", element.Name)
	scribe.UpdateIndent(1)
}

func (p PlantUmlStrategy) EndParentElement(scribe Scribe, element *mdl.Element) {
	scribe.UpdateIndent(-1)
	scribe.WriteLine("}")
}

func (p PlantUmlStrategy) Association(scribe Scribe, association mdl.Relationship) {
	scribe.WriteLine("[%s] -- [%s]", association.Source.Name, association.Destination.Name)
}
