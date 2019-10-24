package server

import (
	"fmt"
	"github.com/briggysmalls/archie"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"net/url"
	"strings"
)

func Serve(address string) error {
	// Create a router
	r := mux.NewRouter()
	r.HandleFunc("/diagram/landscape", landscapeHandler).Methods("POST")
	r.PathPrefix("/diagram/context/").HandlerFunc(contextHandler).Methods("POST")

	// Serve
	return http.ListenAndServe(address, r)
}

func landscapeHandler(w http.ResponseWriter, r *http.Request) {
	// Get the model
	archie, err := readModel(r)
	if err != nil {
		errorHandler(w, err.Error(), http.StatusBadRequest)
	}
	// Create a landscape view
	chart, err := archie.LandscapeView()
	if err != nil {
		errorHandler(w, err.Error(), http.StatusBadRequest)
	}
	// Return diagram in browser
	fmt.Fprintf(w, chart)
}

func contextHandler(w http.ResponseWriter, r *http.Request) {
	// Get the model
	archie, err := readModel(r)
	if err != nil {
		errorHandler(w, err.Error(), http.StatusBadRequest)
	}
	// Determine the item
	item, err := url.PathUnescape(strings.TrimFunc(r.URL.Path, func(r rune) bool {
		return r == '/'
	}))
	if err != nil {
		errorHandler(w, err.Error(), http.StatusBadRequest)
	}
	// Create the view
	chart, err := archie.ContextView(item)
	if err != nil {
		errorHandler(w, err.Error(), http.StatusBadRequest)
	}
	// Return diagram in browser
	fmt.Fprintf(w, chart)
}

// Our custom error page
func errorHandler(w http.ResponseWriter, error string, code int) {
	w.WriteHeader(code)
	fmt.Fprintf(w, "Error %d: %s", code, error)
}

func readModel(r *http.Request) (archie.Archie, err) {
	// Obtain the model from the request body
	model := r.Body
	if model == nil {
		return "", fmt.Errorf("No model found in request")
	}
	// Create an Archie from the model
	archie, err := archie.New(model)
	return archie, nil
}
