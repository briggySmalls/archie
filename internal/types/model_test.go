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
	m.AddElement(&one)
	m.AddElement(&two)

	// Assert
	assert.Assert(t, is.Contains(m.Elements(), &one))
	assert.Assert(t, is.Contains(m.Elements(), &two))
}

// Test parent indexing
func TestParentIndexing(t *testing.T) {
	// Create a simple model
	m := NewModel()

	// Create two items, each with one child
	one := NewItem("One")
	oneChild := NewItem("OneChild")
	oneChildChild := NewItem("OneChildChild")
	one.AddChild(&oneChild)
	oneChild.AddChild(&oneChildChild)

	two := NewItem("Two")
	twoChild := NewItem("TwoChild")
	twoChildChild := NewItem("TwoChildChild")
	two.AddChild(&twoChild)
	twoChild.AddChild(&twoChildChild)

	// Add the items to the model
	m.AddElement(&one)
	m.AddElement(&two)

	// Test parent results
	AssertParent(t, m, &oneChildChild, &oneChild)
	AssertParent(t, m, &oneChild, &one)
	AssertParent(t, m, &one, nil)
	AssertParent(t, m, &twoChildChild, &twoChild)
	AssertParent(t, m, &twoChild, &two)
	AssertParent(t, m, &two, nil)
}

// Test trivial implicit relationships
func TestTrivialImplicitRelationships(t *testing.T) {
	// Create a simple model
	m := NewModel()

	// Create two items, each with one child
	one := NewItem("One")
	two := NewItem("Two")
	m.AddElement(&one)
	m.AddElement(&two)

	// Create a single relationship
	m.AddRelationship(&one, &two)

	// Assert implicit relationships returns trivial solution
	assert.Assert(t, is.Contains(m.ImplicitRelationships(), Relationship{Source: &one, Destination: &two}))
	assert.Assert(t, is.Len(m.ImplicitRelationships(), 1))
}

// Test implicit relationships
func TestDeepImplicitRelationships(t *testing.T) {
	// Create a simple model
	m := NewModel()

	// Create items with children
	one := NewItem("One")
	oneChild := NewItem("OneChild")
	oneChildChild := NewItem("OneChildChild")
	one.AddChild(&oneChild)
	oneChild.AddChild(&oneChildChild)

	two := NewItem("Two")
	twoChild := NewItem("TwoChild")
	twoChildChild := NewItem("TwoChildChild")
	two.AddChild(&twoChild)
	twoChild.AddChild(&twoChildChild)

	// Add items to the model
	m.AddElement(&one)
	m.AddElement(&two)

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
func AssertParent(t *testing.T, m Model, element *Element, parent *Element) {
	result, err := m.Parent(element)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, result, parent)
}
