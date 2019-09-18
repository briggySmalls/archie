package types

import (
	"testing"
	"gotest.tools/assert"
)

// Test that a ModelElement wrapping 'nil' is identified as a root node
func TestIsRoot(t *testing.T) {
	// Create a ModelElement, wrapping 'nil'
	me := NewModelElement(nil)
	// Assert we see a 'root'/virtual node
	assert.Assert(t, me.IsRoot())
}
