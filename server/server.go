package server

import (
	"fmt"
	"github.com/briggysmalls/archie/core"
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

func NewServer(archie core.Archie) (Server, error) {
	// Load the template
	t, err := template.New("page").Parse(page)
	if err != nil {
		return nil, err
	}
	// Create the server
	return &server{
		archie:   archie,
		template: t,
	}, nil
}

type Server interface {
	Serve(address string) error
}

type server struct {
	archie   core.Archie
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
	chart, err := s.archie.LandscapeView()
	if err != nil {
		s.Error(w, err.Error(), http.StatusBadRequest)
	}
	// Create template data, including the graph
	data := templateData{
		Title:    "Architecture explorer",
		ViewName: "landscape",
		Context:  "",
		Chart:    chart,
	}
	// Send populated template as response
	s.template.Execute(w, data)
}

func (s *server) contextHandler(w http.ResponseWriter, r *http.Request) {
	// Determine the item
	item, err := url.PathUnescape(strings.TrimFunc(r.URL.Path, func(r rune) bool {
		return r == '/'
	}))
	if err != nil {
		s.Error(w, err.Error(), http.StatusBadRequest)
	}
	// Create the view
	chart, err := s.archie.ContextView(item)
	if err != nil {
		s.Error(w, err.Error(), http.StatusBadRequest)
	}
	// Create template data, including the graph
	data := templateData{
		Title:    "Architecture explorer",
		ViewName: "context",
		Context:  item,
		Chart:    chart,
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
