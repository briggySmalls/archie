package io

import (
	mdl "github.com/briggysmalls/archie/core/model"
	"github.com/ghodss/yaml"
)

// Parse an API model from a yaml string
func ParseYaml(data string) (*mdl.Model, error) {
	// Convert yaml to json
	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		return nil, err
	}
	// Run usual json parsing
	return ParseJson(string(json))
}

// Convert an API model to yaml
func ToYaml(model *mdl.Model) (string, error) {
	// Convert to serializable model
	sModel := toSerialisable(model)
	// Now marshal this into yaml
	data, err := yaml.Marshal(sModel)
	return string(data), err
}
