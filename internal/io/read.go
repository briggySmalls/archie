package io

import (
	"fmt"
	mdl "github.com/briggysmalls/archie/internal/model"
	"github.com/mitchellh/mapstructure"
)

func toInternalModel(apiModel Model) (*mdl.Model, error) {
	// Copy the parsed elements into the new model
	m := mdl.NewModel()
	for _, rootEl := range apiModel.Elements {
		// Add these first elements to the root
		var el mdl.Element
		if rootEl.Kind == "actor" {
			el = mdl.NewActor(rootEl.Name)
		} else {
			el = mdl.NewItem(rootEl.Name, rootEl.Technology)
		}
		m.AddRootElement(el)
		// Now recursively add children
		err := addChildren(&m, el, rootEl.Children)
		if err != nil {
			return nil, err
		}
	}
	// Copy the parsed relationships into the new model
	for _, ass := range apiModel.Associations {
		// Get the elements
		var src, dest mdl.Element
		var err error
		src, err = m.LookupName(ass.Source)
		if err != nil {
			return nil, err
		}
		dest, err = m.LookupName(ass.Destination)
		if err != nil {
			return nil, err
		}
		// Add a new relationship
		m.AddAssociation(src, dest)
	}
	return &m, nil
}

func addChildren(model *mdl.Model, parent mdl.Element, children []interface{}) error {
	// Iterate through children
	for _, child := range children {
		// Check whether it is a brief entry or not
		switch i := child.(type) {
		case string:
			// This is a shorthand
			new := mdl.NewItem(i, "")
			// Add to the model
			model.AddElement(new, parent)
		case map[string]interface{}:
			// Map into an element structure
			var el Element
			err := mapstructure.Decode(i, &el)
			if err != nil {
				return err
			}
			// Map into an element structure
			updateModelAndRecurse(model, parent, el)
		case map[interface{}]interface{}:
			// See: https://github.com/go-yaml/yaml/issues/139
			// Map into an element structure
			var el Element
			err := mapstructure.Decode(i, &el)
			if err != nil {
				return err
			}
			updateModelAndRecurse(model, parent, el)
		default:
			return fmt.Errorf("Unexpected type %T", i)
		}
	}
	return nil
}

func updateModelAndRecurse(model *mdl.Model, parent mdl.Element, el Element) error {
	// This is a fully-specified element
	new := mdl.NewItem(el.Name, el.Technology)
	// Add to the model
	model.AddElement(new, parent)
	// Add children
	err := addChildren(model, new, el.Children)
	if err != nil {
		return err
	}
	return nil
}
