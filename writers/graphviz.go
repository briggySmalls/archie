package writers

import (
	"fmt"
	"strings"
)

// GraphvizStrategy is the strategy for drawing a graphviz graph from a model/view.
type GraphvizStrategy struct {
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

var colorMap = make(map[string]string)

// Header writes the header
func (p GraphvizStrategy) Header(scribe Scribe) {
	scribe.WriteLine("graph arch {")
	scribe.UpdateIndent(1)
	scribe.WriteLine("graph [fontname=Helvetica]")
	scribe.WriteLine(`edge [fontsize=9; fontname=Helvetica; color="#333333"]`)
	scribe.WriteLine("node [shape=plaintext; margin=0; fontname=Helvetica]")
}

// Footer writes a footer
func (p GraphvizStrategy) Footer(scribe Scribe) {
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
		scribe.WriteLine(`<TABLE BORDER="0" CELLBORDER="0" CELLSPACING="0">`)
		if len(element.Tags()) > 0 {
			scribe.WriteLine(makeTags(element.Tags()))
		}
		scribe.WriteLine(`<TR><TD COLSPAN="%d" CELLPADDING="10" BGCOLOR="#dbdbdb">%s</TD></TR>`, len(element.Tags()), element.Name())
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
	scribe.WriteString(true, `"%s" -- "%s"`, association.Source().ID(), association.Destination().ID())
	if len(association.Tags()) > 0 {
		scribe.WriteString(false, ` [label = "%s"]`, strings.Join(association.Tags(), ", "))
	}
	scribe.WriteString(false, "\n")
}

// makeTags creates the title title containing tags
func makeTags(tags []string) string {
	if len(tags) == 0 {
		return ""
	}

	var sb strings.Builder
	sb.WriteString("<TR>")

	for _, tag := range tags {
		color, hasColor := colorMap[tag]
		if !hasColor {
			selectedColor := colors[len(colorMap)]
			colorMap[tag] = selectedColor
			color = selectedColor
		}

		templ := `<TD CELLPADDING="5" BGCOLOR="%s"><I><FONT POINT-SIZE="9">%s</FONT></I></TD>`
		sb.WriteString(fmt.Sprintf(templ, color, tag))
	}

	sb.WriteString("</TR>")

	return sb.String()
}
