package views

import (
	mdl "github.com/briggysmalls/archie/internal/model"
)

// NewStructureView creates a view that shows the children of the scoped element
func NewStructureView(model *mdl.Model, scope mdl.Element, tag string, maxDepth int) mdl.Model {
	var primaries []mdl.Element
	if scope == nil {
		primaries = model.RootElements()
	} else {
		primaries = []mdl.Element{scope}
	}
	view := mdl.NewModel()
	for _, primary := range primaries {
		if primary.IsActor() ||
			(len(primary.Tags()) > 0 && len(tag) > 0 && !primary.HasTag(tag)) {
			continue
		}

		hasChildren := addChildren(model, &view, primary, tag, 1, maxDepth)
		if hasChildren || len(tag) == 0 || primary.HasTag(tag) {
			view.AddRootElement(primary)
		}
	}

	return view
}

func addChildren(from, to *mdl.Model, of mdl.Element, tag string, depth int, maxDepth int) bool {
	if maxDepth != 0 && depth > maxDepth {
		return false
	}
	atLeastOneChild := false

	for _, kid := range from.Children(of) {
		hasChildren := addChildren(from, to, kid, tag, depth+1, maxDepth)
		if hasChildren || len(tag) == 0 || kid.HasTag(tag) {
			atLeastOneChild = true
			to.AddRootElement(kid)
			to.AddAssociation(of, kid, []string{})
		}
	}

	return atLeastOneChild
}
