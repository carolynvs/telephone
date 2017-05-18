package gophers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestChuckNorrisGopher(t *testing.T) {
	g := ChuckNorrisGopher{}

	// Mock the quote API to always return the same quote
	want := "Chuck Norris always unit tests"
	mockApi := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		switch r.RequestURI {
		case "/":
			fmt.Fprintln(w, `{ "type": "success", "value": { "id": 1, "joke": "`+want+`", "categories": [] } }`)
		}
	}))
	defer mockApi.Close()

	// Tell our gopher to use our mock API
	g.endpoint = mockApi.URL

	got := g.TransformMessage("I love puppies!")
	if got != want {
		t.Fatalf("Expected '%s', Got '%s'", want, got)
	}
}
