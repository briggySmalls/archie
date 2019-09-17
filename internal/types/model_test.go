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
	assert.Assert(t, m.lookup(&one) != nil)
	assert.Assert(t, m.lookup(&two) != nil)
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
	assert.NilError(t, m.AddChild(&one, &oneChild))
	assert.NilError(t, m.AddChild(&oneChild, &oneChildChild))
	assert.NilError(t, m.AddChild(&two, &twoChild))
	assert.NilError(t, m.AddChild(&twoChild, &twoChildChild))

	// Test parent results
	AssertParent(t, m, &one, nil)
	AssertParent(t, m, &oneChild, &one)
	AssertParent(t, m, &oneChildChild, &oneChild)
	AssertParent(t, m, &two, nil)
	AssertParent(t, m, &twoChild, &two)
	AssertParent(t, m, &twoChildChild, &twoChild)
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
	assert.Assert(
		t,
		is.Contains(m.ImplicitRelationships(),
		Relationship{Source: m.lookup(&one), Destination: m.lookup(&two)}))
	assert.Assert(t, is.Len(m.ImplicitRelationships(), 1))
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
	assert.NilError(t, m.AddChild(&one, &oneChild))
	assert.NilError(t, m.AddChild(&oneChild, &oneChildChild))
	assert.NilError(t, m.AddChild(&two, &twoChild))
	assert.NilError(t, m.AddChild(&twoChild, &twoChildChild))

	// Link the children together
	m.AddRelationship(&oneChildChild, &twoChildChild)

	// Assert implicit relationships
	assert.Assert(t, is.Contains(m.Relationships, Relationship{Source: m.lookup(&oneChildChild), Destination: m.lookup(&twoChildChild)}))
	assert.Assert(t, is.Len(m.Relationships, 1))
	implicitRels := m.ImplicitRelationships()
	assert.Assert(t, is.Contains(implicitRels, Relationship{Source: m.lookup(&one), Destination: m.lookup(&two)}))
	assert.Assert(t, is.Contains(implicitRels, Relationship{Source: m.lookup(&one), Destination: m.lookup(&twoChild)}))
	assert.Assert(t, is.Contains(implicitRels, Relationship{Source: m.lookup(&one), Destination: m.lookup(&twoChildChild)}))
	assert.Assert(t, is.Contains(implicitRels, Relationship{Source: m.lookup(&oneChild), Destination: m.lookup(&two)}))
	assert.Assert(t, is.Contains(implicitRels, Relationship{Source: m.lookup(&oneChild), Destination: m.lookup(&twoChild)}))
	assert.Assert(t, is.Contains(implicitRels, Relationship{Source: m.lookup(&oneChild), Destination: m.lookup(&twoChildChild)}))
	assert.Assert(t, is.Contains(implicitRels, Relationship{Source: m.lookup(&oneChildChild), Destination: m.lookup(&two)}))
	assert.Assert(t, is.Contains(implicitRels, Relationship{Source: m.lookup(&oneChildChild), Destination: m.lookup(&twoChild)}))
	assert.Assert(t, is.Contains(implicitRels, Relationship{Source: m.lookup(&oneChildChild), Destination: m.lookup(&twoChildChild)}))
	assert.Assert(t, is.Len(implicitRels, 9))
}

// Helper function to assert expected parent
func AssertParent(t *testing.T, m Model, element *Element, parent *Element) {
	childModelElement := m.lookup(element)
	var parentModelElement *ModelElement
	if parent != nil {
		parentModelElement = m.lookup(parent)
	}
	result, err := m.Parent(childModelElement)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, parentModelElement, result)
}
