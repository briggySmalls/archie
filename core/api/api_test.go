package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var yaml = `
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

func TestLandscape(t *testing.T) {
	// Create an archie
	a, err := NewArchieFromYaml(yaml)
	assert.NotNil(t, err)

	// Create a landscape view
	view, err := a.LandscapeView()
	assert.NotNil(t, err)

	// Assert
	expected := `
{
	"elements": [
		{"name": "user", "kind": "actor"},
		{"name": "sound system"}
	],
	"associations": [
		{"source": "user", "destination": "sound system"},
		{"source": "sound system", "destination": "user"}
	]
}
`
	assert.JSONEq(t, expected, view)
}

func TestContext(t *testing.T) {
	// Create an archie
	a, err := NewArchieFromYaml(yaml)
	assert.NotNil(t, err)

	// Create a landscape view
	view, err := a.ContextView("sound system")
	assert.NotNil(t, err)

	// Assert
	expected := `
{
	"elements": [
		{"name": "user", "kind": "actor"},
		{
			"name": "sound system",
			"children": [
				{"name": "speaker"},
				{"name": "amplifier"}
			]
		}
	],
	"associations": [
		{"source": "sound system/speaker", "destination": "user"},
		{"source": "sound system/amplifier", "destination": "sound system/speaker"},
		{"source": "user", "destination": "sound system/amplifier"},
	]
}
`
	assert.JSONEq(t, expected, view)
}
