package writers

import (
	"gotest.tools/assert"
	"testing"
)

func TestDrawPlantuml(t *testing.T) {
	// Create the test model
	m, elMap := createTestModel()

	// Drawer
	customFooter := `
skinparam shadowing false
skinparam nodesep 10
skinparam ranksep 20
`
	d := New(PlantUmlStrategy{CustomFooter: customFooter})
	output, err := d.Write(*m)
	assert.NilError(t, err)

	const resultFormat = `@startuml
package "One" <<software>> {
    rectangle "OneChild" as %[2]s
}
rectangle "Two" as %[3]s <<software>><<mechanical>>
actor "User" as %[4]s
%[2]s --> %[3]s

skinparam shadowing false
skinparam nodesep 10
skinparam ranksep 20
@enduml
`
	// Assert result
	assertOutput(t, output, resultFormat, elMap)
}
