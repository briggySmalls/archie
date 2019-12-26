package server

import (
  "gotest.tools/assert"
  is "gotest.tools/assert/cmp"
  "net/http"
  "net/http/httptest"
  "strings"
  "testing"
)

var model = `
config:
  footer: |
    skinparam nodesep 10
    skinparam nodesep 10
model:
  elements:
    - name: user
      kind: actor
    - name: sound-system
      children:
        - name: amplifier
          children:
            - name: audio in connector
              tags: [electronic]
            - name: audio out connector
              tags: [electronic]
            - name: bluetooth receiver
              tags: [electronic]
            - name: ac-dc converter
              tags: [electronic]
            - name: mixer
              tags: [electronic]
            - name: amplifier
              tags: [electronic]
            - name: power button
              tags: [electronic, mechanical]
            - name: input select
              tags: [electronic, mechanical]
  associations:
    # Sound system
    - source: user
      destination: sound-system/amplifier/input select
    - source: sound-system/amplifier/input select
      destination: sound-system/amplifier/mixer
    - source: sound-system/amplifier/audio in connector
      destination: sound-system/amplifier/mixer
    - source: sound-system/amplifier/bluetooth receiver
      destination: sound-system/amplifier/mixer
    - source: sound-system/amplifier/ac-dc converter
      destination: sound-system/amplifier/mixer
    - source: sound-system/amplifier/mixer
      destination: sound-system/amplifier/amplifier
    - source: sound-system/amplifier/ac-dc converter
      destination: sound-system/amplifier/amplifier
    - source: sound-system/amplifier/amplifier
      destination: sound-system/amplifier/audio out connector
    - source: sound-system/amplifier/power button
      destination: sound-system/amplifier/ac-dc converter
    - source: user
      destination: sound-system/amplifier/power button
`

func TestLandscapeHandler(t *testing.T) {
  // Create a request to pass to our handler. We don't have any query parameters for now, so we'll
  // pass 'nil' as the third parameter.
  req, err := http.NewRequest("POST", "diagram/landscape", strings.NewReader(model))
  if err != nil {
    t.Fatal(err)
  }

  // We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
  rr := httptest.NewRecorder()
  handler := http.HandlerFunc(landscapeHandler)

  // Our handlers satisfy http.Handler, so we can call their ServeHTTP method
  // directly and pass in our Request and ResponseRecorder.
  handler.ServeHTTP(rr, req)

  // Check the status code is what we expect.
  if status := rr.Code; status != http.StatusOK {
    t.Errorf("handler returned wrong status code: got %v want %v",
      status, http.StatusOK)
  }

  // Check the content is what we expect
  body := rr.Body.String()
  assert.Assert(t, is.Contains(body, "rectangle \"sound-system\" as"))
  assert.Assert(t, is.Contains(body, "actor \"user\" as"))
  assert.Assert(t, is.Contains(body, "skinparam nodesep 10"))
}

func TestContextHandler(t *testing.T) {
  // Create a request to pass to our handler. We don't have any query parameters for now, so we'll
  // pass 'nil' as the third parameter.
  req, err := http.NewRequest("POST", "diagram/context?scope=sound-system", strings.NewReader(model))
  if err != nil {
    t.Fatal(err)
  }

  // We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
  rr := httptest.NewRecorder()
  handler := http.HandlerFunc(contextHandler)

  // Our handlers satisfy http.Handler, so we can call their ServeHTTP method
  // directly and pass in our Request and ResponseRecorder.
  handler.ServeHTTP(rr, req)

  // Check the status code is what we expect.
  if status := rr.Code; status != http.StatusOK {
    t.Errorf("handler returned wrong status code: got %v want %v",
      status, http.StatusOK)
  }
}

func TestTagHandler(t *testing.T) {
  // Create a request to pass to our handler. We don't have any query parameters for now, so we'll
  // pass 'nil' as the third parameter.
  req, err := http.NewRequest("POST", "diagram/tag?scope=sound-system&tag=software", strings.NewReader(model))
  if err != nil {
    t.Fatal(err)
  }

  // We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
  rr := httptest.NewRecorder()
  handler := http.HandlerFunc(tagHandler)

  // Our handlers satisfy http.Handler, so we can call their ServeHTTP method
  // directly and pass in our Request and ResponseRecorder.
  handler.ServeHTTP(rr, req)

  // Check the status code is what we expect.
  if status := rr.Code; status != http.StatusOK {
    t.Errorf("handler returned wrong status code: got %v want %v",
      status, http.StatusOK)
  }
}
