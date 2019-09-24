package types

import (
	"testing"

	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

func TestElements(t *testing.T) {
	// Create a simple model
	m := NewModel()

	// Create two items and add them to the model
	one := NewItem("One")
	two := NewItem("Two")
	m.AddRootElement(&one)
	m.AddRootElement(&two)

	// Assert
	assert.Assert(t, &one != nil)
	assert.Assert(t, &two != nil)
	assertName(t, &m, &one, "One")
	assertName(t, &m, &two, "Two")
}

// Test composition relationships
func TestComposition(t *testing.T) {
	// Create a simple model
	m, elMap := createModel()

	// Test parent results
	assertParent(t, m, elMap["One"], nil)
	assertParent(t, m, elMap["OneChild"], elMap["One"])
	assertParent(t, m, elMap["OneChildChilda"], elMap["OneChild"])
	assertParent(t, m, elMap["OneChildChildb"], elMap["OneChild"])
	assertParent(t, m, elMap["Two"], nil)
	assertParent(t, m, elMap["TwoChild"], elMap["Two"])
	assertParent(t, m, elMap["TwoChildChild"], elMap["TwoChild"])

	// Test children
	assertChildren(t, m, elMap["One"], []*Element{elMap["OneChild"]})
	assertChildren(t, m, elMap["OneChild"], []*Element{elMap["OneChildChilda"], elMap["OneChildChildb"]})
	assertChildren(t, m, elMap["Two"], []*Element{elMap["TwoChild"]})
	assertChildren(t, m, elMap["TwoChild"], []*Element{elMap["TwoChildChild"]})

	// Test lookups too
	assertName(t, m, elMap["One"], "One")
	assertName(t, m, elMap["OneChild"], "One/OneChild")
	assertName(t, m, elMap["OneChildChilda"], "One/OneChild/OneChildChilda")
	assertName(t, m, elMap["OneChildChildb"], "One/OneChild/OneChildChildb")
	assertName(t, m, elMap["Two"], "Two")
	assertName(t, m, elMap["TwoChild"], "Two/TwoChild")
	assertName(t, m, elMap["TwoChildChild"], "Two/TwoChild/TwoChildChild")
}

// Test trivial implicit relationships
func TestTrivialImplicitAssociations(t *testing.T) {
	// Create a simple model
	m := NewModel()

	// Create two items, each with one child
	one := NewItem("One")
	two := NewItem("Two")
	m.AddRootElement(&one)
	m.AddRootElement(&two)

	// Create a single relationship
	m.AddAssociation(&one, &two)

	// Assert implicit relationships returns trivial solution
	implicitRels := m.ImplicitAssociations()
	assert.Assert(t, is.Contains(implicitRels, Relationship{Source: &one, Destination: &two}))
	assert.Assert(t, is.Len(implicitRels, 1))
}

// Test implicit relationships
func TestDeepImplicitRelationships(t *testing.T) {
	// Create a simple model
	m, elMap := createModel()

	// Link the children together
	m.AddAssociation(elMap["OneChildChilda"], elMap["TwoChildChild"])
	m.AddAssociation(elMap["OneChildChilda"], elMap["OneChildChildb"])

	// Assert implicit relationships
	assert.Assert(t, is.Contains(m.Associations, Relationship{Source: elMap["OneChildChilda"], Destination: elMap["TwoChildChild"]}))
	assert.Assert(t, is.Contains(m.Associations, Relationship{Source: elMap["OneChildChilda"], Destination: elMap["OneChildChildb"]}))
	assert.Assert(t, is.Len(m.Associations, 2))
	implicitRels := m.ImplicitAssociations()
	// TODO: Check we never link a child to it's parent
	assert.Assert(t, is.Contains(implicitRels, Relationship{Source: elMap["One"], Destination: elMap["Two"]}))
	assert.Assert(t, is.Contains(implicitRels, Relationship{Source: elMap["One"], Destination: elMap["TwoChild"]}))
	assert.Assert(t, is.Contains(implicitRels, Relationship{Source: elMap["One"], Destination: elMap["TwoChildChild"]}))
	assert.Assert(t, is.Contains(implicitRels, Relationship{Source: elMap["OneChild"], Destination: elMap["Two"]}))
	assert.Assert(t, is.Contains(implicitRels, Relationship{Source: elMap["OneChild"], Destination: elMap["TwoChild"]}))
	assert.Assert(t, is.Contains(implicitRels, Relationship{Source: elMap["OneChild"], Destination: elMap["TwoChildChild"]}))
	assert.Assert(t, is.Contains(implicitRels, Relationship{Source: elMap["OneChildChilda"], Destination: elMap["Two"]}))
	assert.Assert(t, is.Contains(implicitRels, Relationship{Source: elMap["OneChildChilda"], Destination: elMap["TwoChild"]}))
	assert.Assert(t, is.Contains(implicitRels, Relationship{Source: elMap["OneChildChilda"], Destination: elMap["TwoChildChild"]}))
	assert.Assert(t, is.Len(implicitRels, 9))
}

// Helper function to assert expected parent
func assertParent(t *testing.T, m *Model, child *Element, parent *Element) {
	// Get the parent
	result, err := m.Parent(child)
	// Assert the lookup was successful
	assert.NilError(t, err)
	// Assert parent is as expected
	assert.Equal(t, parent, result)
}

func assertName(t *testing.T, m *Model, el *Element, name string) {
	// Try to look up the name
	result, err := m.LookupName(name)
	assert.NilError(t, err)
	// Now check they match
	assert.Equal(t, el, result)
}

func assertChildren(t *testing.T, m *Model, parent *Element, children []*Element) {
	result := m.Children(parent)
	assert.Assert(t, is.Len(result, len(children)))
	for _, expected := range children {
		assert.Assert(t, is.Contains(result, expected))
	}
}

// Helper function to create a model
func createModel() (*Model, map[string]*Element) {
	// Create a simple model
	m := NewModel()

	// Create the map
	elMap := make(map[string]*Element)

	// Create the items we'll be testing
	for _, name := range []string{"One", "OneChild", "OneChildChilda", "OneChildChildb", "Two", "TwoChild", "TwoChildChild"} {
		// Create the element
		el := NewItem(name)
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
