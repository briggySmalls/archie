package drawers

import (
	"github.com/briggysmalls/archie/internal/views"
)

type TextDrawer interface {
	Draw(views.View) string
}
