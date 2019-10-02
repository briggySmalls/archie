package writers

import (
	"fmt"
	mdl "github.com/briggysmalls/archie/core/model"
	"strings"
)

const (
	SPACES_IN_TAB = 4
)

type Writer interface {
	Write(mdl.Model) (string, error)
}

type WriterStrategy interface {
	Header(writer Scribe)
	Footer(writer Scribe)
	Element(writer Scribe, element *mdl.Element)
	StartParentElement(writer Scribe, element *mdl.Element)
	EndParentElement(writer Scribe, element *mdl.Element)
	Association(writer Scribe, association mdl.Relationship)
}

type Scribe interface {
	FullName(*mdl.Element) (string, error)
	WriteLine(string, ...interface{})
	UpdateIndent(int)
}

type writer struct {
	strategy WriterStrategy
	writer   Writer
	indent   uint
	builder  strings.Builder
	model    *mdl.Model
}

func New(strategy WriterStrategy) writer {
	return writer{strategy: strategy}
}

// Entrypoint for the writer
func (d *writer) Write(model mdl.Model) (string, error) {
	// Reset the writer
	d.indent = 0
	d.builder.Reset()
	d.model = &model
	// Add the header
	d.strategy.Header(d)
	// Draw the elements, recursively
	var err error
	for _, el := range model.Children(nil) {
		err = d.writeElement(&model, el)
		if err != nil {
			return "", err
		}
	}
	// Now draw the relationships
	for _, rel := range model.Associations {
		d.strategy.Association(d, rel)
	}
	// Write footer
	d.strategy.Footer(d)
	// Return result
	return d.builder.String(), nil
}

func (d *writer) FullName(element *mdl.Element) (string, error) {
	name, err := d.model.Name(element)
	return name, err
}

// Recursive function for drawing elements
func (d *writer) writeElement(model *mdl.Model, el *mdl.Element) error {
	var err error
	children := model.Children(el)
	if len(children) == 0 {
		// Write a simple component
		d.strategy.Element(d, el)
		return nil
	}
	// Start a new package
	d.strategy.StartParentElement(d, el)
	for _, child := range children {
		// Recurse through children
		err = d.writeElement(model, child)
		if err != nil {
			return err
		}
	}
	d.strategy.EndParentElement(d, el)
	return nil
}

func (d *writer) WriteLine(format string, args ...interface{}) {
	// Write the string
	d.builder.WriteString(fmt.Sprintf("%*s%s\n", d.indent*SPACES_IN_TAB, "", fmt.Sprintf(format, args...)))
}

func (d *writer) UpdateIndent(indicator int) {
	switch {
	case indicator > 0:
		d.indent++
	case indicator < 0:
		d.indent--
	default:
		// Do nothing
	}
}
