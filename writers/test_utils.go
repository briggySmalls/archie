package writers

import (
	"fmt"
	mdl "github.com/briggysmalls/archie/internal/model"
	"gotest.tools/assert"
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
	m.AddAssociation(oneChild, two, "")

	// Create the map
	elMap := map[string]mdl.Element{
		"User":     actor,
		"One":      one,
		"OneChild": oneChild,
		"Two":      two,
	}

	return &m, elMap
}

func assertOutput(t *testing.T, output string, formatString string, IDs []string) {
	var IDsInterface []interface{} = make([]interface{}, len(IDs))
	for i, d := range IDs {
		IDsInterface[i] = d
	}
	assert.Equal(t, output, fmt.Sprintf(formatString, IDsInterface...))
}
