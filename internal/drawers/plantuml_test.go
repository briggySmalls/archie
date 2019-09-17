package drawers

import (
	"fmt"
	"testing"

	"github.com/briggysmalls/archie/internal/types"
	"github.com/briggysmalls/archie/internal/views"
)

func TestDraw(t *testing.T) {
	// Create a simple model
	m := types.NewModel()

	// Create two items, each with one child
	one := types.NewItem("SystemOne")
	oneChild := types.NewItem("SystemOneChild")
	one.AddChild(&oneChild)
	two := types.NewItem("SystemTwo")
	twoChild := types.NewItem("SystemTwoChild")
	two.AddChild(&twoChild)

	// Add to model
	m.AddElement(&one)
	m.AddElement(&two)

	// Add relationship
	m.AddRelationship(&one, &two)

	// Create the landscape view
	l := views.NewLandscape(m)

	// Create the drawer
	d := PlantUmlDrawer{}
	fmt.Print(d.Draw(&l))
}
