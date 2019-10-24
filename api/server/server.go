package server

import (
	"fmt"
	"github.com/briggysmalls/archie"
	"github.com/briggysmalls/archie/writers"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func Serve(address string) error {
	// Create a router
	r := mux.NewRouter()
	r.HandleFunc("/diagram/landscape", landscapeHandler).Methods("POST")
	r.PathPrefix("/diagram/context").HandlerFunc(contextHandler).Methods("POST")

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
	items := r.URL.Query()["scope"]
	if len(items) != 1 {
		errorHandler(w, fmt.Sprintf("Invalid scope '%s'", items), http.StatusBadRequest)
	}
	item := items[0]
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

func readModel(r *http.Request) (archie.Archie, error) {
	// Obtain the model from the request body
	model, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("Cannot read body")
	}
	if model == nil {
		return nil, fmt.Errorf("No model found in request")
	}
	// Create an Archie from the model
	archie, err := archie.New(writers.PlantUmlStrategy{}, string(model))
	return archie, err
}
