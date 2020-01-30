package main

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type DiagramEvent struct {
	Type           string `json:""`
	Scope          string `json:""`
	Tag            string `json:""`
	ModelAndConfig string `json:""`
}

func handler(ctx context.Context, event DiagramEvent) (string, error) {
	// Obtain yaml from the base64 encoded field
	modelAndConfigBytes, err := base64.StdEncoding.DecodeString(event.ModelAndConfig)
	if err != nil {
		return "", fmt.Errorf("Couldn't decode model: %w", err)
	}
	modelAndConfig := string(modelAndConfigBytes)
	// Call relevant diagramming function
	switch event.Type {
	case "landscape":
		return landscapeDiagram(modelAndConfig)
	case "context":
		return contextDiagram(modelAndConfig, event.Scope)
	case "tag":
		return tagDiagram(modelAndConfig, event.Scope, event.Tag)
	default:
		return "", fmt.Errorf("Invalid diagram type: %s", event.Type)
	}
}

func main() {
	lambda.Start(handler)
}
