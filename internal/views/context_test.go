package views

import (
	"testing"

	mdl "github.com/briggysmalls/archie/internal/model"
	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
	"strings"
)

// Test creating view for a scope with no children
func TestContextElements(t *testing.T) {
	// Create a simple model
	m, elMap := createModel()

	// Link the children together
	m.AddAssociation(elMap["1/1/1"], elMap["2/1/1"], []string{"t1"})
	m.AddAssociation(elMap["1/1/1"], elMap["2/1"], []string{"t2"})
	m.AddAssociation(elMap["1/1/1"], elMap["1/1/2"], nil)
	m.AddAssociation(elMap["1/1/1"], elMap["1/2/1"], nil)
	m.AddAssociation(elMap["1/1/2"], elMap["1/2/1"], nil)

	// Create the view
	l := NewContextView(m, elMap["1/1/1"])

	// Check elements are correct
	assertExpectedElements(t, l.Elements, []mdl.Element{
		elMap["1"],
		elMap["1/1"],
		elMap["1/2"],
		elMap["1/1/1"],
		elMap["1/1/2"],
		elMap["2"],
	})

	// Check children are correct
	assert.Assert(t, is.Len(l.Children(elMap["2"]), 0))
	assert.Assert(t, is.Len(l.Children(elMap["1/2"]), 0))

	// Check relationships are correct
	assert.Assert(t, is.Contains(l.Associations, mdl.NewAssociation(elMap["1/1/1"], elMap["1/1/2"], nil)))
	assert.Assert(t, is.Contains(l.Associations, mdl.NewAssociation(elMap["1/1/1"], elMap["2"], []string{"t1", "t2"})))
	assert.Assert(t, is.Contains(l.Associations, mdl.NewAssociation(elMap["1/1/1"], elMap["1/2"], nil)))
	assert.Assert(t, is.Len(l.Associations, 3))
}

// Test creating view for a scope with children
func TestContextChildElements(t *testing.T) {
	// Create a simple model
	m, elMap := createModel()

	// Link the children together
	m.AddAssociation(elMap["1/1/1"], elMap["2/1/1"], nil)
	m.AddAssociation(elMap["1/1/1"], elMap["1/1/2"], nil)
	m.AddAssociation(elMap["1/1/1"], elMap["1/2/1"], nil)

	// Create the view
	l := NewContextView(m, elMap["1/1"])

	// Check elements are correct
	assertExpectedElements(t, l.Elements, []mdl.Element{
		elMap["1"],
		elMap["1/1"],
		elMap["1/2"],
		elMap["1/1/1"],
		elMap["1/1/2"],
		elMap["2"],
	})

	// Check children are correct
	assert.Assert(t, is.Len(l.Elements, 6))
	assert.Assert(t, is.Len(l.Children(elMap["2"]), 0))
	assert.Assert(t, is.Len(l.Children(elMap["1/2"]), 0))

	// Check relationships are correct
	assert.Assert(t, is.Contains(l.Associations, mdl.NewAssociation(elMap["1/1/1"], elMap["1/1/2"], nil)))
	assert.Assert(t, is.Contains(l.Associations, mdl.NewAssociation(elMap["1/1/1"], elMap["2"], nil)))
	assert.Assert(t, is.Contains(l.Associations, mdl.NewAssociation(elMap["1/1/1"], elMap["1/2"], nil)))
	assert.Assert(t, is.Len(l.Associations, 3))
}

// Test creating view for landscape (context of nil)
func TestContextLandscape(t *testing.T) {
	// Create a simple model
	m, elMap := createModel()

	// Link the children together
	m.AddAssociation(elMap["1/1/1"], elMap["2/1/1"], nil)
	m.AddAssociation(elMap["1/1/1"], elMap["1/1/2"], nil)
	m.AddAssociation(elMap["1/1/1"], elMap["1/2/1"], nil)

	// Create the view
	l := NewContextView(m, nil)

	// Check elements are correct
	assertExpectedElements(t, l.Elements, []mdl.Element{
		elMap["1"],
		elMap["2"],
	})

	// Check children are correct
	assert.Assert(t, is.Len(l.Children(elMap["1"]), 0))
	assert.Assert(t, is.Len(l.Children(elMap["2"]), 0))

	// Check relationships are correct
	assert.Assert(t, is.Contains(l.Associations, mdl.NewAssociation(elMap["1"], elMap["2"], nil)))
	assert.Assert(t, is.Len(l.Associations, 1))
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
		"1/2",
		"1/1/1",
		"1/1/2",
		"1/2/1",
		"2",
		"2/1",
		"2/1/1",
	}
	for _, name := range elements {
		// Create the element
		el := mdl.NewItem(name, nil)
		// Record it
		elMap[name] = el
	}

	// Add the items to the model
	for name, el := range elMap {
		nesting := strings.Split(name, "/")
		if len(nesting) == 1 {
			m.AddRootElement(el)
		} else {
			m.AddElement(el, elMap[strings.Join(nesting[:len(nesting)-1], "/")])
		}
	}

	return &m, elMap
}

// Helper function to test expected elements
func assertExpectedElements(t *testing.T, expected, actual []mdl.Element) {
	// They must be the same length
	assert.Assert(t, is.Len(actual, len(expected)))
	// Now check each expected value is in the actual
	for _, el := range expected {
		// Assert elements are present
		assert.Assert(t, is.Contains(actual, el))
	}
}
