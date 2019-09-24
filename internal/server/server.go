package server

import (
	"http"
	"github.com/gorilla/mux"
	"github.com/briggysmalls/archie/internal/types"
	"github.com/briggysmalls/archie/internal/drawers"
	"github.com/briggysmalls/archie/internal/views"
)

func NewServer(model *types.Model) Server {
	// Create the server
	return server{
		model: model,
		drawer: drawers.NewPlantUmlDrawer(),
	}
}

type Server interface {
	Serve(address string) error
}

type server struct {
	model *types.Model
	drawer drawers.Drawer
}

func (s *server) Serve(address string) {
	// Create a router
	r := mux.NewRouter()
	r.HandleFunc("/", s.home)
	r.HandleFunc("/{item}", s.item)

	// Serve
	log.Fatal(http.ListenAndServe(address, nil))
}

func (s *server) home(w http.ResponseWriter, r *http.Request) {
	// Create a landscape view
	viewModel := views.NewLandscapeView(s.model)
	// Write plantuml
	output, err := s.drawer.Draw(viewModel)
	if err != nil {
		w.Write(err)
	}
	// Print
	w.Write(html.EscapeString(output))
}

func (s *server) item(w http.ResponseWriter, r *http.Request) {
	// Determine the item
	itemName := mux.Vars(r)["item"]
	item := s.model.LookupName(itemName)
	// Create the view
	viewModel := views.NewItemContextView(s.model)
	// Write plantuml
	output, err := d.drawer.Draw(viewModel)
	if err != nil {
		w.Write(err)
	}
	// Print
	w.Write(html.EscapeString(output))
}
