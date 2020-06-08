package views

import (
	mdl "github.com/briggysmalls/archie/internal/model"
)

// NewStructureView creates a view that shows the children of the scoped element
func NewStructureView(model *mdl.Model, scope mdl.Element) mdl.Model {
	var primaries []mdl.Element
	if scope == nil {
		primaries = model.RootElements()
	} else {
		primaries = []mdl.Element{scope}
	}
	view := mdl.NewModel()
	for _, primary := range primaries {
		if primary.IsActor() {
			continue
		}

		view.AddRootElement(primary)
		addChildren(model, &view, primary)
	}

	return view
}

func addChildren(from, to *mdl.Model, of mdl.Element) {
	for _, kid := range from.Children(of) {
		to.AddRootElement(kid)
		to.AddAssociation(of, kid, []string{})
		addChildren(from, to, kid)
	}
}
