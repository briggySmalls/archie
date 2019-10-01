package api

import (
	mdl "github.com/briggysmalls/archie/core/model"
	"gopkg.in/yaml.v3"
)

// Parse an API model from a yaml string
func ParseYaml(data string) (*mdl.Model, error) {
	// Parse the yaml using the package
	var apiModel Model
	err := yaml.Unmarshal([]byte(data), &apiModel)
	if err != nil {
		return nil, err
	}
	// Convert to an internal model
	return toInternalModel(apiModel)
}

// Convert an API model to yaml
func ToYaml(model *mdl.Model) (string, error) {
	// Convert to serializable model
	apiModel := toApiModel(model)
	// Now marshal this into yaml
	data, err := yaml.Marshal(apiModel)
	return string(data), err
}

func (e Element) MarshalYAML() (interface{}, error) {
	// Check if we are an actor
	// Check if all we need to write is the name
	if e.Kind == "" && e.Technology == "" && len(e.Children) == 0 {
		return e.Name, nil
	}
	return ElementWithChildren(e), nil
}
