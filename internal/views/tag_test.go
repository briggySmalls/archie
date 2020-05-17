package views

import (
	"testing"

	mdl "github.com/briggysmalls/archie/internal/model"
	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
	"strings"
)

// Test creating view for a scope with no children
func TestTagWideScope(t *testing.T) {
	// Create a simple model
	m, elMap := createTagModel()

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
		elMap["2"],
		elMap["2/1"],
		elMap["2/1/1"],
		elMap["2/1/2"],
	}
	for _, el := range expectedElements {
		// Assert elements are present
		assert.Assert(t, is.Contains(l.Elements, el))
	}
	assert.Assert(t, is.Len(l.Elements, len(expectedElements)))

	// Check relationships are correct
	expectedAssociations := []mdl.Association{
		mdl.NewAssociation(elMap["1/1/2"], elMap["1/1/1"], []string{"detect"}),
		mdl.NewAssociation(elMap["1/1/1"], elMap["1/2/1"], []string{"start"}),
		mdl.NewAssociation(elMap["1/2/1"], elMap["1/2/2"], []string{"display"}),
		mdl.NewAssociation(elMap["1/3"], elMap["1/1/1"], []string{"listen"}),
		mdl.NewAssociation(elMap["2/1/1"], elMap["1/2/1"], []string{"receive"}),
		mdl.NewAssociation(elMap["1/2/1"], elMap["2/1/2"], []string{"send"}),
	}
	for _, ass := range expectedAssociations {
		// Assert elements are present
		assert.Assert(t, is.Contains(l.Associations, ass))
	}
	assert.Assert(t, is.Len(l.Associations, len(expectedAssociations)))
}

// Test creating view for a scope with no children
func TestTagNarrowScope(t *testing.T) {
	// Create a simple model
	m, elMap := createTagModel()

	// Create the view
	l := NewTagView(m, elMap["1/2/1"], "software")

	// Check elements are correct
	expectedElements := []mdl.Element{
		elMap["1"],
		elMap["1/1"],
		elMap["1/2"],
		elMap["1/1/1"],
		elMap["1/2/1"],
		elMap["1/2/1/1"],
		elMap["1/2/1/2"],
		elMap["1/2/2"],
		elMap["2"],
		elMap["2/1"],
		elMap["2/1/1"],
		elMap["2/1/2"],
	}
	for _, el := range expectedElements {
		// Assert elements are present
		assert.Assert(t, is.Contains(l.Elements, el))
	}
	assert.Assert(t, is.Len(l.Elements, len(expectedElements)))

	// Check relationships are correct
	expectedAssociations := []mdl.Association{
		mdl.NewAssociation(elMap["1/1/1"], elMap["1/2/1/1"], []string{"start"}),
		mdl.NewAssociation(elMap["1/2/1/2"], elMap["1/2/2"], []string{"display"}),
		mdl.NewAssociation(elMap["2/1/1"], elMap["1/2/1/1"], []string{"receive"}),
		mdl.NewAssociation(elMap["1/2/1/1"], elMap["2/1/2"], []string{"send"}),
	}
	for _, ass := range expectedAssociations {
		// Assert elements are present
		assert.Assert(t, is.Contains(l.Associations, ass))
	}
	assert.Assert(t, is.Len(l.Associations, len(expectedAssociations)))
}

// Test creating view for a scope with no children
func TestTagScopeNoChildren(t *testing.T) {
	// Create a simple model
	m, elMap := createTagModel()

	// Create the view
	l := NewTagView(m, elMap["1/2/1/1"], "software")

	// Check elements are correct
	expectedElements := []mdl.Element{
		elMap["1"],
		elMap["1/1"],
		elMap["1/1/1"],
		elMap["1/2/1"],
		elMap["1/2/1/1"],
		elMap["1/2/1/2"],
		elMap["2"],
		elMap["2/1"],
		elMap["2/1/1"],
		elMap["2/1/2"],
	}
	for _, el := range expectedElements {
		// Assert elements are present
		assert.Assert(t, is.Contains(l.Elements, el))
	}
	assert.Assert(t, is.Len(l.Elements, len(expectedElements)))

	// Check relationships are correct
	expectedAssociations := []mdl.Association{
		mdl.NewAssociation(elMap["1/1/1"], elMap["1/2/1/1"], []string{"start"}),
		mdl.NewAssociation(elMap["2/1/1"], elMap["1/2/1/1"], []string{"receive"}),
		mdl.NewAssociation(elMap["1/2/1/1"], elMap["2/1/2"], []string{"send"}),
	}
	for _, ass := range expectedAssociations {
		// Assert elements are present
		assert.Assert(t, is.Contains(l.Associations, ass))
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
		{"1/2/1/1", []string{"software", "driver"}},
		{"1/2/1/2", []string{"software", "ui"}},
		{"1/2/2", []string{"electronics"}},
		{"1/3", []string{"software"}},
		{"2", []string{}},
		{"2/1", []string{}},
		{"2/1/1", []string{}},
		{"2/1/2", []string{"software"}},
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

	// Link the children together
	m.AddAssociation(elMap["1/1/2"], elMap["1/1/1"], []string{"detect"})
	m.AddAssociation(elMap["1/1/1"], elMap["1/2/1/1"], []string{"start"})
	m.AddAssociation(elMap["1/2/1/2"], elMap["1/2/2"], []string{"display"})
	m.AddAssociation(elMap["1/3"], elMap["1/1/1"], []string{"listen"})
	m.AddAssociation(elMap["1/1/2"], elMap["1/2/2"], []string{"spin"})
	m.AddAssociation(elMap["2/1/1"], elMap["1/2/1/1"], []string{"receive"})
	m.AddAssociation(elMap["1/2/1/1"], elMap["2/1/2"], []string{"send"})

	return &m, elMap
}
