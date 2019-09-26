package views

import (
	"testing"

	mdl "github.com/briggysmalls/archie/core/model"
	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

// Test creating view for a scope with no children
func TestContextElements(t *testing.T) {
	// Create a simple model
	m, elMap := createModel()

	// Link the children together
	m.AddAssociation(elMap["OneChildChilda"], elMap["TwoChildChild"])
	m.AddAssociation(elMap["OneChildChilda"], elMap["OneChildChildb"])

	// Create the view
	l := NewContextView(m, elMap["OneChildChilda"])

	// Check elements are correct
	assert.Assert(t, is.Contains(l.Elements, elMap["One"]))
	assert.Assert(t, is.Contains(l.Elements, elMap["OneChild"]))
	assert.Assert(t, is.Contains(l.Elements, elMap["OneChildChilda"]))
	assert.Assert(t, is.Contains(l.Elements, elMap["OneChildChildb"]))
	assert.Assert(t, is.Contains(l.Elements, elMap["Two"]))
	assert.Assert(t, is.Len(l.Elements, 5))
	assert.Assert(t, is.Len(l.Children(elMap["Two"]), 0))

	// Check relationships are correct
	assert.Assert(t, is.Contains(l.Associations, mdl.Relationship{Source: elMap["OneChildChilda"], Destination: elMap["OneChildChildb"]}))
	assert.Assert(t, is.Contains(l.Associations, mdl.Relationship{Source: elMap["OneChildChilda"], Destination: elMap["Two"]}))
	assert.Assert(t, is.Len(l.Associations, 2))
}

// Test creating view for a scope with children
func TestContextChildElements(t *testing.T) {
	// Create a simple model
	m, elMap := createModel()

	// Link the children together
	m.AddAssociation(elMap["OneChildChilda"], elMap["TwoChildChild"])
	m.AddAssociation(elMap["OneChildChilda"], elMap["OneChildChildb"])

	// Create the view
	l := NewContextView(m, elMap["OneChild"])

	// Check elements are correct
	assert.Assert(t, is.Contains(l.Elements, elMap["One"]))
	assert.Assert(t, is.Contains(l.Elements, elMap["OneChild"]))
	assert.Assert(t, is.Contains(l.Elements, elMap["OneChildChilda"]))
	assert.Assert(t, is.Contains(l.Elements, elMap["OneChildChildb"]))
	assert.Assert(t, is.Contains(l.Elements, elMap["Two"]))
	assert.Assert(t, is.Len(l.Elements, 5))
	assert.Assert(t, is.Len(l.Children(elMap["Two"]), 0))

	// Check relationships are correct
	assert.Assert(t, is.Contains(l.Associations, mdl.Relationship{Source: elMap["OneChildChilda"], Destination: elMap["OneChildChildb"]}))
	assert.Assert(t, is.Contains(l.Associations, mdl.Relationship{Source: elMap["OneChildChilda"], Destination: elMap["Two"]}))
	assert.Assert(t, is.Len(l.Associations, 2))
}

// Helper function to create a model
func createModel() (*mdl.Model, map[string]*mdl.Element) {
	// Create a simple model
	m := mdl.NewModel()

	// Create the map
	elMap := make(map[string]*mdl.Element)

	// Create the items we'll be testing
	for _, name := range []string{"One", "OneChild", "OneChildChilda", "OneChildChildb", "Two", "TwoChild", "TwoChildChild"} {
		// Create the element
		el := mdl.NewItem(name)
		// Record it
		elMap[name] = &el
	}

	// Add the items to the model
	m.AddRootElement(elMap["One"])
	m.AddRootElement(elMap["Two"])
	m.AddElement(elMap["OneChild"], elMap["One"])
	m.AddElement(elMap["OneChildChilda"], elMap["OneChild"])
	m.AddElement(elMap["OneChildChildb"], elMap["OneChild"])
	m.AddElement(elMap["TwoChild"], elMap["Two"])
	m.AddElement(elMap["TwoChildChild"], elMap["TwoChild"])

	return &m, elMap
}
