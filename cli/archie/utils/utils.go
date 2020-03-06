package utils

import (
	"github.com/briggysmalls/archie"
	"github.com/briggysmalls/archie/writers"
	"gopkg.in/yaml.v2"
)

type config struct {
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
