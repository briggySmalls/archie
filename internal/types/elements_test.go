package types

import (
	"testing"

	"gotest.tools/assert"
)

// Test creating an item
func TestItem(t *testing.T) {
	// Create a test item
	el := NewItem("MyItem")
	// Verify fields
	assert.Assert(t, el.Kind == ITEM)
	assert.Equal(t, el.Name, "MyItem")
}

// Test creating an actor
func TestActor(t *testing.T) {
	// Create a test item
	el := NewActor("MyActor")
	// Verify fields
	assert.Assert(t, el.Kind == ACTOR)
	assert.Equal(t, el.Name, "MyActor")
}

// Test creating a model root
func TestModelRoot(t *testing.T) {
	// Create a test item
	el := newModelRoot()
	// Verify it is not an actor
	assert.Assert(t, el.Kind == MODEL_ROOT)
}
