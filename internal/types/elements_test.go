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

// Test making children
func TestChildren(t *testing.T) {
	// Create some items
	first := NewItem("First")
	second := NewItem("Second")

	// Associate
	first.AddChild(&second)

	// Assert
	assert.Equal(t, first.Depth(), uint(0))
	assert.Equal(t, second.Depth(), uint(1))
	assert.Equal(t, second.Parent, &first)
}
