package utils

import (
	"fmt"

	"github.com/briggysmalls/archie"
	"github.com/briggysmalls/archie/writers"
	"gopkg.in/yaml.v2"
)

type config struct {
	Writer string `yaml:""`
	Footer string `yaml:""`
}

type parsedModelAndConfig struct {
	Model  interface{} `yaml:""`
	Config config      `yaml:""`
}

// ReadModel reads a yaml into a archie structure
func ReadModel(modelAndConfig []byte) (archie.Archie, error) {
	// Separate config and model
	p := parsedModelAndConfig{}
	err := yaml.UnmarshalStrict(modelAndConfig, &p)
	if err != nil {
		return nil, err
	}
	// Get the model as yaml
	model, err := yaml.Marshal(p.Model)
	if err != nil {
		return nil, err
	}
	// Create a writer, using the provided config
	var writer writers.Strategy
	switch p.Config.Writer {
	case "":
		// Assume plantuml if not specified explicitly
		fallthrough
	case "plantuml":
		writer = writers.PlantUmlStrategy{CustomFooter: p.Config.Footer}
	case "graphviz":
		writer = writers.GraphvizStrategy{CustomFooter: p.Config.Footer}
	default:
		return nil, fmt.Errorf("Unexpected writer strategy: %s", p.Config.Writer)
	}
	// Create an archie instance with the writer and model
	archie, err := archie.New(writer, string(model))
	return archie, err
}
