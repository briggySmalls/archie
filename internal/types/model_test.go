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
}

// Test parent indexing
func TestParentIndexing(t *testing.T) {
	// Create a simple model
	m, elMap := createModel()

	// Test parent results
	assertParent(t, m, elMap["One"], &m.Root)
	assertParent(t, m, elMap["OneChild"], elMap["One"])
	assertParent(t, m, elMap["OneChildChild"], elMap["OneChild"])
	assertParent(t, m, elMap["Two"], &m.Root)
	assertParent(t, m, elMap["TwoChild"], elMap["Two"])
	assertParent(t, m, elMap["TwoChildChild"], elMap["TwoChild"])
}

// Test trivial implicit relationships
func TestTrivialImplicitRelationships(t *testing.T) {
	// Create a simple model
	m := NewModel()

	// Create two items, each with one child
	one := NewItem("One")
	two := NewItem("Two")
	m.AddRootElement(&one)
	m.AddRootElement(&two)

	// Create a single relationship
	m.AddRelationship(&one, &two)

	// Assert implicit relationships returns trivial solution
	implicitRels := m.ImplicitRelationships()
	assert.Assert(t, is.Contains(implicitRels, Relationship{Source: &one, Destination: &two}))
	assert.Assert(t, is.Len(implicitRels, 1))
}

// Test implicit relationships
func TestDeepImplicitRelationships(t *testing.T) {
	// Create a simple model
	m, elMap := createModel()

	// Link the children together
	m.AddRelationship(elMap["OneChildChild"], elMap["TwoChildChild"])

	// Assert implicit relationships
	assert.Assert(t, is.Contains(m.Relationships, Relationship{Source: elMap["OneChildChild"], Destination: elMap["TwoChildChild"]}))
	assert.Assert(t, is.Len(m.Relationships, 1))
	implicitRels := m.ImplicitRelationships()
	assert.Assert(t, is.Contains(implicitRels, Relationship{Source: elMap["One"], Destination: elMap["Two"]}))
	assert.Assert(t, is.Contains(implicitRels, Relationship{Source: elMap["One"], Destination: elMap["TwoChild"]}))
	assert.Assert(t, is.Contains(implicitRels, Relationship{Source: elMap["One"], Destination: elMap["TwoChildChild"]}))
	assert.Assert(t, is.Contains(implicitRels, Relationship{Source: elMap["OneChild"], Destination: elMap["Two"]}))
	assert.Assert(t, is.Contains(implicitRels, Relationship{Source: elMap["OneChild"], Destination: elMap["TwoChild"]}))
	assert.Assert(t, is.Contains(implicitRels, Relationship{Source: elMap["OneChild"], Destination: elMap["TwoChildChild"]}))
	assert.Assert(t, is.Contains(implicitRels, Relationship{Source: elMap["OneChildChild"], Destination: elMap["Two"]}))
	assert.Assert(t, is.Contains(implicitRels, Relationship{Source: elMap["OneChildChild"], Destination: elMap["TwoChild"]}))
	assert.Assert(t, is.Contains(implicitRels, Relationship{Source: elMap["OneChildChild"], Destination: elMap["TwoChildChild"]}))
	assert.Assert(t, is.Len(implicitRels, 9))
}

func TestCopy(t *testing.T) {
	// Create a simple model
	m, _ := createModel()

	// Create a copy
	copy, elMap := m.Copy()

	// Check that the models are similar
	checkChildrenAreEqual(t, elMap, &m.Root, &copy.Root)
}

func checkChildrenAreEqual(t *testing.T, elMap map[*Element]*Element, orig *Element, copy *Element) {
	// Check fields are equal
	assert.Equal(t, orig.Name, copy.Name)
	// Check mapping is correct
	assert.Equal(t, copy, elMap[orig])
	// Ensure we have matching number of children
	assert.Equal(t, len(orig.Children), len(copy.Children))
	// Look at children
	for i := range orig.Children {
		// Pull out the relevant child
		origChild := orig.Children[i]
		copyChild := copy.Children[i]
		// Check parent field is updated too
		assert.Equal(t, copyChild.Parent, copy)
		// Check the child's fields are the same
		assert.Equal(t, origChild.Name, copyChild.Name)
		// Now recurse to these element's children
		checkChildrenAreEqual(t, elMap, origChild, copyChild)
	}
}

// Helper function to assert expected parent
func assertParent(t *testing.T, m *Model, child *Element, parent *Element) {
	// Assert parent is as expected
	assert.Equal(t, parent, child.Parent)
}

// Helper function to create a model
func createModel() (*Model, map[string]*Element) {
	// Create a simple model
	m := NewModel()

	// Create the map
	elMap := make(map[string]*Element)

	// Create the items we'll be testing
	for _, name := range []string{"One", "OneChild", "OneChildChild", "Two", "TwoChild", "TwoChildChild"} {
		// Create the element
		el := NewItem(name)
		// Record it
		elMap[name] = &el
	}

	// Add the items to the model
	m.AddRootElement(elMap["One"])
	m.AddRootElement(elMap["Two"])
	elMap["One"].AddChild(elMap["OneChild"])
	elMap["OneChild"].AddChild(elMap["OneChildChild"])
	elMap["Two"].AddChild(elMap["TwoChild"])
	elMap["TwoChild"].AddChild(elMap["TwoChildChild"])

	return &m, elMap
}
