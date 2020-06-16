package views

import (
	mdl "github.com/briggysmalls/archie/internal/model"
)

// NewStructureView creates a view that shows the children of the scoped element
func NewStructureView(model *mdl.Model, scope mdl.Element, tag string) mdl.Model {
	var primaries []mdl.Element
	if scope == nil {
		primaries = model.RootElements()
	} else {
		primaries = []mdl.Element{scope}
	}
	view := mdl.NewModel()
	for _, primary := range primaries {
		if primary.IsActor() || hasOnlyOtherTags(tag, primary.Tags()) {
			continue
		}

		view.AddRootElement(primary)
		addChildren(model, &view, primary, tag)
	}

	return view
}

func addChildren(from, to *mdl.Model, of mdl.Element, tag string) {
	for _, kid := range from.Children(of) {
		if hasOnlyOtherTags(tag, kid.Tags()) {
			continue
		}

		to.AddRootElement(kid)
		to.AddAssociation(of, kid, []string{})
		addChildren(from, to, kid, tag)
	}
}

func hasOnlyOtherTags(tag string, tags []string) bool {
	if len(tags) == 0 || len(tag) == 0 {
		return false
	}

	for _, t := range tags {
		if t == tag {
			return false
		}
	}

	return true
}
