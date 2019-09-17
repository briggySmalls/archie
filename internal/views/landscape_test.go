package views

import (
	"github.com/briggysmalls/archie/internal/types"
	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
	"testing"
)

func TestElements(t *testing.T) {
	// Create a simple model
	m := types.Model{}

	// Create two items, each with one child
	one := types.NewItem("SystemOne")
	oneChild := types.NewItem("SystemOneChild")
	one.AddChild(&oneChild)
	two := types.NewItem("SystemTwo")
	twoChild := types.NewItem("SystemTwoChild")
	two.AddChild(&twoChild)

	// Add to model
	m.AddElement(&one)
	m.AddElement(&two)

	// Create the landscape view
	l := NewLandscape(m)
	// Check elements are correct
	assert.Assert(t, is.Len(l.Elements(), 2))
	assert.Assert(t, is.Contains(l.Elements(), &one))
	assert.Assert(t, is.Contains(l.Elements(), &two))
}

func TestRelationships(t *testing.T) {
	// Create a simple model
	m := types.Model{}

	// Create two items, each with one child
	one := types.NewItem("SystemOne")
	oneChild := types.NewItem("SystemOneChild")
	one.AddChild(&oneChild)
	two := types.NewItem("SystemTwo")
	twoChild := types.NewItem("SystemTwoChild")
	two.AddChild(&twoChild)

	// Add to model
	m.AddElement(&one)
	m.AddElement(&two)

	// Add relationship
	m.AddRelationship(&one, &two)

	// Create the landscape view
	l := NewLandscape(m)

	// Check relationships are correct
	assert.Assert(t, is.Contains(l.Relationships(), types.Relationship{Source: &one, Destination: &two}))
	assert.Assert(t, is.Len(l.Relationships(), 1))
	rel := l.Relationships()[0]
	assert.Equal(t, rel.Source, &one)
	assert.Equal(t, rel.Destination, &two)

}
