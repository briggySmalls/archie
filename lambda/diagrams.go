package main

import (
	"github.com/briggysmalls/archie"
	"github.com/briggysmalls/archie/writers"
	"gopkg.in/yaml.v2"
)

type config struct {
	Footer string `yaml:""`
}

type payload struct {
	Model  interface{} `yaml:""`
	Config config      `yaml:""`
}

func landscapeDiagram(yaml string) (string, error) {
	// Get the model
	archie, err := readModel(yaml)
	if err != nil {
		return "", err
	}
	// Create a landscape view
	chart, err := archie.LandscapeView()
	if err != nil {
		return "", err
	}
	// Return diagram in browser
	return chart, nil
}

func contextDiagram(yaml, scope string) (string, error) {
	// Get the model
	archie, err := readModel(yaml)
	if err != nil {
		return "", err
	}
	// Create the view
	chart, err := archie.ContextView(scope)
	if err != nil {
		return "", err
	}
	// Return diagram in browser
	return chart, nil
}

func tagDiagram(yaml, scope, tag string) (string, error) {
	// Get the model
	archie, err := readModel(yaml)
	if err != nil {
		return "", err
	}
	// Create the view
	chart, err := archie.TagView(scope, tag)
	if err != nil {
		return "", err
	}
	// Return diagram in browser
	return chart, nil
}

func readModel(yml string) (archie.Archie, error) {
	// Separate config and model
	p := payload{}
	err := yaml.Unmarshal([]byte(yml), &p)
	if err != nil {
		return nil, err
	}
	// Create an Archie from the model & config
	model, err := yaml.Marshal(p.Model)
	if err != nil {
		return nil, err
	}
	// Create a writer, using the provided config
	writer := writers.PlantUmlStrategy{CustomFooter: p.Config.Footer}
	// Create an archie instance with the writer and model
	archie, err := archie.New(writer, string(model))
	return archie, err
}
