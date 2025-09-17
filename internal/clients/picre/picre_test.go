package picre

import (
	"booru-server/pkg/booru"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_Search(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/image.json" {
			t.Errorf("Expected to request '/image.json', got: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		response := Image{
			ID:      123,
			FileURL: "https://example.com/image.jpg",
			Tags:    []string{"tag1"},
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := New()
	client.BaseURL = server.URL
	params := booru.SearchParams{Tags: []string{"tag1"}}
	images, err := client.Search(context.Background(), params)

	if err != nil {
		t.Fatalf("Search failed: %v", err)
	}

	if len(images) != 1 {
		t.Fatalf("Expected 1 image, got %d", len(images))
	}

	if images[0].ID != "123" {
		t.Errorf("Expected image ID '123', got '%s'", images[0].ID)
	}
}
