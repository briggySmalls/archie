package views

import (
	"github.com/briggysmalls/archie/internal/types"
)

type View interface {
	Elements() []*types.Element
	Relationships() []types.Relationship
}
