package nekosapi

import (
	"booru-server/pkg/booru"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestModule_Search(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/images/random" {
			t.Errorf("Expected to request '/images/random', got: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		response := []Image{
			{
				ID:      "test-id",
				FileURL: "https://example.com/image.jpg",
				Tags:    []Tag{{Name: "tag1"}},
			},
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	module := New(WithBaseURL(server.URL))
	params := booru.SearchParams{Tags: []string{"tag1"}}
	images, err := module.Search(context.Background(), params)

	if err != nil {
		t.Fatalf("Search failed: %v", err)
	}

	if len(images) != 1 {
		t.Fatalf("Expected 1 image, got %d", len(images))
	}

	if images[0].ID != "test-id" {
		t.Errorf("Expected image ID 'test-id', got '%s'", images[0].ID)
	}
}

func TestModule_Search_Error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	module := New(WithBaseURL(server.URL))
	_, err := module.Search(context.Background(), booru.SearchParams{})
	if err == nil {
		t.Fatal("Expected an error, but got none")
	}
}
