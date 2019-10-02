package writers

import (
	mdl "github.com/briggysmalls/archie/core/model"
	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
	"strings"
	"testing"
)

func TestDraw(t *testing.T) {
	// Create a simple model
	m := mdl.NewModel()

	// Create the items we'll be testing
	one := mdl.NewItem("One", "")
	oneChild := mdl.NewItem("OneChild", "")
	two := mdl.NewItem("Two", "")

	// Add the items, and their relationships to the model
	m.AddRootElement(&one)
	m.AddElement(&oneChild, &one)
	m.AddRootElement(&two)

	// Link the children together
	m.AddAssociation(&oneChild, &two)

	// Drawer
	d := New(PlantUmlStrategy{})
	output, err := d.Write(m)
	assert.NilError(t, err)

	// Assert result
	lines := strings.Split(output, "\n")
	assert.Equal(t, lines[0], "@startuml")
	// Find the line with the parent object definition
	var parentLine uint
	for i, line := range lines[1:] {
		if line == "package \"One\" {" {
			parentLine = uint(i + 1)
		}
	}
	assert.Equal(t, lines[parentLine], "package \"One\" {")
	assert.Equal(t, lines[parentLine+1], "    [OneChild]")
	assert.Equal(t, lines[parentLine+2], "}")
	assert.Assert(t, is.Contains(lines, "[Two]"))
	assert.Assert(t, is.Contains(lines, "[OneChild] -- [Two]"))
	assert.Equal(t, lines[len(lines)-2], "@enduml")
}
