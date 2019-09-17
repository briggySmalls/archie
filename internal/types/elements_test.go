package types

import (
	"gotest.tools/assert"
	"testing"
)

// Test creating an item
func TestItem(t *testing.T) {
	// Create a test item
	el := NewItem("MyItem")
	// Verify it is not an actor
	assert.Assert(t, !el.IsActor())
	// Verify name correct
	assert.Equal(t, el.name, "MyItem")
}

// Test creating an actor
func TestActor(t *testing.T) {
	// Create a test item
	el := NewActor("MyActor")
	// Verify it is not an actor
	assert.Assert(t, el.IsActor())
	// Verify name correct
	assert.Equal(t, el.name, "MyActor")
}
