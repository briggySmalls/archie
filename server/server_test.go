package server

import (
	"github.com/briggysmalls/archie/core"
	"github.com/briggysmalls/archie/core/writers"
	"github.com/gorilla/mux"
	"gotest.tools/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var model = `
elements:
  - name: user
    kind: actor
  - name: sound system
    children:
      - one
      - two
      - three
associations:
  - source: user
    destination: sound system/one
  - source: sound system/two
    destination: sound system/three
  - source: sound system/three
    destination: user
`

func TestHomeHandler(t *testing.T) {
	// Create the server
	s := newServer(t)

	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(s.homeHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestContextHandler(t *testing.T) {
	// Create the server
	s := newServer(t)

	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/One", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	// Need to create a router that we can pass the request through
	// so that the vars will be added to the context
	router := mux.NewRouter()
	router.HandleFunc("/{item}", s.contextHandler)
	router.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func newServer(t *testing.T) *server {
	// Create a simple model
	// Create the api
	a, err := core.New(writers.MermaidStrategy{}, model)
	assert.NilError(t, err)

	// Create the server
	s, err := NewServer(a)
	assert.NilError(t, err)

	return s.(*server)
}
