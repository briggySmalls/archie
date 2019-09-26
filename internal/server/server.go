package server

import (
	"fmt"
	"github.com/briggysmalls/archie/internal/drawers"
	"github.com/briggysmalls/archie/internal/types"
	"github.com/briggysmalls/archie/internal/views"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"net/url"
)

const (
	TEMPLATE_FILE = "/Users/sambriggs/Code/go/archie/internal/server/page.html"
)

func NewServer(model *types.Model) (Server, error) {
	// Load the template
	t, err := template.ParseFiles(TEMPLATE_FILE)
	if err != nil {
		return nil, err
	}
	// Create the server
	return &server{
		model:    model,
		drawer:   drawers.NewMermaidDrawer("http://localhost:8080/"),
		template: t,
	}, nil
}

type Server interface {
	Serve(address string) error
}

type server struct {
	model    *types.Model
	drawer   drawers.Drawer
	template *template.Template
}

type templateData struct {
	Title    string
	ViewName string
	Context  string
	Chart    string
}

func (s *server) Serve(address string) error {
	// Create a router
	r := mux.NewRouter()
	r.HandleFunc("/", s.homeHandler)
	r.HandleFunc("/{item:.*}", s.contextHandler)

	// Serve
	return http.ListenAndServe(address, r)
}

func (s *server) homeHandler(w http.ResponseWriter, r *http.Request) {
	// Create a landscape view
	viewModel := views.NewLandscapeView(s.model)
	// Write plantuml
	output, err := s.drawer.Draw(viewModel)
	if err != nil {
		s.Error(w, err.Error(), 500)
		return
	}
	// Create template data, including the graph
	data := templateData{
		Title:    "Architecture explorer",
		ViewName: "landscape",
		Context:  "",
		Chart:    output,
	}
	// Send populated template as response
	s.template.Execute(w, data)
}

func (s *server) contextHandler(w http.ResponseWriter, r *http.Request) {
	// Determine the item
	itemName, err := url.PathUnescape(mux.Vars(r)["item"])
	if err != nil {
		s.Error(w, err.Error(), http.StatusBadRequest)
	}
	item, err := s.model.LookupName(itemName)
	if err != nil {
		// Failed to find the supplied item (user error)
		s.NotFound(w, r)
		return
	}
	// Create the view
	viewModel := views.NewItemContextView(s.model, item)
	// Write plantuml
	output, err := s.drawer.Draw(viewModel)
	if err != nil {
		// Server error
		s.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Create template data, including the graph
	data := templateData{
		Title:    "Architecture explorer",
		ViewName: "context",
		Context:  itemName,
		Chart:    output,
	}
	// Send populated template as response
	s.template.Execute(w, data)
}

// Our custom HTTP 404 error page
func (s *server) NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Error 404: Can't find that page :'(")
}

// Our custom error page
func (s *server) Error(w http.ResponseWriter, error string, code int) {
	w.WriteHeader(code)
	fmt.Fprintf(w, "Error %d: %s", code, error)
}
