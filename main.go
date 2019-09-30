package main

import "C"
import "github.com/briggysmalls/archie/core/views"
import "github.com/briggysmalls/archie/io/yaml"
import "github.com/briggysmalls/archie/io/writers"

//export ParseYaml
func ParseYaml(data string) (*mdl.Model, error) {
	return yaml.ParseYaml(data)
}

//export NewLandscapeView
func NewLandscapeView(model *mdl.Model) mdl.Model {
	return views.NewLandscapeView(model)
}

//export NewContextView
func NewContextView(model *mdl.Model, scope *mdl.Element) mdl.Model {
	return views.NewContextView(model, scope)
}

func main() {}
