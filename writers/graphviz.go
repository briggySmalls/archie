package writers

import (
	"fmt"
	"strings"
)

// GraphvizStrategy is the strategy for drawing a graphviz graph from a model/view.
type GraphvizStrategy struct {
	CustomFooter string
}

var colors = []string{
	"#8dd3c7",
	"#ffffb3",
	"#bebada",
	"#fb8072",
	"#80b1d3",
	"#fdb462",
	"#b3de69",
	"#fccde5",
	"#d9d9d9",
	"#bc80bd",
	"#ccebc5",
	"#ffed6f",
}

// Preallocate a map from tag to colour
var colourMap = make(map[string]string)

// Header writes the header
func (p GraphvizStrategy) Header(scribe Scribe) {
	scribe.WriteLine("digraph arch {")
	scribe.UpdateIndent(1)
	scribe.WriteLine("graph [fontname=Helvetica]")
	scribe.WriteLine(`edge [fontsize=9; fontname=Helvetica; color="#333333"]`)
	scribe.WriteLine("node [shape=plaintext; margin=0; fontname=Helvetica]")
}

// Footer writes a footer
func (p GraphvizStrategy) Footer(scribe Scribe) {
	for _, line := range strings.Split(p.CustomFooter, "\n") {
		scribe.WriteLine(line)
	}
	scribe.UpdateIndent(-1)
	scribe.WriteLine("}")
}

// Element writes an element
func (p GraphvizStrategy) Element(scribe Scribe, element Element) {
	scribe.WriteLine(`"%p" [`, element)
	scribe.UpdateIndent(1)

	if element.IsActor() {
		scribe.WriteLine(`color = "#333333"`)
		scribe.WriteLine("shape = circle")
		scribe.WriteLine("margin = 0.04")
		scribe.WriteLine("label = <%s>", element.Name())
	} else {
		scribe.WriteLine(`label = <`)
		scribe.UpdateIndent(1)
		// We render items as a table
		scribe.WriteLine(`<TABLE BORDER="0" CELLBORDER="0" CELLSPACING="0">`)
		// Create a header row for tags, if present
		if len(element.Tags()) > 0 {
			scribe.WriteLine(makeTags(element.Tags()))
		}
		// Start a row for the item name
		scribe.WriteString(true, `<TR><TD`)
		// If there are multiple tags then we want this column to span all of them
		if len(element.Tags()) > 0 {
			scribe.WriteString(false, ` COLSPAN="%d"`, len(element.Tags()))
		}
		scribe.WriteString(false, " CELLPADDING=\"10\" BGCOLOR=\"#dbdbdb\">%s</TD></TR>\n", element.Name())
		scribe.WriteLine("</TABLE>>")
		scribe.UpdateIndent(-1)
	}
	scribe.UpdateIndent(-1)
	scribe.WriteLine(`];`)
}

// StartParentElement writes the start of an enclosing/parent element
func (p GraphvizStrategy) StartParentElement(scribe Scribe, element Element) {
	scribe.WriteLine(`subgraph "cluster_%p" {`, element)
	scribe.UpdateIndent(1)
	scribe.WriteLine(`label = <%s>`, element.Name())
}

// EndParentElement writes the end of an enclosing/parent element
func (p GraphvizStrategy) EndParentElement(scribe Scribe, element Element) {
	scribe.UpdateIndent(-1)
	scribe.WriteLine("}")
}

// Association writes an association
func (p GraphvizStrategy) Association(scribe Scribe, association Association) {
	scribe.WriteString(true, `"%s" -> "%s"`, association.Source().ID(), association.Destination().ID())
	if len(association.Tags()) > 0 {
		scribe.WriteString(false, ` [label = "%s"]`, strings.Join(association.Tags(), ",\\n"))
	}
	scribe.WriteString(false, "\n")
}

// makeTags creates a column for every tag
func makeTags(tags []string) string {
	// Short-circuit if there are no tags
	if len(tags) == 0 {
		return ""
	}
	// Start building a row
	var sb strings.Builder
	sb.WriteString("<TR>")
	// Create a column per tag
	for _, tag := range tags {
		// First check if we've encountered the tag already
		color, hasColor := colourMap[tag]
		// Allocate a new colour for the tag
		if !hasColor {
			selectedColor := colors[len(colourMap)%len(colors)]
			colourMap[tag] = selectedColor
			color = selectedColor
		}
		// Write the column
		sb.WriteString(fmt.Sprintf(`<TD CELLPADDING="5" BGCOLOR="%s"><I><FONT POINT-SIZE="9">%s</FONT></I></TD>`, color, tag))
	}
	sb.WriteString("</TR>")
	return sb.String()
}
