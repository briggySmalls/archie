package drawers

import (
	"fmt"
	mdl "github.com/briggysmalls/archie/core/model"
	"strings"
)

const (
	SPACES_IN_TAB = 4
)

type Drawer interface {
	Draw(mdl.Model) (string, error)
}

type DrawConfig interface {
	Header(writer Writer)
	Footer(writer Writer)
	Element(writer Writer, element *mdl.Element)
	StartParentElement(writer Writer, element *mdl.Element)
	EndParentElement(writer Writer, element *mdl.Element)
	Association(writer Writer, association mdl.Relationship)
}

type Writer interface {
	FullName(*mdl.Element) (string, error)
	Write(string, ...interface{})
	UpdateIndent(int)
}

type drawer struct {
	config  DrawConfig
	writer  Writer
	indent  uint
	builder strings.Builder
	model   *mdl.Model
}

func newDrawer(config DrawConfig) drawer {
	return drawer{config: config}
}

// Entrypoint for the drawer
func (d *drawer) Draw(model mdl.Model) (string, error) {
	// Reset the drawer
	d.indent = 0
	d.builder.Reset()
	d.model = &model
	// Add the header
	d.config.Header(d)
	// Draw the elements, recursively
	var err error
	for _, el := range model.Children(nil) {
		err = d.drawElement(&model, el)
		if err != nil {
			return "", err
		}
	}
	// Now draw the relationships
	for _, rel := range model.Associations {
		d.config.Association(d, rel)
	}
	// Write footer
	d.config.Footer(d)
	// Return result
	return d.builder.String(), nil
}

func (d *drawer) FullName(element *mdl.Element) (string, error) {
	name, err := d.model.Name(element)
	return name, err
}

// Recursive function for drawing elements
func (d *drawer) drawElement(model *mdl.Model, el *mdl.Element) error {
	var err error
	children := model.Children(el)
	if len(children) == 0 {
		// Write a simple component
		d.config.Element(d, el)
		return nil
	}
	// Start a new package
	d.config.StartParentElement(d, el)
	for _, child := range children {
		// Recurse through children
		err = d.drawElement(model, child)
		if err != nil {
			return err
		}
	}
	d.config.EndParentElement(d, el)
	return nil
}

func (d *drawer) Write(format string, args ...interface{}) {
	// Write the string
	d.builder.WriteString(fmt.Sprintf("%*s%s\n", d.indent*SPACES_IN_TAB, "", fmt.Sprintf(format, args...)))
}

func (d *drawer) UpdateIndent(indicator int) {
	switch {
	case indicator > 0:
		d.indent++
	case indicator < 0:
		d.indent--
	default:
		// Do nothing
	}
}
