package server

import (
	"fmt"
	"github.com/briggysmalls/archie"
	"github.com/briggysmalls/archie/cli/archie/utils"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Serve a REST API exposing Archie functionality
func Serve(address string) error {
	// Create a router
	r := mux.NewRouter()
	r.PathPrefix("/diagram/context").HandlerFunc(contextHandler).Methods("POST")
	r.PathPrefix("/diagram/tag").HandlerFunc(tagHandler).Methods("POST")

	// Serve
	return http.ListenAndServe(address, r)
}

func contextHandler(w http.ResponseWriter, r *http.Request) {
	// Get the model
	archie, err := readModel(r)
	if err != nil {
		errorHandler(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Determine the item
	scope, err := readSingleParameter(r.URL, "scope", true)
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
	scope, err := readSingleParameter(r.URL, "scope", true)
	if err != nil {
		errorHandler(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Determine the tag
	tag, err := readSingleParameter(r.URL, "tag", false)
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

func readSingleParameter(url *url.URL, parameter string, allowEmpty bool) (value string, err error) {
	items := url.Query()[parameter]
	switch len(items) {
	case 1:
		// If there is exactly one, we're always happy
		return items[0], nil
	case 0:
		if allowEmpty {
			// We're allowed to return an empty value
			return "", nil
		}
		// Otherwise error
		fallthrough
	default:
		return "", fmt.Errorf("Invalid %s '%s'", parameter, items)
	}
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
	// Parse the model and config into an archie instance
	return utils.ReadModel(body)
}
