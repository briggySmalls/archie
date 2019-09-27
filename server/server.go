package server

import (
	"fmt"
	mdl "github.com/briggysmalls/archie/core/model"
	"github.com/briggysmalls/archie/core/views"
	"github.com/briggysmalls/archie/io/writers"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"net/url"
	"strings"
)

const page = `
<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="utf-8">
        <title>{{.Title}}</title>
        {{/* Include mermaid stylesheets */}}
        <link rel="stylesheet" href="mermaid.min.css">
    </head>
    <body>
        <h1>{{.Title}}</h1>
        {{if .ViewName}}<h2>View: {{.ViewName}}</h2>{{end}}
        {{if .Context}}<h3>Context: {{.Context}} <a href=".">⬆️</a></h3>{{end}}
        <div class="mermaid">
            {{.Chart}}
        </div>
        {{/* Include mermaid.js */}}
        <script src="https://cdnjs.cloudflare.com/ajax/libs/mermaid/8.3.1/mermaid.min.js"></script>
        {{/* Configure/run mermaid.js */}}
        <script>
            mermaid.initialize({
                startOnLoad:true, // TODO: Can we remove this?
                securityLevel: 'loose', // Allow click events
            });
        </script>
    </body>
</html>
`

func NewServer(model *mdl.Model) (Server, error) {
	// Load the template
	t, err := template.New("page").Parse(page)
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
	model    *mdl.Model
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
	r.PathPrefix("/").HandlerFunc(s.contextHandler)

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
	itemName, err := url.PathUnescape(strings.TrimFunc(r.URL.Path, func(r rune) bool {
		return r == '/'
	}))
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
	viewModel := views.NewContextView(s.model, item)
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
