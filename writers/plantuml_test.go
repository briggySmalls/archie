package writers

import (
	"fmt"
	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
	"strings"
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

	// Assert result
	lines := strings.Split(output, "\n")
	assert.Equal(t, lines[0], "@startuml")
	// Find the line with the parent object definition
	parentString := "package \"One\" <<software>> {"
	assert.Assert(t, is.Contains(lines, parentString))
	var parentLine uint
	for i, line := range lines[1:] {
		if line == parentString {
			parentLine = uint(i + 1)
		}
	}
	assert.Equal(t, lines[parentLine+1], fmt.Sprintf("    rectangle \"OneChild\" as %s", elMap["OneChild"].ID()))
	assert.Equal(t, lines[parentLine+2], "}")
	assert.Assert(t, is.Contains(lines, fmt.Sprintf("actor \"User\" as %s", elMap["User"].ID())))
	assert.Assert(t, is.Contains(lines, fmt.Sprintf("rectangle \"Two\" as %s <<software>><<mechanical>>", elMap["Two"].ID())))
	assert.Assert(t, is.Contains(lines, fmt.Sprintf("%s --> %s", elMap["OneChild"].ID(), elMap["Two"].ID())))

	assert.Assert(t, is.Contains(lines, "skinparam shadowing false"))
	assert.Assert(t, is.Contains(lines, "skinparam nodesep 10"))
	assert.Assert(t, is.Contains(lines, "skinparam ranksep 20"))
	assert.Equal(t, lines[len(lines)-2], "@enduml")
}
