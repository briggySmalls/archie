package writers

import (
	"fmt"
	mdl "github.com/briggysmalls/archie/internal/model"
	"sort"
	"strings"
)

const (
	spacesInTab = 4
)

// Element wraps an internal model element for use in defining writers
type Element interface {
	mdl.Element
}

// Association wraps an internal model association for use in defining writers
type Association interface {
	mdl.Association
}

// Writer is an interface for writing a diagram from a model
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

// Strategy is an interface for writing diagram strings from model entities
type Strategy interface {
	Header(writer Scribe)
	Footer(writer Scribe)
	Element(writer Scribe, element Element)
	StartParentElement(writer Scribe, element Element)
	EndParentElement(writer Scribe, element Element)
	Association(writer Scribe, association Association)
}

// Scribe is an interface for
type Scribe interface {
	FullName(mdl.Element) (string, error)
	WriteLine(string, ...interface{})
	WriteString(bool, string, ...interface{})
	UpdateIndent(int)
}

// New creates a writer from a strategy
func New(strategy Strategy) Writer {
	return &writer{strategy: strategy}
}

// Write produces a diagram string from a model
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
		d.strategy.Association(d, Association(rel))
	}
	// Write footer
	d.strategy.Footer(d)
	// Return result
	return d.builder.String(), nil
}

// FullName gets the full name of an element
func (d *writer) FullName(element mdl.Element) (string, error) {
	name, err := d.model.Name(element)
	return name, err
}

// Recursive function for drawing elements
func (d *writer) writeElement(model *mdl.Model, el mdl.Element) error {
	var err error
	// Try to collapse parents if possible
	collapsed, actual := collapseSingleParents(model, el)
	// collapsed := el
	// actual := el
	// Determine how to draw our collapsed element
	children := model.Children(actual)
	if len(children) == 0 {
		// Write a simple component
		d.strategy.Element(d, Element(collapsed))
		return nil
	}
	// Start a parent
	d.strategy.StartParentElement(d, Element(collapsed))
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

// WriteString appends the provided string to the current line
func (d *writer) WriteString(withIndent bool, format string, args ...interface{}) {
	// Add an indent, if requested
	if withIndent {
		d.builder.WriteString(fmt.Sprintf("%*s", d.indent*spacesInTab, ""))
	}
	// Write the string
	d.builder.WriteString(fmt.Sprintf(format, args...))
}

// UpdateIndent modifies the indentation of the writer.
// Positive values increase the indent.
// Negative values decrease the indent.
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

func collapseSingleParents(model *mdl.Model, el mdl.Element) (collapsed, actual mdl.Element) {
	// Get the collapsable elements
	collapsable := getCollapsable(model, el, []mdl.Element{el})
	// Short-circuit if there is no collapsing to be done
	if len(collapsable) == 1 {
		collapsed = collapsable[0]
		actual = collapsable[0]
		return
	}
	// Pull out the names
	names := make([]string, len(collapsable))
	for i, el := range collapsable {
		names[i] = el.Name()
	}
	// Create a new collapsed element
	collapsed = mdl.NewItem(strings.Join(names, "/"), collapsable[0].Tags())
	actual = collapsable[len(collapsable)-1]
	return
}

func getCollapsable(model *mdl.Model, el mdl.Element, els []mdl.Element) []mdl.Element {
	children := model.Children(el)
	// If the next element is not a single-parent, bail
	if len(children) != 1 {
		return els
	}
	// If tags differ, bail
	child := children[0]
	if !areTagsEqual(el.Tags(), child.Tags()) {
		return els
	}
	// Recurse (there may be more!)
	return getCollapsable(model, child, append(els, child))
}

func areTagsEqual(a, b []string) bool {
	// Short-circuit if lengths differ
	if len(a) != len(b) {
		return false
	}
	sort.Strings(a)
	sort.Strings(b)
	// Ensure each entry matches
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
