package writers

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
	if element.IsActor() {
		scribe.WriteLine("actor \"%s\" as %s", element.Name(), element.ID())
	} else {
		scribe.WriteLine("rectangle \"%s\" as %s", element.Name(), element.ID())
	}
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
	scribe.WriteLine("%s --> %s", association.Source().ID(), association.Destination().ID())
}
