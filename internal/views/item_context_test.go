package views

import (
	"testing"

	"github.com/briggysmalls/archie/internal/types"
	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

func TestItemContextElements(t *testing.T) {
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
	l := NewItemContextView(&m, &one)

	// Check elements are correct
	assert.Assert(t, is.Contains(l.Elements(), &one))
	assert.Assert(t, is.Contains(l.Elements(), &oneChild))
	assert.Assert(t, is.Contains(l.Elements(), &two))
	assert.Assert(t, is.Len(l.Elements(), 3))

	// Check relationships are correct
	assert.Assert(t, is.Len(l.Relationships, 1))
	rel := l.Relationships[0]
	assert.Equal(t, rel.Source, &oneChild)
	assert.Equal(t, rel.Destination, &two)
}
