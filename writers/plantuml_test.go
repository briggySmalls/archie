package writers

import (
	"fmt"
	mdl "github.com/briggysmalls/archie/internal/model"
	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
	"strings"
	"testing"
)

func TestDraw(t *testing.T) {
	// Create a simple model
	m := mdl.NewModel()

	// Create the items we'll be testing
	actor := mdl.NewActor("User")
	one := mdl.NewItem("One", []string{"software"})
	oneChild := mdl.NewItem("OneChild", nil)
	two := mdl.NewItem("Two", []string{"software", "mechanical"})

	// Add the items, and their relationships to the model
	m.AddRootElement(actor)
	m.AddRootElement(one)
	m.AddElement(oneChild, one)
	m.AddRootElement(two)

	// Link the children together
	m.AddAssociation(oneChild, two, "")

	// Drawer
	customFooter := `
skinparam shadowing false
skinparam nodesep 10
skinparam ranksep 20
`
	d := New(PlantUmlStrategy{CustomFooter: customFooter})
	output, err := d.Write(m)
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
	assert.Equal(t, lines[parentLine+1], fmt.Sprintf("    rectangle \"OneChild\" as %s", oneChild.ID()))
	assert.Equal(t, lines[parentLine+2], "}")
	assert.Assert(t, is.Contains(lines, fmt.Sprintf("actor \"User\" as %s", actor.ID())))
	assert.Assert(t, is.Contains(lines, fmt.Sprintf("rectangle \"Two\" as %s <<software>><<mechanical>>", two.ID())))
	assert.Assert(t, is.Contains(lines, fmt.Sprintf("%s --> %s", oneChild.ID(), two.ID())))

	assert.Assert(t, is.Contains(lines, "skinparam shadowing false"))
	assert.Assert(t, is.Contains(lines, "skinparam nodesep 10"))
	assert.Assert(t, is.Contains(lines, "skinparam ranksep 20"))
	assert.Equal(t, lines[len(lines)-2], "@enduml")
}
