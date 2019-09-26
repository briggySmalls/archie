package readers

import (
	"testing"

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
            technology: physical
          - name: driver
            technology: electro-mechanical
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
            technology: electro-mechanical
          - name: input select
            technology: electro-mechanical

associations:
  - source: user
    destination: sound system/amplifier/input select
  - source: sound system/amplifier/input select
    destination: sound system/amplifier/mixer
  - source: sound system/amplifier/mixer
    destination: sound system/amplifier/audio in connector
  - source: sound system/amplifier/ac-dc converter
    destination: sound system/amplifier/mixer
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
}
