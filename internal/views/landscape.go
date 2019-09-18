package views

import (
	"github.com/briggysmalls/archie/internal/types"
)

// Create a system landscape view
func NewLandscapeView(model *types.Model) types.Model {
	// Create a model from the model's root elements
	return CreateSubmodel(model, model.Elements())
}
