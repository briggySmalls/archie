package api

import (
	mdl "github.com/briggysmalls/archie/core/model"
)

// Convert an internal model into the API form (serializable)
func toApiModel(model *mdl.Model) (Model) {
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

	return m
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
		// We are converting an internal model, which should be consistent
		panic(err)
	}
	return name
}
