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
	m := NewModel()

	// Create the items we'll be testing
	one := NewItem("One")
	oneChild := NewItem("OneChild")
	oneChildChild := NewItem("OneChildChild")
	two := NewItem("Two")
	twoChild := NewItem("TwoChild")
	twoChildChild := NewItem("TwoChildChild")

	// Add the items, and their relationships to the model
	m.AddRootElement(&one)
	m.AddRootElement(&two)
	one.AddChild(&oneChild)
	oneChild.AddChild(&oneChildChild)
	two.AddChild(&twoChild)
	twoChild.AddChild(&twoChildChild)

	// Test parent results
	AssertParent(t, &m, &one, &m.Root)
	AssertParent(t, &m, &oneChild, &one)
	AssertParent(t, &m, &oneChildChild, &oneChild)
	AssertParent(t, &m, &two, &m.Root)
	AssertParent(t, &m, &twoChild, &two)
	AssertParent(t, &m, &twoChildChild, &twoChild)
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
	m := NewModel()

	// Create the items we'll be testing
	one := NewItem("One")
	oneChild := NewItem("OneChild")
	oneChildChild := NewItem("OneChildChild")
	two := NewItem("Two")
	twoChild := NewItem("TwoChild")
	twoChildChild := NewItem("TwoChildChild")

	// Add the items, and their relationships to the model
	m.AddRootElement(&one)
	m.AddRootElement(&two)
	one.AddChild(&oneChild)
	oneChild.AddChild(&oneChildChild)
	two.AddChild(&twoChild)
	twoChild.AddChild(&twoChildChild)

	// Link the children together
	m.AddRelationship(&oneChildChild, &twoChildChild)

	// Assert implicit relationships
	assert.Assert(t, is.Contains(m.Relationships, Relationship{Source: &oneChildChild, Destination: &twoChildChild}))
	assert.Assert(t, is.Len(m.Relationships, 1))
	implicitRels := m.ImplicitRelationships()
	assert.Assert(t, is.Contains(implicitRels, Relationship{Source: &one, Destination: &two}))
	assert.Assert(t, is.Contains(implicitRels, Relationship{Source: &one, Destination: &twoChild}))
	assert.Assert(t, is.Contains(implicitRels, Relationship{Source: &one, Destination: &twoChildChild}))
	assert.Assert(t, is.Contains(implicitRels, Relationship{Source: &oneChild, Destination: &two}))
	assert.Assert(t, is.Contains(implicitRels, Relationship{Source: &oneChild, Destination: &twoChild}))
	assert.Assert(t, is.Contains(implicitRels, Relationship{Source: &oneChild, Destination: &twoChildChild}))
	assert.Assert(t, is.Contains(implicitRels, Relationship{Source: &oneChildChild, Destination: &two}))
	assert.Assert(t, is.Contains(implicitRels, Relationship{Source: &oneChildChild, Destination: &twoChild}))
	assert.Assert(t, is.Contains(implicitRels, Relationship{Source: &oneChildChild, Destination: &twoChildChild}))
	assert.Assert(t, is.Len(implicitRels, 9))
}

// Helper function to assert expected parent
func AssertParent(t *testing.T, m *Model, child *Element, parent *Element) {
	// Assert parent is as expected
	assert.Equal(t, parent, child.Parent)
}
