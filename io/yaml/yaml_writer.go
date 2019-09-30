package yaml

import (
	mdl "github.com/briggysmalls/archie/core/model"
	"gopkg.in/yaml.v3"
)

func ToYaml(model *mdl.Model) (string, error) {
	m := Model{}

	// Copy elements into yaml model
	for _, child := range model.RootElements() {
		// Create a yaml element from the root element
		el := copyElementToYamlModel(model, child)
		// Now add this element to our yaml model
		m.Elements = append(m.Elements, el)
	}

	// Copy associations into yaml model
	for _, rel := range model.Associations {
		// Create a relationship
		yamlRel := Association{
			Source:      name(model, rel.Source),
			Destination: name(model, rel.Destination),
		}
		// Add it to the yaml model
		m.Associations = append(m.Associations, yamlRel)
	}

	// Now marshall this into yaml
	data, err := yaml.Marshal(m)
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

func copyElementToYamlModel(model *mdl.Model, modelElement *mdl.Element) Element {
	// Create an yaml element from the model element
	el := Element{
		Name:       modelElement.Name,
		Technology: modelElement.Technology,
	}
	// Update with type if relevant
	if modelElement.IsActor() {
		el.Kind = "actor"
	}
	// Now copy children
	for _, child := range model.Children(modelElement) {
		// Create a new yaml child from the model child
		yamlChild := copyElementToYamlModel(model, child)
		// Append the child
		el.Children = append(el.Children, yamlChild)
	}
	return el
}

func name(model *mdl.Model, element *mdl.Element) string {
	name, err := model.Name(element)
	if err != nil {
		panic(err)
	}
	return name
}
