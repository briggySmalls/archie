package writers

import (
	"fmt"
	"strings"
)

// GraphvizStrategy is the strategy for drawing a graphviz graph from a model/view.
type GraphvizStrategy struct {
}

var colors = []string{
	"#b7eb8f",
	"#87e8de",
	"#d3adf7",
	"#ffadd2",
	"#ffa39e",
}

var colorMap = make(map[string]string)

// Header writes the header mermaid syntax.
func (p GraphvizStrategy) Header(scribe Scribe) {
	scribe.WriteLine("graph arch {")
	scribe.UpdateIndent(1)
	scribe.WriteLine("graph [fontname=Helvetica]")
	scribe.WriteLine("edge [fontsize=9; fontname=Helvetica]")
	scribe.WriteLine("node [shape=plaintext; margin=0; fontname=Helvetica]")
}

// Footer writes a footer in mermaid syntax.
func (p GraphvizStrategy) Footer(scribe Scribe) {
	scribe.UpdateIndent(-1)
	scribe.WriteLine("}")
}

// Element writes an element in mermaid syntax.
func (p GraphvizStrategy) Element(scribe Scribe, element Element) {
	scribe.WriteLine(`"%p" [`, element)
	scribe.UpdateIndent(1)

	scribe.WriteLine(`label = <<TABLE BORDER="0" CELLBORDER="0" CELLSPACING="0">%s<TR><TD CELLPADDING="10" BGCOLOR="#EEEEEE">%s</TD></TR></TABLE>>`, makeTags(element.Tags()), element.Name())

	scribe.UpdateIndent(-1)
	scribe.WriteLine(`];`)
}

// StartParentElement writes the start of an enclosing/parent element in mermaid syntax.
func (p GraphvizStrategy) StartParentElement(scribe Scribe, element Element) {
	scribe.WriteLine(`subgraph "cluster_%p" {`, element)
	scribe.UpdateIndent(1)
	scribe.WriteLine(`label = <%s>`, element.Name())
}

// EndParentElement writes the end of an enclosing/parent element in mermaid syntax.
func (p GraphvizStrategy) EndParentElement(scribe Scribe, element Element) {
	scribe.WriteLine("}")
	scribe.UpdateIndent(-1)
}

// Association writes an association in mermaid syntax
func (p GraphvizStrategy) Association(scribe Scribe, association Association) {
	scribe.WriteString(true, `"%s" -- "%s"`, association.Source().ID(), association.Destination().ID())
	if association.Tag() != "" {
		scribe.WriteString(false, " [label = %s]", association.Tag())
	}
	scribe.WriteString(false, "\n")
}

func makeTags(tags []string) string {
	if len(tags) == 0 {
		return ""
	}

	tagsStr := strings.Join(tags, ", ")
	color, hasColor := colorMap[tagsStr]
	if !hasColor {
		selectedColor := colors[len(colorMap)]
		colorMap[tagsStr] = selectedColor
		color = selectedColor
	}

	return fmt.Sprintf(`<TR><TD CELLPADDING="5" BGCOLOR="%s"><I><FONT POINT-SIZE="9">%s</FONT></I></TD></TR>`, color, tagsStr)
}
