package model

import (
	"testing"

	"gotest.tools/assert"
)

// Test creating an item
func TestItem(t *testing.T) {
	// Create a test item
	el := NewItem("MyItem")
	// Verify fields
	assert.Assert(t, el.Kind == item)
	assert.Equal(t, el.Name, "MyItem")
}

// Test creating an actor
func TestActor(t *testing.T) {
	// Create a test item
	el := NewActor("MyActor")
	// Verify fields
	assert.Assert(t, el.Kind == actor)
	assert.Equal(t, el.Name, "MyActor")
}
