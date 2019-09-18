package views

import (
	"testing"

	"github.com/briggysmalls/archie/internal/types"
	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

func TestElements(t *testing.T) {
	// Create a simple model
	m := types.NewModel()

	// Create the items we'll be testing
	one := types.NewItem("One")
	oneChild := types.NewItem("OneChild")
	two := types.NewItem("Two")
	twoChild := types.NewItem("TwoChild")

	// Add the items, and their relationships to the model
	m.AddRootElement(&one)
	m.AddChild(&one, &oneChild)
	m.AddRootElement(&two)
	m.AddChild(&two, &twoChild)

	// Link the children together
	m.AddRelationship(&oneChild, &twoChild)

	// Create the landscape view
	l := NewLandscapeView(&m)

	// Check elements are correct
	assert.Assert(t, is.Len(l.Elements(), 2))
	assert.Assert(t, is.Contains(l.Elements(), &one))
	assert.Assert(t, is.Contains(l.Elements(), &two))

	// Check relationships are correct
	// assert.Assert(t, is.Contains(
	// 	l.Relationships,
	// 	types.Relationship{Source: one.Data, Destination: &two.Data}))
	assert.Assert(t, is.Len(l.Relationships, 1))
	rel := l.Relationships[0]
	assert.Equal(t, rel.Source.Data, &one)
	assert.Equal(t, rel.Destination.Data, &two)
}
