package writers

import (
	"fmt"
	mdl "github.com/briggysmalls/archie/internal/model"
	"strings"
)

const (
	SPACES_IN_TAB = 4
)

type Element interface {
	mdl.Element
}

type Relationship interface {
	mdl.Relationship
}

type Writer interface {
	Write(mdl.Model) (string, error)
}

type writer struct {
	strategy Strategy
	writer   Writer
	indent   uint
	builder  strings.Builder
	model    *mdl.Model
}

type Strategy interface {
	Header(writer Scribe)
	Footer(writer Scribe)
	Element(writer Scribe, element Element)
	StartParentElement(writer Scribe, element Element)
	EndParentElement(writer Scribe, element Element)
	Association(writer Scribe, association Relationship)
}

type Scribe interface {
	FullName(mdl.Element) (string, error)
	WriteLine(string, ...interface{})
	WriteString(bool, string, ...interface{})
	UpdateIndent(int)
}

func New(strategy Strategy) writer {
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
		d.strategy.Association(d, Relationship(rel))
	}
	// Write footer
	d.strategy.Footer(d)
	// Return result
	return d.builder.String(), nil
}

func (d *writer) FullName(element mdl.Element) (string, error) {
	name, err := d.model.Name(element)
	return name, err
}

// Recursive function for drawing elements
func (d *writer) writeElement(model *mdl.Model, el mdl.Element) error {
	var err error
	children := model.Children(el)
	if len(children) == 0 {
		// Write a simple component
		d.strategy.Element(d, Element(el))
		return nil
	}
	// Start a new package
	d.strategy.StartParentElement(d, Element(el))
	for _, child := range children {
		// Recurse through children
		err = d.writeElement(model, child)
		if err != nil {
			return err
		}
	}
	d.strategy.EndParentElement(d, Element(el))
	return nil
}

// Write a correctly-indented string terminated with a newline
func (d *writer) WriteLine(format string, args ...interface{}) {
	// Write the string (with indent)
	d.WriteString(true, fmt.Sprintf(format, args...))
	// Append a newline
	d.WriteString(false, "\n")
}

// Append the provided string to the current line
func (d *writer) WriteString(withIndent bool, format string, args ...interface{}) {
	// Add an indent, if requested
	if withIndent {
		d.builder.WriteString(fmt.Sprintf("%*s", d.indent*SPACES_IN_TAB, ""))
	}
	// Write the string
	d.builder.WriteString(fmt.Sprintf(format, args...))
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
