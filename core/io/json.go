package io

import (
	"encoding/json"
	mdl "github.com/briggysmalls/archie/core/model"
)

// Parse an API model from a yaml string
func ParseJson(data string) (*mdl.Model, error) {
	// Parse the yaml using the package
	var sModel Model
	err := json.Unmarshal([]byte(data), &sModel)
	if err != nil {
		return nil, err
	}
	// Convert to an internal model
	return toInternalModel(sModel)
}
