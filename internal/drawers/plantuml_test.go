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
	one.AddChild(&oneChild)
	m.AddRootElement(&two)
	two.AddChild(&twoChild)

	// Link the children together
	m.AddRelationship(&oneChild, &twoChild)

	// Create the landscape view
	l := views.NewLandscapeView(&m)

	// Drawer
	d := PlantUmlDrawer{}
	output := d.Draw(l)

	// Assert result
	lines := strings.Split(output, "\n")
	assert.Equal(t, lines[0], "@startuml")
	assert.Assert(t, is.Contains(lines, "[One]"))
	assert.Assert(t, is.Contains(lines, "[Two]"))
	assert.Assert(t, is.Contains(lines, "[One] --> [Two]"))
	assert.Equal(t, lines[len(lines)-2], "@enduml")
}
