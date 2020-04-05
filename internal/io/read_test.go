package io

import (
	"testing"

	mdl "github.com/briggysmalls/archie/internal/model"

	"fmt"

	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

var data = `
elements:
  - name: user
    kind: actor
  - name: sound system
    children:
      - name: speaker
        children:
          - name: enclosure
            tags: [physical]
          - name: driver
            tags: [electronics, mechanical]
          - connector
          - cable
      - name: amplifier
        children:
          - audio in connector
          - audio out connector
          - ac-dc converter
          - mixer
          - amplifier circuit
          - name: power button
            tags: [electronics, mechanical]
          - name: input select
            tags: [electronics, mechanical]
        associations:
          - source: ac-dc converter
            destination: amplifier circuit
    associations:
      - source: amplifier/audio out connector
        destination: speaker/cable

associations:
  - source: user
    destination: sound system/amplifier/input select
    tags: [press]
  - source: sound system/amplifier/input select
    destination: sound system/amplifier/mixer
    tags: [signal]
  - source: sound system/amplifier/mixer
    destination: sound system/amplifier/audio in connector
  - source: sound system/amplifier/ac-dc converter
    destination: sound system/amplifier/mixer
    tags: [power]
  - source: sound system/amplifier/amplifier circuit
    destination: sound system/amplifier/audio out connector
  - source: sound system/speaker/cable
    destination: sound system/speaker/connector
  - source: sound system/speaker/connector
    destination: sound system/speaker/driver
  - source: sound system/speaker/driver
    destination: sound system/speaker/enclosure
  - source: sound system/amplifier/power button
    destination: sound system/amplifier/ac-dc converter
`

// Test creating an item
func TestRead(t *testing.T) {
	// Read the model
	m, err := ParseYaml(data)
	// Assert some stuff
	assert.NilError(t, err)
	assert.Assert(t, is.Len(m.Elements, 15))
	assert.Assert(t, is.Len(m.Composition, 15))
	assert.Assert(t, is.Len(m.Associations, 11))
	// Be a bit more in-depth
	assert.Assert(t, is.Len(m.RootElements(), 2))
	assertChildrenCount(t, m, "sound system", 2)
	assertChildrenCount(t, m, "sound system/speaker", 4)
	assertChildrenCount(t, m, "sound system/amplifier", 7)
	// Check some element tags
	assertElementTags(t, m, "sound system/speaker/driver", []string{"electronics", "mechanical"})

	// Check some association tags
	ass, err := findFirstAssociation(m, "user", "input select")
	assert.NilError(t, err)
	assert.Assert(t, is.Contains(ass.Tags(), "press"))
	assert.Equal(t, len(ass.Tags()), 1)

	// Check nested association
	ass, err = findFirstAssociation(m, "ac-dc converter", "amplifier circuit")
	assert.NilError(t, err)

	// Check nested association on root element
	ass, err = findFirstAssociation(m, "audio out connector", "cable")
	assert.NilError(t, err)
}

func assertChildrenCount(t *testing.T, m *mdl.Model, name string, length int) {
	// Lookup the name
	el, err := m.LookupName(name, nil)
	assert.NilError(t, err)
	// Now assert the number of children is as expected
	assert.Assert(t, is.Len(m.Children(el), length))
}

func assertElementTags(t *testing.T, m *mdl.Model, name string, expected []string) {
	// Lookup the name
	el, err := m.LookupName(name, nil)
	assert.NilError(t, err)
	// Assert the tag slices match
	tags := el.Tags()
	assert.Equal(t, len(expected), len(tags))
	for _, tag := range tags {
		assert.Assert(t, is.Contains(expected, tag))
	}
}

func findFirstAssociation(model *mdl.Model, source, destination string) (mdl.Association, error) {
	for _, rel := range model.Associations {
		if rel.Source().Name() == source && rel.Destination().Name() == destination {
			// We've found a relationship that links the two
			return rel, nil
		}
	}
	return nil, fmt.Errorf("Association not found")
}
