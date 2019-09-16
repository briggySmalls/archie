package types

import (
	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
	"testing"
)

// Test implicit relationships
func TestImplicitRelationships(t *testing.T) {
	// Create a simple model
	m := Model{}

	// Create two items, each with one child
	systemOne := NewItem("SystemOne")
	systemOneChild := NewItem("SystemOneChild")
	systemOne.AddChild(&systemOneChild)
	systemTwo := NewItem("SystemTwo")
	systemTwoChild := NewItem("SystemTwoChild")
	systemTwo.AddChild(&systemTwoChild)

	// Link the children together
	m.AddRelationship(systemOneChild, systemTwoChild)

	// Assert implicit relationships
	assert.Assert(t, is.Len(m.Relationships, 1))
	implicitRels := m.ImplicitRelationships()
	assert.Assert(t, is.Len(implicitRels, 3))
}
