package waifuim

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
		if r.URL.Path != "/search" {
			t.Errorf("Expected to request '/search', got: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		response := Response{
			Images: []Image{
				{
					ImageID: 123,
					URL:     "https://example.com/image.jpg",
					Tags:    []Tag{{Name: "tag1"}, {Name: "tag2"}},
				},
			},
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := New(server.URL)
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
	if images[0].URL != "https://example.com/image.jpg" {
		t.Errorf("Expected image URL 'https://example.com/image.jpg', got '%s'", images[0].URL)
	}
}

func TestClient_Search_Error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	client := New(server.URL)
	_, err := client.Search(context.Background(), booru.SearchParams{})
	if err == nil {
		t.Fatal("Expected an error, but got none")
	}
}
