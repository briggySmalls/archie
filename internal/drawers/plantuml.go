package drawers

import (
	"github.com/briggysmalls/archie/internal/types"
)

func NewPlantUmlDrawer() Drawer {
	// Create a new drawer with correct config
	d := newDrawer(PlantUmlConfig{})
	return &d
}

type PlantUmlConfig struct {
}

func (p PlantUmlConfig) Header(writer Writer) {
	writer.Write("@startuml")
}

func (p PlantUmlConfig) Footer(writer Writer) {
	writer.Write("@enduml")
}

func (p PlantUmlConfig) Element(writer Writer, element *types.Element) {
	writer.Write("[%s]", element.Name)
}

func (p PlantUmlConfig) StartParentElement(writer Writer, element *types.Element) {
	writer.Write("package \"%s\" {", element.Name)
	writer.UpdateIndent(1)
}

func (p PlantUmlConfig) EndParentElement(writer Writer, element *types.Element) {
	writer.UpdateIndent(-1)
	writer.Write("}")
}

func (p PlantUmlConfig) Association(writer Writer, association types.Relationship) {
	writer.Write("[%s] -- [%s]", association.Source.Name, association.Destination.Name)
}
