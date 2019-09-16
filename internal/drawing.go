package internal

import (
	"github.com/briggysmalls/archie/internal/views"
)

type Drawer interface {
	Draw(views.View)
}
