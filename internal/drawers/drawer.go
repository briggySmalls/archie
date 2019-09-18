package drawers

import (
	"github.com/briggysmalls/archie/internal/types"
)

type TextDrawer interface {
	Draw(types.Model) string
}
