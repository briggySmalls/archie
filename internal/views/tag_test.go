package views

import (
	"testing"

	mdl "github.com/briggysmalls/archie/internal/model"
	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
	"strings"
)

// Test creating view for a scope with no children
func TestTagElements(t *testing.T) {
	// Create a simple model
	m, elMap := createTagModel()

	// Link the children together
	m.AddAssociation(elMap["1/1/2"], elMap["1/1/1"], "detect")
	m.AddAssociation(elMap["1/1/1"], elMap["1/2/1"], "start")
	m.AddAssociation(elMap["1/2/1"], elMap["1/2/2"], "actuate")
	m.AddAssociation(elMap["1/3"], elMap["1/1/1"], "listen")
	m.AddAssociation(elMap["1/1/2"], elMap["1/2/2"], "spin")

	// Create the view
	l := NewTagView(m, elMap["1"], "software")

	// Check elements are correct
	expectedElements := []mdl.Element{
		elMap["1"],
		elMap["1/1"],
		elMap["1/2"],
		elMap["1/3"],
		elMap["1/1/1"],
		elMap["1/1/2"],
		elMap["1/2/1"],
		elMap["1/2/2"],
	}
	for _, el := range expectedElements {
		// Assert elements are present
		assert.Assert(t, is.Contains(l.Elements, el))
	}
	assert.Assert(t, is.Len(l.Elements, len(expectedElements)))

	// Check relationships are correct
	expectedAssociations := []mdl.Relationship{
		mdl.NewRelationship(elMap["1/1/2"], elMap["1/1/1"], "detect"),
		mdl.NewRelationship(elMap["1/1/1"], elMap["1/2/1"], "start"),
		mdl.NewRelationship(elMap["1/2/1"], elMap["1/2/2"], "actuate"),
		mdl.NewRelationship(elMap["1/3"], elMap["1/1/1"], "listen"),
	}
	for _, el := range expectedAssociations {
		// Assert elements are present
		assert.Assert(t, is.Contains(l.Associations, el))
	}
	assert.Assert(t, is.Len(l.Associations, len(expectedAssociations)))
}

// Helper function to create a model
func createTagModel() (*mdl.Model, map[string]mdl.Element) {
	// Create a simple model
	m := mdl.NewModel()

	// Create the map
	elMap := make(map[string]mdl.Element)

	// Create the items we'll be testing
	elements := []struct {
		name string
		tags []string
	}{
		{"1", []string{}},
		{"1/1", []string{}},
		{"1/2", []string{}},
		{"1/1/1", []string{"software"}},
		{"1/1/2", []string{"electronics"}},
		{"1/1/2/1", []string{"passive"}},
		{"1/1/2/2", []string{"IC"}},
		{"1/2/1", []string{"software"}},
		{"1/2/1/1", []string{"driver"}},
		{"1/2/1/2", []string{"ui"}},
		{"1/2/2", []string{"electronics"}},
		{"1/3", []string{"software"}},
		{"2", []string{}},
		{"2/1", []string{}},
		{"2/1/1", []string{}},
	}
	for _, data := range elements {
		// Create the element
		el := mdl.NewItem(data.name, data.tags)
		// Record it
		elMap[data.name] = el
	}

	// Add the items to the model
	for name, el := range elMap {
		nesting := strings.Split(name, "/")
		if len(nesting) == 1 {
			m.AddRootElement(el)
		} else {
			m.AddElement(el, elMap[strings.Join(nesting[:len(nesting)-1], "/")])
		}
	}

	return &m, elMap
}
