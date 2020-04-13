package writers

import (
	"fmt"
	mdl "github.com/briggysmalls/archie/internal/model"
	"gotest.tools/v3/assert"
	"sort"
	"testing"
)

func createTestModel() (*mdl.Model, map[string]mdl.Element) {
	// Create a simple model
	m := mdl.NewModel()

	// Create the items we'll be testing
	actor := mdl.NewActor("User")
	one := mdl.NewItem("One", []string{"software"})
	oneChild := mdl.NewItem("OneChild", nil)
	two := mdl.NewItem("Two", []string{"software", "mechanical"})

	// Add the items, and their relationships to the model
	m.AddRootElement(actor)
	m.AddRootElement(one)
	m.AddElement(oneChild, one)
	m.AddRootElement(two)

	// Link the children together
	m.AddAssociation(oneChild, two, nil)

	// Create the map
	elMap := map[string]mdl.Element{
		"User":     actor,
		"One":      one,
		"OneChild": oneChild,
		"Two":      two,
	}

	return &m, elMap
}

func assertOutput(t *testing.T, output string, formatString string, elMap map[string]mdl.Element) {
	// Get the elements
	els := make([]mdl.Element, 0, len(elMap))
	for _, v := range elMap {
		els = append(els, v)
	}
	// Sort them by name
	sort.Slice(els, func(i, j int) bool {
		return els[i].Name() < els[j].Name()
	})
	// Get the IDs
	var IDsInterface []interface{} = make([]interface{}, len(els))
	for i, e := range els {
		IDsInterface[i] = e.ID()
	}
	assert.Equal(t, output, fmt.Sprintf(formatString, IDsInterface...))
}
