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
	oneChildChilda := types.NewItem("OneChildChilda")
	oneChildChildb := types.NewItem("OneChildChildb")
	two := types.NewItem("Two")
	twoChild := types.NewItem("TwoChild")
	twoChildChild := types.NewItem("TwoChildChild")

	// Add the items, and their relationships to the model
	m.AddRootElement(&one)
	m.AddElement(&oneChild, &one)
	m.AddElement(&oneChildChilda, &oneChild)
	m.AddElement(&oneChildChildb, &oneChild)
	m.AddRootElement(&two)
	m.AddElement(&twoChild, &two)
	m.AddElement(&twoChildChild, &twoChild)

	// Link the children together
	m.AddAssociation(&oneChildChilda, &twoChildChild)
	m.AddAssociation(&oneChildChilda, &oneChildChildb)

	// Create the view
	l := NewItemContextView(&m, &oneChild)

	// Check elements are correct
	assert.Assert(t, is.Contains(l.Elements, &one))
	assert.Assert(t, is.Contains(l.Elements, &oneChild))
	assert.Assert(t, is.Contains(l.Elements, &oneChildChilda))
	assert.Assert(t, is.Contains(l.Elements, &oneChildChildb))
	assert.Assert(t, is.Contains(l.Elements, &two))
	assert.Assert(t, is.Len(l.Elements, 5))
	assert.Assert(t, is.Len(l.Children(&two), 0))

	// Check relationships are correct
	assert.Assert(t, is.Contains(l.Associations, types.Relationship{Source: &oneChildChilda, Destination: &oneChildChildb}))
	assert.Assert(t, is.Contains(l.Associations, types.Relationship{Source: &oneChildChilda, Destination: &two}))
	assert.Assert(t, is.Len(l.Associations, 2))
}
