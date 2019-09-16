package types

import (
	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
	"testing"
)

func TestElements(t *testing.T) {
	// Create a simple model
	m := Model{}

	// Create two items and add them to the model
	one := NewItem("One")
	two := NewItem("Two")
	m.AddElement(one)
	m.AddElement(two)

	// Assert
	assert.Assert(t, is.Contains(m.Elements, &one))
	assert.Assert(t, is.Contains(m.Elements, &two))
}

// Test implicit relationships
func TestImplicitRelationships(t *testing.T) {
	// Create a simple model
	m := Model{}

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

	// Link the children together
	m.AddRelationship(oneChildChild, twoChildChild)

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
