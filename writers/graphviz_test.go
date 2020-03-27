package writers

import (
	"gotest.tools/assert"
	"testing"
)

func TestDrawGraphviz(t *testing.T) {
	// Create the test model
	m, elMap := createTestModel()

	// Drawer
	d := New(GraphvizStrategy{})
	output, err := d.Write(*m)
	assert.NilError(t, err)

	const resultFormat = `graph arch {
    graph [fontname=Helvetica]
    edge [fontsize=9; fontname=Helvetica; color="#333333"]
    node [shape=plaintext; margin=0; fontname=Helvetica]
    subgraph "cluster_%[1]s" {
        label = <One>
        "%[2]s" [
            label = <
                <TABLE BORDER="0" CELLBORDER="0" CELLSPACING="0">

                <TR><TD COLSPAN="0" CELLPADDING="10" BGCOLOR="#dbdbdb">OneChild</TD></TR>
                </TABLE>>
        ];
        }
    "%[3]s" [
        label = <
            <TABLE BORDER="0" CELLBORDER="0" CELLSPACING="0">
            <TR><TD CELLPADDING="5" BGCOLOR="#8dd3c7"><I><FONT POINT-SIZE="9">software</FONT></I></TD><TD CELLPADDING="5" BGCOLOR="#ffffb3"><I><FONT POINT-SIZE="9">mechanical</FONT></I></TD></TR>
            <TR><TD COLSPAN="2" CELLPADDING="10" BGCOLOR="#dbdbdb">Two</TD></TR>
            </TABLE>>
    ];
    "%[4]s" [
        color = "#333333"
        shape = circle
        margin = 0.04
        label = User
    ];
    "%[2]s" -- "%[3]s"
}
`
	// Assert result
	assertOutput(t, output, resultFormat, []string{elMap["One"].ID(), elMap["OneChild"].ID(), elMap["Two"].ID(), elMap["User"].ID()})
}
