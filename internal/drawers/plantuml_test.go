package drawers

import (
	"fmt"
	"testing"
	"gotest.tools/assert"
	"github.com/briggysmalls/archie/internal/types"
	"github.com/briggysmalls/archie/internal/views"
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
	assert.NilError(t, m.AddChild(&one, &oneChild))
	m.AddRootElement(&two)
	assert.NilError(t, m.AddChild(&two, &twoChild))

	// Link the children together
	m.AddRelationship(&oneChild, &twoChild)

	// Create the landscape view
	l := views.NewLandscapeView(&m)

	// Create the drawer
	d := PlantUmlDrawer{}
	fmt.Print(d.Draw(l))
}
