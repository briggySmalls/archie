package api

import (
	mdl "github.com/briggysmalls/archie/core/model"
	"encoding/json"
)

// Parse an API model from a yaml string
func ParseJson(data string) (*mdl.Model, error) {
	// Parse the yaml using the package
	var apiModel Model
	err := json.Unmarshal([]byte(data), &apiModel)
	if err != nil {
		return nil, err
	}
	// Convert to an internal model
	return toInternalModel(apiModel)
}

// Convert an API model to yaml
func ToJson(model *mdl.Model) ([]byte, error) {
	// Convert to serializable model
	apiModel := toApiModel(model)
	// Now marshall this into yaml
	data, err := json.Marshal(apiModel)
	return data, err
}

func (e Element) MarshalJSON() (interface{}, error) {
	// Check if we are an actor
	// Check if all we need to write is the name
	if e.Kind == "" && e.Technology == "" && len(e.Children) == 0 {
		return e.Name, nil
	}
	return ElementWithChildren(e), nil
}
