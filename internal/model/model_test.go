package model

import (
	"testing"

	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

func TestElements(t *testing.T) {
	// Create a simple model
	m := NewModel()

	// Create two items and add them to the model
	one := NewItem("One", nil)
	two := NewItem("Two", nil)
	m.AddRootElement(one)
	m.AddRootElement(two)

	// Assert
	assert.Assert(t, one != nil)
	assert.Assert(t, two != nil)
	assertName(t, &m, one, "One")
	assertName(t, &m, two, "Two")
}

// Test composition associations
func TestParent(t *testing.T) {
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
}

func TestChildren(t *testing.T) {
	// Create a simple model
	m, elMap := createModel()

	// Test children
	assertChildren(t, m, elMap["One"], []Element{elMap["OneChild"]})
	assertChildren(t, m, elMap["OneChild"], []Element{elMap["OneChildChilda"], elMap["OneChildChildb"]})
	assertChildren(t, m, elMap["OneChildChilda"], []Element{})
	assertChildren(t, m, elMap["OneChildChildb"], []Element{})
	assertChildren(t, m, elMap["Two"], []Element{elMap["TwoChild"]})
	assertChildren(t, m, elMap["TwoChild"], []Element{elMap["TwoChildChild"]})
}

func TestLookupName(t *testing.T) {
	// Create a simple model
	m, elMap := createModel()

	// Test lookups too
	assertName(t, m, elMap["One"], "One")
	assertName(t, m, elMap["OneChild"], "One/OneChild")
	assertName(t, m, elMap["OneChildChilda"], "One/OneChild/OneChildChilda")
	assertName(t, m, elMap["OneChildChildb"], "One/OneChild/OneChildChildb")
	assertName(t, m, elMap["Two"], "Two")
	assertName(t, m, elMap["TwoChild"], "Two/TwoChild")
	assertName(t, m, elMap["TwoChildChild"], "Two/TwoChild/TwoChildChild")
}

func TestIsAnscestor(t *testing.T) {
	// Create a simple model
	m, elMap := createModel()

	// Test valid anscestors
	assert.Assert(t, m.IsAncestor(elMap["OneChild"], elMap["One"]))
	assert.Assert(t, m.IsAncestor(elMap["OneChildChilda"], elMap["OneChild"]))
	assert.Assert(t, m.IsAncestor(elMap["OneChildChilda"], elMap["One"]))
	assert.Assert(t, m.IsAncestor(elMap["OneChildChildb"], elMap["OneChild"]))
	assert.Assert(t, m.IsAncestor(elMap["OneChildChildb"], elMap["One"]))
	assert.Assert(t, m.IsAncestor(elMap["TwoChild"], elMap["Two"]))
	assert.Assert(t, m.IsAncestor(elMap["TwoChildChild"], elMap["Two"]))
	assert.Assert(t, m.IsAncestor(elMap["TwoChildChild"], elMap["TwoChild"]))

	// Test invalid anscestors
	assert.Assert(t, !m.IsAncestor(elMap["OneChild"], elMap["Two"]))
	assert.Assert(t, !m.IsAncestor(elMap["OneChildChilda"], elMap["TwoChild"]))
	assert.Assert(t, !m.IsAncestor(elMap["OneChildChilda"], elMap["Two"]))
	assert.Assert(t, !m.IsAncestor(elMap["OneChildChildb"], elMap["TwoChild"]))
	assert.Assert(t, !m.IsAncestor(elMap["OneChildChildb"], elMap["Two"]))
	assert.Assert(t, !m.IsAncestor(elMap["TwoChild"], elMap["One"]))
	assert.Assert(t, !m.IsAncestor(elMap["TwoChildChild"], elMap["One"]))
	assert.Assert(t, !m.IsAncestor(elMap["TwoChildChild"], elMap["OneChild"]))
}

// Test trivial implicit associations
func TestTrivialImplicitAssociations(t *testing.T) {
	// Create a simple model
	m := NewModel()

	// Create two items, each with one child
	one := NewItem("One", nil)
	two := NewItem("Two", nil)
	m.AddRootElement(one)
	m.AddRootElement(two)

	// Create a single association
	m.AddAssociation(one, two, []string{"force"})

	// Assert implicit associations returns trivial solution
	implicitRels := m.ImplicitAssociations()
	assert.Assert(t, is.Contains(implicitRels, association{source: one, destination: two, tags: []string{"force"}}))
	assert.Assert(t, is.Len(implicitRels, 1))
}

// Test implicit associations
func TestDeepImplicitAssociations(t *testing.T) {
	// Create a simple model
	m, elMap := createModel()

	// Link the children together
	m.AddAssociation(elMap["OneChildChilda"], elMap["TwoChildChild"], []string{"force"})
	m.AddAssociation(elMap["OneChildChilda"], elMap["OneChildChildb"], []string{"heat"})
	m.AddAssociation(elMap["OneChild"], elMap["TwoChild"], []string{"friction"})

	// Assert associations
	assert.Assert(t, is.Contains(m.Associations, NewAssociation(elMap["OneChildChilda"], elMap["TwoChildChild"], []string{"force"})))
	assert.Assert(t, is.Contains(m.Associations, NewAssociation(elMap["OneChildChilda"], elMap["OneChildChildb"], []string{"heat"})))
	assert.Assert(t, is.Contains(m.Associations, NewAssociation(elMap["OneChild"], elMap["TwoChild"], []string{"friction"})))
	assert.Assert(t, is.Len(m.Associations, 3))
	// Assert implicit associations
	implicitRels := m.ImplicitAssociations()
	// TODO: Check we never link a child to it's parent
	assert.Assert(t, is.Contains(implicitRels, NewAssociation(elMap["One"], elMap["Two"], []string{"friction"})))
	assert.Assert(t, is.Contains(implicitRels, NewAssociation(elMap["One"], elMap["Two"], []string{"force"})))
	assert.Assert(t, is.Contains(implicitRels, NewAssociation(elMap["One"], elMap["TwoChild"], []string{"friction"})))
	assert.Assert(t, is.Contains(implicitRels, NewAssociation(elMap["One"], elMap["TwoChild"], []string{"force"})))
	assert.Assert(t, is.Contains(implicitRels, NewAssociation(elMap["One"], elMap["TwoChildChild"], []string{"force"})))
	assert.Assert(t, is.Contains(implicitRels, NewAssociation(elMap["OneChild"], elMap["Two"], []string{"force"})))
	assert.Assert(t, is.Contains(implicitRels, NewAssociation(elMap["OneChild"], elMap["Two"], []string{"friction"})))
	assert.Assert(t, is.Contains(implicitRels, NewAssociation(elMap["OneChild"], elMap["TwoChild"], []string{"friction"})))
	assert.Assert(t, is.Contains(implicitRels, NewAssociation(elMap["OneChild"], elMap["TwoChild"], []string{"force"})))
	assert.Assert(t, is.Contains(implicitRels, NewAssociation(elMap["OneChild"], elMap["TwoChildChild"], []string{"force"})))
	assert.Assert(t, is.Contains(implicitRels, NewAssociation(elMap["OneChildChilda"], elMap["Two"], []string{"force"})))
	assert.Assert(t, is.Contains(implicitRels, NewAssociation(elMap["OneChildChilda"], elMap["TwoChild"], []string{"force"})))
	assert.Assert(t, is.Contains(implicitRels, NewAssociation(elMap["OneChildChilda"], elMap["TwoChildChild"], []string{"force"})))
	assert.Assert(t, is.Contains(implicitRels, NewAssociation(elMap["OneChildChilda"], elMap["OneChildChildb"], []string{"heat"})))
	assert.Assert(t, is.Len(implicitRels, 14))
}

// Helper function to assert expected parent
func assertParent(t *testing.T, m *Model, child Element, parent Element) {
	// Get the parent
	result, err := m.Parent(child)
	// Assert the lookup was successful
	assert.NilError(t, err)
	// Assert parent is as expected
	assert.Equal(t, parent, result)
}

func assertName(t *testing.T, m *Model, el Element, name string) {
	// Try to look up the name
	result, err := m.LookupName(name, nil)
	assert.NilError(t, err)
	// Now check they match
	assert.Equal(t, el, result)
	// Also check name function works too
	builtName, err := m.Name(el)
	assert.NilError(t, err)
	assert.Equal(t, builtName, name)
}

func assertChildren(t *testing.T, m *Model, parent Element, children []Element) {
	result := m.Children(parent)
	assert.Assert(t, is.Len(result, len(children)))
	for _, expected := range children {
		assert.Assert(t, is.Contains(result, expected))
	}
}

// Helper function to create a model
func createModel() (*Model, map[string]Element) {
	// Create a simple model
	m := NewModel()

	// Create the map
	elMap := make(map[string]Element)

	// Create the items we'll be testing
	for _, name := range []string{"One", "OneChild", "OneChildChilda", "OneChildChildb", "Two", "TwoChild", "TwoChildChild"} {
		// Create the element
		el := NewItem(name, nil)
		// Record it
		elMap[name] = el
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
