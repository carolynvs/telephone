package gophers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestYodaGopher(t *testing.T) {
	g := NewYodaGopher()

	// Mock the quote API to always return the same quote
	want := "Writes tests, yoda always does."
	mockApi := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Accept", "test/plain")

		switch r.RequestURI {
		case "/?sentence=Yoda+always+writes+tests.":
			fmt.Fprint(w, "Writes tests, yoda always does.")
		}
	}))
	defer mockApi.Close()

	// Tell our gopher to use our mock API
	g.endpoint = mockApi.URL

	got := g.TransformMessage("Yoda always writes tests.")
	if got != want {
		t.Fatalf("Expected '%s', Got '%s'", want, got)
	}
}
