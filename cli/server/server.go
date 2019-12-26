package server

import (
	"fmt"
	"github.com/briggysmalls/archie"
	"github.com/briggysmalls/archie/writers"
	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"net/url"
)

type payload struct {
	Model  interface{} `yaml:""`
	Config interface{} `yaml:""`
}

var customFooter string

func Serve(address, footer string) error {
	// Record the custom footer
	customFooter = footer
	// Create a router
	r := mux.NewRouter()
	r.HandleFunc("/diagram/landscape", landscapeHandler).Methods("POST")
	r.PathPrefix("/diagram/context").HandlerFunc(contextHandler).Methods("POST")
	r.PathPrefix("/diagram/tag").HandlerFunc(tagHandler).Methods("POST")

	// Serve
	return http.ListenAndServe(address, r)
}

func landscapeHandler(w http.ResponseWriter, r *http.Request) {
	// Get the model
	archie, err := readModel(r)
	if err != nil {
		errorHandler(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Create a landscape view
	chart, err := archie.LandscapeView()
	if err != nil {
		errorHandler(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Return diagram in browser
	fmt.Fprintf(w, chart)
}

func contextHandler(w http.ResponseWriter, r *http.Request) {
	// Get the model
	archie, err := readModel(r)
	if err != nil {
		errorHandler(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Determine the item
	scope, err := readSingleParameter(r.URL, "scope")
	if err != nil {
		errorHandler(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Create the view
	chart, err := archie.ContextView(scope)
	if err != nil {
		errorHandler(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Return diagram in browser
	fmt.Fprintf(w, chart)
}

func tagHandler(w http.ResponseWriter, r *http.Request) {
	// Get the model
	archie, err := readModel(r)
	if err != nil {
		errorHandler(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Determine the item
	scope, err := readSingleParameter(r.URL, "scope")
	if err != nil {
		errorHandler(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Determine the tag
	tag, err := readSingleParameter(r.URL, "tag")
	if err != nil {
		errorHandler(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Create the view
	chart, err := archie.TagView(scope, tag)
	if err != nil {
		errorHandler(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Return diagram in browser
	fmt.Fprintf(w, chart)
}

func readSingleParameter(url *url.URL, parameter string) (value string, err error) {
	items := url.Query()[parameter]
	if len(items) != 1 {
		return "", fmt.Errorf("Invalid %s '%s'", parameter, items)
	}
	return items[0], nil
}

// Our custom error page
func errorHandler(w http.ResponseWriter, error string, code int) {
	w.WriteHeader(code)
	fmt.Fprintf(w, "Error %d: %s", code, error)
}

func readModel(r *http.Request) (archie.Archie, error) {
	// Obtain the model and config from the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("Cannot read body")
	}
	if body == nil {
		return nil, fmt.Errorf("No payload found in request")
	}
	// Separate config and model
	p := payload{}
	err = yaml.Unmarshal(body, &p)
	if err != nil {
		return nil, err
	}
	// Create an Archie from the model
	model, err := yaml.Marshal(p.Model)
	if err != nil {
		return nil, err
	}
	archie, err := archie.New(writers.PlantUmlStrategy{CustomFooter: customFooter}, string(model))
	return archie, err
}
