package drawers

import (
	"github.com/briggysmalls/archie/internal/views"
	"strings"
)

type PlantUmlDrawer struct {
}

func (p *PlantUmlDrawer) Draw(view views.View) string {
	// Create a builder
	var b strings.Builder
	// Add the header
	b.WriteString("@startuml\n")
	// Add nested blocks
	for _, el in range view.Elements() {

	}
}
