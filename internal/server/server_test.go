package server

import (
	"github.com/briggysmalls/archie/internal/types"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

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

	mytest(t, s)
	mytest(t, s)
}

func mytest(t *testing.T, s *server) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/context/One", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	// Need to create a router that we can pass the request through
	// so that the vars will be added to the context
	router := mux.NewRouter()
	router.HandleFunc("/context/{item}", s.contextHandler)
	router.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func newServer(t *testing.T) *server {
	// Create a simple model
	m := types.NewModel()
	one := types.NewItem("One")
	two := types.NewItem("Two")
	m.AddRootElement(&one)
	m.AddRootElement(&two)

	// Create the server
	s, err := NewServer(&m)
	if err != nil {
		t.Fatal(err)
	}

	return s.(*server)
}
