package nekosapi

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
		if r.URL.Path != "/images/random" {
			t.Errorf("Expected to request '/images/random', got: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		response := Response{
			Data: []Image{
				{
					ID:      "test-id",
					FileURL: "https://example.com/image.jpg",
					Tags:    []Tag{{Name: "tag1"}},
				},
			},
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

	if images[0].ID != "test-id" {
		t.Errorf("Expected image ID 'test-id', got '%s'", images[0].ID)
	}
}
