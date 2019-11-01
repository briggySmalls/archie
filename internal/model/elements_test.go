package model

import (
	"testing"

	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

// Test creating an item
func TestNewItem(t *testing.T) {
	// Create a test item
	el := NewItem("MyItem", []string{"electronics"})
	// Verify fields
	assert.Equal(t, el.Name(), "MyItem")
	assert.Assert(t, is.Len(el.Tags(), 1))
	assert.Assert(t, is.Contains(el.Tags(), "electronics"))
	assert.Assert(t, !el.IsActor())
}

// Test creating an actor
func TestNewActor(t *testing.T) {
	// Create a test item
	el := NewActor("MyActor")
	// Verify fields
	assert.Equal(t, el.Name(), "MyActor")
	assert.Assert(t, is.Len(el.Tags(), 0))
	assert.Assert(t, el.IsActor())
}
