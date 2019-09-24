package drawers

import (
	"github.com/briggysmalls/archie/internal/types"
	"github.com/briggysmalls/archie/internal/views"
	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
	"strings"
	"testing"
)

func TestDraw(t *testing.T) {
	// Create a simple model
	m := types.NewModel()

	// Create the items we'll be testing
	one := types.NewItem("One")
	oneChild := types.NewItem("OneChild")
	two := types.NewItem("Two")
	twoChild := types.NewItem("TwoChild")

	// Add the items, and their relationships to the model
	m.AddRootElement(&one)
	m.AddElement(&oneChild, &one)
	m.AddRootElement(&two)
	m.AddElement(&twoChild, &two)

	// Link the children together
	m.AddAssociation(&oneChild, &twoChild)

	// Create the landscape view
	l := views.NewItemContextView(&m, &one)

	// Drawer
	d := NewPlantUmlDrawer()
	output, err := d.Draw(l)
	assert.NilError(t, err)

	// Assert result
	lines := strings.Split(output, "\n")
	assert.Equal(t, lines[0], "@startuml")
	assert.Equal(t, lines[2], "    [OneChild]")
	assert.Equal(t, lines[1], "package \"One\" {")
	assert.Equal(t, lines[3], "}")
	assert.Assert(t, is.Contains(lines, "[Two]"))
	assert.Assert(t, is.Contains(lines, "[OneChild] -- [Two]"))
	assert.Equal(t, lines[len(lines)-2], "@enduml")
}
