package io

import (
  mdl "github.com/briggysmalls/archie/internal/model"
  "testing"

  "gotest.tools/assert"
  is "gotest.tools/assert/cmp"
  "fmt"
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
  - source: user
    destination: sound system/amplifier/input select
    tag: press
  - source: sound system/amplifier/input select
    destination: sound system/amplifier/mixer
    tag: signal
  - source: sound system/amplifier/mixer
    destination: sound system/amplifier/audio in connector
  - source: sound system/amplifier/ac-dc converter
    destination: sound system/amplifier/mixer
    tag: power
  - source: sound system/amplifier/ac-dc converter
    destination: sound system/amplifier/amplifier circuit
  - source: sound system/amplifier/amplifier circuit
    destination: sound system/amplifier/audio out connector
  - source: sound system/amplifier/audio out connector
    destination: sound system/speaker/cable
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
  ass, err := findFirstRelationship(m, "user", "input select")
  assert.NilError(t, err)
  assert.Equal(t, ass.Tag(), "press")
}

func assertChildrenCount(t *testing.T, m *mdl.Model, name string, length int) {
  // Lookup the name
  el, err := m.LookupName(name)
  assert.NilError(t, err)
  // Now assert the number of children is as expected
  assert.Assert(t, is.Len(m.Children(el), length))
}

func assertElementTags(t *testing.T, m *mdl.Model, name string, expected []string) {
  // Lookup the name
  el, err := m.LookupName(name)
  assert.NilError(t, err)
  // Assert the tag slices match
  tags := el.Tags()
  assert.Equal(t, len(expected), len(tags))
  for _, tag := range tags {
    assert.Assert(t, is.Contains(expected, tag))
  }
}

func findFirstRelationship(model *mdl.Model, source, destination string) (mdl.Relationship, error) {
  for _, rel := range model.Associations {
    if rel.Source().Name() == source && rel.Destination().Name() == destination {
      // We've found a relationship that links the two
      return rel, nil
    }
  }
  return nil, fmt.Errorf("Association not found")
}
