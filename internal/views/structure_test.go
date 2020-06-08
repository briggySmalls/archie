package views

import (
	"testing"

	mdl "github.com/briggysmalls/archie/internal/model"
	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

// TestStrucureScoped Test generating non scoped diagram
func TestStrucureWithutScope(t *testing.T) {
	m := mdl.NewModel()
	m.AddRootElement(mdl.NewItem("1", []string{}))
	m.AddRootElement(mdl.NewItem("2", []string{}))
	m.AddRootElement(mdl.NewActor("no"))

	v := NewStructureView(&m, nil)
	assert.Assert(t, is.Len(v.Elements, 2))
	assert.Assert(t, is.Len(v.Associations, 0))
}

// TestStrucureScoped Test generating scoped structure diagram
func TestStrucureScoped(t *testing.T) {
	m := mdl.NewModel()
	two := mdl.NewItem("2", []string{})
	twokid := mdl.NewItem("2.1", []string{})
	m.AddRootElement(mdl.NewItem("1", []string{}))
	m.AddRootElement(two)
	m.AddElement(twokid, two)

	v := NewStructureView(&m, two)
	assert.Assert(t, is.Len(v.Elements, 2))
	assert.Assert(t, is.Len(v.Associations, 1))
	assert.Assert(t, is.Equal(v.Associations[0].Source(), two))
	assert.Assert(t, is.Equal(v.Associations[0].Destination(), twokid))
}
