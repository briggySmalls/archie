package readers

import (
	"fmt"
	"github.com/briggysmalls/archie/core/types"
	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v3"
)

type Model struct {
	Elements     []Element
	Associations []Association
}

type Element struct {
	Name       string
	Type       string        `yaml:",omitempty"`
	Technology string        `yaml:",omitempty"`
	Children   []interface{} `yaml:",omitempty"`
}

type Association struct {
	Source      string
	Destination string
}

func ParseYaml(data string) (*types.Model, error) {
	// Parse the yaml using the package
	var yamlModel Model
	err := yaml.Unmarshal([]byte(data), &yamlModel)
	if err != nil {
		return nil, err
	}
	// Copy the parsed elements into the new model
	m := types.NewModel()
	for _, rootEl := range yamlModel.Elements {
		// Add these first elements to the root
		var el types.Element
		if rootEl.Type == "actor" {
			el = types.NewActor(rootEl.Name)
		} else {
			el = types.NewItem(rootEl.Name)
		}
		m.AddRootElement(&el)
		// Now recursively add children
		err = addChildren(&m, &el, rootEl.Children)
		if err != nil {
			return nil, err
		}
	}
	// Copy the parsed relationships into the new model
	for _, ass := range yamlModel.Associations {
		// Get the elements
		var src, dest *types.Element
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
	return &m, err
}

func addChildren(model *types.Model, parent *types.Element, children []interface{}) error {
	// Iterate through children
	for _, child := range children {
		// Check whether it is a brief entry or not
		switch i := child.(type) {
		case string:
			// This is a shorthand
			new := types.NewItem(i)
			// Add to the model
			model.AddElement(&new, parent)
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

func updateModelAndRecurse(model *types.Model, parent *types.Element, el Element) error {
	// This is a fully-specified element
	new := types.NewItem(el.Name)
	// Add to the model
	model.AddElement(&new, parent)
	// Add children
	err := addChildren(model, &new, el.Children)
	if err != nil {
		return err
	}
	return nil
}
