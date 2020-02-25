package io

import (
	"fmt"

	mdl "github.com/briggysmalls/archie/internal/model"
	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v2"
)

// ParseYaml parses an API model from a yaml string
func ParseYaml(data string) (*mdl.Model, error) {
	// Parse the yaml using the package
	var sModel Model
	err := yaml.UnmarshalStrict([]byte(data), &sModel)
	if err != nil {
		return nil, err
	}
	// Convert to an internal model
	return toInternalModel(sModel)
}

func toInternalModel(apiModel Model) (*mdl.Model, error) {
	// Copy the parsed elements into the new model
	m := mdl.NewModel()
	for _, rootEl := range apiModel.Elements {
		err := updateModelAndRecurse(&m, nil, rootEl)
		if err != nil {
			return nil, err
		}
	}

	// Add top level associations
	err := addAssociations(&m, nil, apiModel.Associations)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

// Add associations for element of a model
func addAssociations(model *mdl.Model, parent mdl.Element, assocs []Association) error {
	// Copy the parsed relationships into the new model
	for _, ass := range assocs {
		// Get the elements
		var src, dest mdl.Element
		var err error
		src, err = model.LookupName(ass.Source, parent)
		if err != nil {
			return err
		}
		dest, err = model.LookupName(ass.Destination, parent)
		if err != nil {
			return err
		}
		// Add a new relationship
		model.AddAssociation(src, dest, ass.Tag)
	}
	return nil
}

func addChildren(model *mdl.Model, parent mdl.Element, children []interface{}) error {
	// Iterate through children
	for _, child := range children {
		// Check whether it is a brief entry or not
		switch i := child.(type) {
		case string:
			// This is a shorthand
			new := mdl.NewItem(i, nil)
			// Add to the model
			model.AddElement(new, parent)
		case map[string]interface{}:
			// Map into an element structure
			var el Element
			var err error
			err = mapstructure.Decode(i, &el)
			if err != nil {
				return err
			}
			// Map into an element structure
			err = updateModelAndRecurse(model, parent, el)
			if err != nil {
				return err
			}
		case map[interface{}]interface{}:
			// See: https://github.com/go-yaml/yaml/issues/139
			// Map into an element structure
			var el Element
			var err error
			err = mapstructure.Decode(i, &el)
			if err != nil {
				return err
			}
			err = updateModelAndRecurse(model, parent, el)
			if err != nil {
				return err
			}
		default:
			return fmt.Errorf("Unexpected type %T", i)
		}
	}
	return nil
}

func updateModelAndRecurse(model *mdl.Model, parent mdl.Element, el Element) error {

	// Handle creating actors
	if el.Kind == "actor" {
		if parent != nil {
			return fmt.Errorf("Only root elements can be actors, element %s is not.", el.Name)
		}
		new := mdl.NewActor(el.Name)
		model.AddElement(new, parent)
		return nil
	}

	// This is a fully-specified element
	new := mdl.NewItem(el.Name, el.Tags)

	// Add to the model
	model.AddElement(new, parent)
	// Add children
	err := addChildren(model, new, el.Children)
	if err != nil {
		return err
	}

	if len(el.Associations) != 0 {
		err := addAssociations(model, new, el.Associations)
		if err != nil {
			return err
		}
	}

	return nil
}
