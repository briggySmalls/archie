package views

import (
	"testing"

	mdl "github.com/briggysmalls/archie/internal/model"
	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

// Test creating view for a scope with no children
func TestContextElements(t *testing.T) {
	// Create a simple model
	m, elMap := createModel()

	// Link the children together
	m.AddAssociation(elMap["1/1/1"], elMap["2/1/1"])
	m.AddAssociation(elMap["1/1/1"], elMap["1/1/2"])

	// Create the view
	l := NewContextView(m, elMap["1/1/1"])

	// Check elements are correct
	assert.Assert(t, is.Contains(l.Elements, elMap["1"]))
	assert.Assert(t, is.Contains(l.Elements, elMap["1/1"]))
	assert.Assert(t, is.Contains(l.Elements, elMap["1/1/1"]))
	assert.Assert(t, is.Contains(l.Elements, elMap["1/1/2"]))
	assert.Assert(t, is.Contains(l.Elements, elMap["2"]))
	assert.Assert(t, is.Len(l.Elements, 5))
	assert.Assert(t, is.Len(l.Children(elMap["2"]), 0))

	// Check relationships are correct
	assert.Assert(t, is.Contains(l.Associations, mdl.NewRelationship(elMap["1/1/1"], elMap["1/1/2"])))
	assert.Assert(t, is.Contains(l.Associations, mdl.NewRelationship(elMap["1/1/1"], elMap["2"])))
	assert.Assert(t, is.Len(l.Associations, 2))
}

// Test creating view for a scope with children
func TestContextChildElements(t *testing.T) {
	// Create a simple model
	m, elMap := createModel()

	// Link the children together
	m.AddAssociation(elMap["1/1/1"], elMap["2/1/1"])
	m.AddAssociation(elMap["1/1/1"], elMap["1/1/2"])

	// Create the view
	l := NewContextView(m, elMap["1/1"])

	// Check elements are correct
	assert.Assert(t, is.Contains(l.Elements, elMap["1"]))
	assert.Assert(t, is.Contains(l.Elements, elMap["1/1"]))
	assert.Assert(t, is.Contains(l.Elements, elMap["1/1/1"]))
	assert.Assert(t, is.Contains(l.Elements, elMap["1/1/2"]))
	assert.Assert(t, is.Contains(l.Elements, elMap["2"]))
	assert.Assert(t, is.Len(l.Elements, 5))
	assert.Assert(t, is.Len(l.Children(elMap["2"]), 0))

	// Check relationships are correct
	assert.Assert(t, is.Contains(l.Associations, mdl.NewRelationship(elMap["1/1/1"], elMap["1/1/2"])))
	assert.Assert(t, is.Contains(l.Associations, mdl.NewRelationship(elMap["1/1/1"], elMap["2"])))
	assert.Assert(t, is.Len(l.Associations, 2))
}

// Helper function to create a model
func createModel() (*mdl.Model, map[string]mdl.Element) {
	// Create a simple model
	m := mdl.NewModel()

	// Create the map
	elMap := make(map[string]mdl.Element)

	// Create the items we'll be testing
	elements := []string{
		"1",
		"1/1",
		"1/1/1",
		"1/1/2",
		"2",
		"2/1",
		"2/1/1",
	}
	for _, name := range elements {
		// Create the element
		el := mdl.NewItem(name, "")
		// Record it
		elMap[name] = el
	}

	// Add the items to the model
	m.AddRootElement(elMap["1"])
	m.AddRootElement(elMap["2"])
	m.AddElement(elMap["1/1"], elMap["1"])
	m.AddElement(elMap["1/1/1"], elMap["1/1"])
	m.AddElement(elMap["1/1/2"], elMap["1/1"])
	m.AddElement(elMap["2/1"], elMap["2"])
	m.AddElement(elMap["2/1/1"], elMap["2/1"])

	return &m, elMap
}
