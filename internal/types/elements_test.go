package types

import (
	"testing"
)

// Test creating an item
func TestItem(t *testing.T) {
	// Create a test item
	item := NewItem("MyItem")

	// Verify it is not an actor
	if item.IsActor() {
		t.Error("Item identified as an actor")
	}

	// Verify name correct
	if item.name != "MyItem" {
		t.Errorf("Item name wrong (%s)", item.name)
	}
}

// Test creating an actor
func TestActor(t *testing.T) {
	// Create a test item
	item := NewActor("MyActor")

	// Verify it is not an actor
	if !item.IsActor() {
		t.Errorf("Actor not identified")
	}

	// Verify name correct
	if item.name != "MyActor" {
		t.Errorf("Item name wrong (%s)", item.name)
	}
}
