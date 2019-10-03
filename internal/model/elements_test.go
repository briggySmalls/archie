package model

import (
	"testing"

	"gotest.tools/assert"
)

// Test creating an item
func TestNewItem(t *testing.T) {
	// Create a test item
	el := NewItem("MyItem", "electronics")
	// Verify fields
	assert.Equal(t, el.Name(), "MyItem")
	assert.Equal(t, el.Technology(), "electronics")
	assert.Assert(t, !el.IsActor())
}

// Test creating an actor
func TestNewActor(t *testing.T) {
	// Create a test item
	el := NewActor("MyActor")
	// Verify fields
	assert.Equal(t, el.Name(), "MyActor")
	assert.Equal(t, el.Technology(), "")
	assert.Assert(t, el.IsActor())
}
