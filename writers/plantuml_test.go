package writers

import (
	"fmt"
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
	rectangle "OneChild" as %[1]s
}
rectangle "Two" as %[2]s <<software>><<mechanical>>
%[1]s --> %[2]s
skinparam shadowing false
skinparam nodesep 10
skinparam ranksep 20
@enduml
`
	// Assert result
	assert.Equal(t, output, fmt.Sprintf(resultFormat, elMap["OneChild"].ID(), elMap["Two"].ID()))
}
